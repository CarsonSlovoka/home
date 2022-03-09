package main

import (
	io2 "carson.io/pkg/io"
	"carson.io/pkg/tpl/funcs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

type Config struct {
	excludeFiles []string
}

var config *Config

func init() {
	config = &Config{
		[]string{
			`url\\static\\css\\.*\.md`,
			`url\\static\\img\\.*\.md`,
			`url\\static\\img\\.*\.md`,
			`url\\static\\js\\.*\.md`,
			`url\\static\\sass\\.*`,
			`url\\tmpl\\.*`, // 樣版在release不需要再給，已經遷入到source之中
		},
	}
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

	t := template.Must(
		template.New(filepath.Base(src)).
			Funcs(funcs.GetUtilsFuncMap()).
			ParseFiles(parseFiles...),
	)
	context := struct {
		Version string
	}{
		"0.0.0",
	}
	return t.Execute(dstFile, context)
}

func main() {
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
		if err = mirrorDir("url\\", "..\\docs\\", []string{
			"url\\pkg",
			"url\\static\\sass",
		}); err != nil {
			panic(err)
		}
	}

	{ // and then copy file
		filePathList, _ := collectFiles("url\\", config.excludeFiles)
		for _, src := range filePathList {
			dst := filepath.Join("../docs/", strings.Replace(src, "url\\", "", 1))
			// fmt.Println(dst)
			if filepath.Ext(dst) == ".gohtml" {
				dst = dst[:len(dst)-6] + "html"
				if err = render(src, dst); err != nil {
					panic(err)
				}
				continue
			}

			if err = io2.CopyFile(src, dst); err != nil {
				panic(err)
			}
		}
	}
}
