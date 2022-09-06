package funcs

import (
	"bytes"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/tpl/funcs"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"html/template"
	"os"
	"path/filepath"
)

var markdown goldmark.Markdown

func init() {
	markdown = goldmark.New(goldmark.WithRendererOptions(
		html.WithUnsafe(),
	))
}

func GetUtilsFuncMap() map[string]any {
	funcMap := funcs.GetUtilsFuncMap()
	funcMap["md"] = func(srcPath string, ctx any) template.HTML { // 回傳值如果是普通的string，不會轉成HTML會被當成一般文字
		rootDir := "url"
		buf := bytes.NewBuffer(make([]byte, 0))
		srcBytes, err := os.ReadFile(filepath.Join(rootDir, srcPath))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "markdown readfile error. srcPath:%s, err: %s\n", srcPath, err)
			return ""
		}
		if err = markdown.Convert(srcBytes, buf); err != nil {
			panic(err)
		}
		return template.HTML(buf.String())
	}
	return funcMap
}
