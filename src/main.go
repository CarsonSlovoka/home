package main

import (
	io2 "carson.io/pkg/io"
	"carson.io/pkg/tpl/funcs"
	"errors"
	"fmt"
	htmlTemplate "html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	textTemplate "text/template"
	"time"
)

type Config struct {
	*Server
	excludeFiles []string
	*SiteContext
}

type SiteContext struct {
	Title       string
	Version     string
	LastModTime string
}

type Server struct {
	Port int
}

var (
	config   *Config
	chanQuit chan error
)

func init() {
	now := time.Now()
	config = &Config{
		&Server{8888},
		[]string{
			`url\\static\\css\\.*\.md`,
			`url\\static\\img\\.*\.md`,
			`url\\static\\img\\.*\.md`,
			`url\\static\\js\\.*\.md`,
			`url\\static\\sass\\.*`,
			`url\\tmpl\\.*`, // 樣版在release不需要再給，已經遷入到source之中
		}, &SiteContext{
			"Carson-Blog",
			"0.0.0",
			now.Format("2006-01-02 15:04"),
		},
	}
	chanQuit = make(chan error)
}

func render(src, dst string) error {
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	tmplDir := "url/tmpl"
	parseFiles := []string{src}
	for _, filename := range []string{"head", "navbar"} {
		parseFiles = append(parseFiles, filepath.Join(tmplDir, filename+".gohtml"))
	}

	t := textTemplate.Must(
		textTemplate.New(filepath.Base(src)).
			Funcs(funcs.GetUtilsFuncMap()).
			ParseFiles(parseFiles...),
	)
	return t.Execute(dstFile, config.SiteContext)
}

func build(outputDir string) error {
	mirrorDir := func(rootSrc string, dst string, excludeList []string) error {
		return filepath.Walk(rootSrc, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && (
				path != rootSrc &&
					func(curPath string) bool { // filter
						for _, excludeItem := range excludeList {
							if strings.HasPrefix(curPath, excludeItem) {
								return false
							}
						}
						return true
					}(path)) {

				dstPath := filepath.Join(dst, strings.Replace(path, rootSrc, "", 1))
				// fmt.Println(dstPath)
				if err = os.MkdirAll(dstPath, os.FileMode(666)); err != nil {
					return err
				}
			}
			return nil
		})
	}

	collectFiles := func(dir string, excludeList []string) (fileList []string, err error) {
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if regexp.MustCompile(strings.Join(excludeList, "|")).Match([]byte(path)) {
				// fmt.Printf("%s\n", path)
				return nil
			}

			// fmt.Println(path)
			fileList = append(fileList, path)
			return nil
		})
		if err != nil {
			log.Fatalf("walk error [%v]\n", err)
			return nil, err
		}
		return fileList, nil
	}

	var err error
	{ // Copy Dir only
		if err = mirrorDir("url\\", outputDir, []string{
			"url\\pkg",
			"url\\static\\sass",
		}); err != nil {
			panic(err)
		}
	}

	{ // and then copy file
		filePathList, _ := collectFiles("url\\", config.excludeFiles)
		for _, src := range filePathList {
			// dst := filepath.Join("../docs/", strings.Replace(src, "url\\", "", 1)) // filepath.Join反斜線會自動修正，所以這樣也可以
			dst := filepath.Join(outputDir, strings.Replace(src, "url\\", "", 1))
			// fmt.Println(dst)
			if filepath.Ext(dst) == ".gohtml" {
				dst = dst[:len(dst)-6] + "html"
				if err = render(src, dst); err != nil {
					return err
				}
				continue
			}

			if err = io2.CopyFile(src, dst); err != nil {
				return err
			}
		}
	}
	return nil
}

type RootDir struct {
	http.Dir
}

func (dir *RootDir) Open(name string) (http.File, error) {
	if filepath.Ext(name) == ".sass" {
		return nil, errors.New(fmt.Sprintf("%d", http.StatusForbidden))
	}
	return dir.Dir.Open(name)
}

type RootHandler struct {
	http.HandlerFunc
}

func (handler *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	switch filepath.Ext(r.URL.Path) {
	case ".html":
		r.URL.Path = r.URL.Path[:len(r.URL.Path)-4] + "gohtml" // treat all of html as gohtml
		fallthrough
	case ".gohtml":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		src := filepath.Join("./url" + r.URL.Path)
		tmplDir := "url/tmpl"
		parseFiles := []string{src}
		for _, filename := range []string{"head", "navbar"} {
			parseFiles = append(parseFiles, filepath.Join(tmplDir, filename+".gohtml"))
		}

		t := htmlTemplate.Must(
			htmlTemplate.New(filepath.Base(src)).
				Funcs(htmlTemplate.FuncMap(funcs.GetUtilsFuncMap())).
				ParseFiles(parseFiles...),
		)

		if err := t.Execute(w, config.SiteContext); err != nil {
			_ = fmt.Errorf("%s\n", err.Error())
		}
		return
		/* // 交給http.FileServer(http.Dir()).ServeHTTP(w, r)已經會自行處理MIME_types
		case ".js":
			// w.Header().Set("Content-Type", "text/javascript; charset=utf-8") // Expected a JavaScript module script but the server responded with a MIME type of "
		case ".css":
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		*/
	}
	handler.HandlerFunc(w, r)
}

func run() error {
	mux := http.NewServeMux()
	rootDir := &RootDir{http.Dir("./url/")}
	rootHandler := &RootHandler{func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(rootDir).ServeHTTP(w, r)
	}}

	mux.Handle("/", rootHandler)
	server := http.Server{Addr: fmt.Sprintf(":%d", config.Server.Port), Handler: mux}

	fmt.Printf("http://localhost:%d\n", config.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		chanQuit <- err
		return err
	}
	return nil
}

func main() {
	go startCMD(&chanQuit)
	for {
		select {
		// case <-chanQuit:
		case err := <-chanQuit:
			log.Printf("Close App. %+v\n", err)
			close(chanQuit)
			return
		}
	}
}
