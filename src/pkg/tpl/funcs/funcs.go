package funcs

import (
	"bytes"
	. "carson.io/pkg/utils"
	"encoding/json"
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
	"regexp"
	"strconv"
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

type SiteContext struct {
	Title            string    // 頁面的title
	Version          string    // 可以考慮是否移除，目前用處可能不大，或者放到about?
	LastBuildTime    time.Time // 表示此頁面被建立的日期
	LastModTime      time.Time // 頁面的修改日期，建議在各別頁面在自己設定
	EnableMarkMapToc bool      // 預設啟用

	TableOfContents template.HTML // 以ul的形式，產生出toc的內容

}

func (s *SiteContext) String() string {
	v, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		PErr.Printf("[SiteContext] json marshal error. %s", err)
		return ""
	}
	return string(v)
}

type TOCNode struct {
	Depth    int
	Content  string
	parent   *TOCNode
	Children []*TOCNode
}

func renderToc(nodes []*TOCNode, ulClassName string) template.HTML {
	var result string
	if ulClassName != "" {
		result = fmt.Sprintf(`<ul class="%s">`, ulClassName)
	} else {
		result = "<ul>"
	}

	for _, node := range nodes {
		result += "<li>" + node.Content
		if len(node.Children) > 0 {
			result += string(renderToc(node.Children, ""))
		}
		result += "</li>"
	}
	result += "</ul>"
	return template.HTML(result)
}

func GetUtilsFuncMap() map[string]any {
	funcMap := funcs.GetUtilsFuncMap()
	funcMap["safeHTML"] = func(val string) template.HTML { // 承諾此數值是安全的，不需要額外的跳脫字元來輔助
		return template.HTML(val)
	}
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
		content := buf.String()

		// 建立toc物件
		var rootNode []*TOCNode
		{
			reToc := regexp.MustCompile(`(?m)^<h(\d)>(.*)<\/h\d>`)
			matchList := reToc.FindAllStringSubmatch(content, -1)
			var preNode *TOCNode
			for _, match := range matchList {
				depthStr, heading := match[1], match[2] // match[0]是所有匹配的項目，0之後才是每一個group的內容
				depth, err := strconv.Atoi(depthStr)
				if err != nil {
					PErr.Printf("error strconv.Atoi %s\n", err)
					return ""
				}
				curNode := &TOCNode{depth, heading, preNode, nil}
				if rootNode == nil {
					rootNode = make([]*TOCNode, 0)
					rootNode = append(rootNode, curNode)
					preNode = curNode
					continue
				}

				if preNode != nil && depth > preNode.Depth {
					if preNode.Children == nil {
						preNode.Children = make([]*TOCNode, 0)
					}
					preNode.Children = append(preNode.Children, curNode)
					preNode = curNode
					continue
				}

				// 往回找，直到前一個深度與它相等
				for {
					preNode = preNode.parent
					if preNode == nil {
						rootNode = append(rootNode, curNode)
						preNode = curNode
						break
					}
					if preNode.Depth < curNode.Depth {
						preNode.Children = append(preNode.Children, curNode)
						preNode = curNode
						break
					}
				}
			}
		}

		c := ctx.(*SiteContext)
		c.TableOfContents = renderToc(rootNode, "toc")
		return template.HTML(buf.String())
	}
	funcMap["debug"] = func(a ...any) string {
		log.Printf("%+v", a)
		return "" // fmt.Sprintf("%+v", a) // 只把訊息顯示在console，避免放到html之中
	}
	funcMap["timeStr"] = func(t time.Time) string {
		// t.Format("2006-01-02 15:04") // 到分感覺沒有意義
		return t.Format("2006-01-02")
	}

	funcMap["time"] = func(value string) (time.Time, error) {
		return time.Parse("2006-01-02", value)
	}

	funcMap["set"] = func(obj any, key string, val any) (string, error) {
		ps := reflect.ValueOf(obj)
		s := ps.Elem()
		if s.Kind() != reflect.Struct {
			log.Printf("type error. 'Struct' expected\n")
			return "", nil
		}
		field := s.FieldByName(key)
		if !field.IsValid() {
			log.Printf("key not found: %s\n", key)
			return "", nil
		}

		if !field.CanSet() {
			log.Printf("The field[%s] is unchangeable. You can't change it.\n", key)
			return "", nil
		}
		field.Set(reflect.ValueOf(val))
		return "", nil
	}
	return funcMap
}
