package funcs

import (
	"bytes"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/tpl/funcs"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"reflect"
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
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
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
	funcMap["debug"] = func(a ...any) string {
		log.Printf("%+v", a)
		return fmt.Sprintf("%+v", a)
	}
	funcMap["timeStr"] = func(t time.Time) string {
		// t.Format("2006-01-02 15:04") // 到分感覺沒有意義
		return t.Format("2006-01-02")
	}

	funcMap["setVal"] = func(obj any, key string, val any) (string, error) {
		// ps := reflect.ValueOf(obj) // pointer to struct - addressable
		// s := ps.Elem()             // struct
		s := reflect.ValueOf(obj)
		if s.Kind() != reflect.Struct {
			return "", fmt.Errorf("type error. 'Struct' expected\n")
		}
		field := s.FieldByName(key)
		if !field.IsValid() {
			return "", fmt.Errorf("key not found: %s\n", key)
		}

		if !field.CanSet() { // 只能對pointer類才能異動數值
			return "", fmt.Errorf("The field[%s] is unchangeable. You can't change it.\n", key)
		}
		field.Set(reflect.ValueOf(val))
		return "", nil
	}
	return funcMap
}
