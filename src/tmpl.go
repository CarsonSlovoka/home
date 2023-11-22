package main

import (
	"bytes"
	bytes2 "carson.io/pkg/bytes"
	"carson.io/pkg/tpl/funcs"
	. "carson.io/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

var markdown goldmark.Markdown

func init() {
	markdown = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, // 包含 Linkify, Table, Strikethrough, TaskList
			extension.Footnote,
			highlighting.Highlighting,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // 會自動加上id，但如果有中文heading會不支持

			// https://github.com/yuin/goldmark#attributes
			parser.WithAttribute(), // 推薦補上這個，可以在heading旁邊利用## MyH1{class="cls1 cls2"}來補上一些屬性 // https://www.markdownguide.org/extended-syntax/#heading-ids // https://github.com/gohugoio/hugo/issues/7548
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
}

type SiteContext struct {
	RootTitle     string // 預設title
	Version       string
	LastBuildTime time.Time `json:"bTime"`
}

// PageContext template.Execute 所用到的上下文
type PageContext struct {
	// SiteContext 網站全域的上下文
	SiteContext

	// Filepath 可以是html的路徑或者是md的路徑，主要用途是給md用layout渲染所用
	Filepath string

	// FrontMatter 單獨md檔案寫的frontMatter資料，不過我們使用嵌入，因此就算你的一般HTML之中沒有frontMatter，也可以取用到相關的欄位
	FrontMatter

	// TODO 產生出toc的內容 只會在單獨md使用layout的方式才會自動帶入 (需要規劃爬取md檔案，而非用HTML的h1~h6)
	// TableOfContents 以ul的形式
	// 如果您在一般的自定義HTML頁面，也想用某個md的toc來製作，可以用toc來幫助: ex: {{toc (md "my.md" .)}}
	// TableOfContents template.HTML

	context.Context
}

var tmplFuncMaps map[string]any

func (s *SiteContext) String() string {
	v, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		PErr.Printf("[SiteContext] json marshal error. %s", err)
		return ""
	}
	return string(v)
}

func init() {
	tmplFuncMaps = funcs.GetUtilsFuncMap()

	tmplFuncMaps["md"] = func(srcPath string, ctx *PageContext) template.HTML { // 回傳值如果是普通的string，不會轉成HTML會被當成一般文字
		rootDir := "url"
		buf := bytes.NewBuffer(make([]byte, 0))
		var (
			srcBytes []byte
			err      error
		)

		srcBytes, err = os.ReadFile(filepath.Join(rootDir, srcPath))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "markdown readfile error. srcPath:%s, err: %s\n", srcPath, err)
			return ""
		}

		_, srcBytes, err = bytes2.GetFrontMatter[any](srcBytes, true)

		if err = markdown.Convert(srcBytes, buf); err != nil {
			panic(err)
		}

		return template.HTML(buf.String())
	}

	tmplFuncMaps["toc"] = func(html template.HTML) template.HTML {
		tocNodes := ParseHTMLAsTOC(string(html))
		if tocNodes == nil {
			return ""
		}
		return renderToc(tocNodes, "toc")
	}
}
