package main

import (
	io2 "carson.io/pkg/io"
	"fmt"
	filepath2 "github.com/CarsonSlovoka/go-pkg/v2/path/filepath"
	"github.com/CarsonSlovoka/go-pkg/v2/tpl/funcs"
	"github.com/CarsonSlovoka/go-pkg/v2/tpl/template"
	htmlTemplate "html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
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
		&Server{0}, // port為0可以自動找尋沒有被使用到的port
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

func render(src, dst string, tmplFiles []string) error {
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	parseFiles, err := template.GetAllTmplName(os.ReadFile, src, tmplFiles)
	if err != nil {
		return err
	}
	parseFiles = append(parseFiles, src)

	t := textTemplate.Must(
		textTemplate.New(filepath.Base(src)).
			Funcs(funcs.GetUtilsFuncMap()).
			ParseFiles(parseFiles...),
	)
	return t.Execute(dstFile, config.SiteContext)
}

func build(outputDir string) error {
	var (
		mirrorDir func(rootSrc string, dst string, excludeList []string) error
	)
	{ // init function
		mirrorDir = func(rootSrc string, dst string, excludeList []string) error {
			return filepath.Walk(rootSrc, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() && (path != rootSrc &&
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
	}

	{ // Copy Dir only
		if err := mirrorDir("url\\", outputDir, []string{
			"url\\pkg",
			"url\\static\\sass",
		}); err != nil {
			panic(err)
		}
	}

	{ // and then copy file
		tmplFiles, err := filepath2.CollectFiles("url/tmpl", []string{"\\.md$"})
		if err != nil {
			return err
		}

		filePathList, _ := filepath2.CollectFiles("url\\", config.excludeFiles)
		for _, src := range filePathList {
			// dst := filepath.Join("../docs/", strings.Replace(src, "url\\", "", 1)) // filepath.Join反斜線會自動修正，所以這樣也可以
			dst := filepath.Join(outputDir, strings.Replace(src, "url\\", "", 1))
			// fmt.Println(dst)
			if filepath.Ext(dst) == ".gohtml" {
				dst = dst[:len(dst)-6] + "html"
				if err = render(src, dst, tmplFiles); err != nil {
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

func BuildServer(isLocalMode bool) (server *http.Server, listener net.Listener) {
	mux := http.NewServeMux()

	tmplFiles, err := filepath2.CollectFiles("url/tmpl", []string{"\\.md$"})
	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rootDir := http.Dir("./url/")
		curFilepath := filepath.Join(string(rootDir), r.URL.Path)
		extName := filepath.Ext(curFilepath)
		switch extName {
		case ".sass":
			w.WriteHeader(http.StatusForbidden) // 如果直接返回, status沒有註明，得到的會是一個空頁面(會不曉得對錯)
			return
		case ".html":
			r.URL.Path = r.URL.Path[:len(r.URL.Path)-4] + "gohtml" // treat all of html as gohtml
			fallthrough
		case ".gohtml":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			src := filepath.Join(string(rootDir), r.URL.Path)
			var parseFiles []string
			parseFiles, err = template.GetAllTmplName(os.ReadFile, src, tmplFiles)
			parseFiles = append(parseFiles, src)

			t := htmlTemplate.Must(
				htmlTemplate.New(filepath.Base(src)).
					Funcs(funcs.GetUtilsFuncMap()).
					ParseFiles(parseFiles...),
			)

			if err := t.Execute(w, config.SiteContext); err != nil {
				log.Printf("%s\n", err.Error())
			}
		/* 交給http.FileServer(http.Dir()).ServeHTTP(w, r)已經會自行處理MIME_types
		case ".js":
			// w.Header().Set("Content-Type", "text/javascript; charset=utf-8") // Expected a JavaScript module script but the server responded with a MIME type of "
		case ".css":
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		*/
		default:
			http.FileServer(rootDir).ServeHTTP(w, r)
		}
	})

	if isLocalMode {
		server = &http.Server{Addr: fmt.Sprintf("127.0.0.1:%d", config.Server.Port), Handler: mux}
	} else {
		server = &http.Server{Addr: fmt.Sprintf(":%d", config.Server.Port), Handler: mux}
	}

	listener, err = net.Listen("tcp", server.Addr)
	if err != nil {
		panic(err)
	}

	return server, listener
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go startCMD(&wg)
	wg.Wait()
	log.Printf("Close App.")
}
