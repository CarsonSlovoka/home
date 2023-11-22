package html

import (
	"bufio"
	"strings"
)

// QuerySelector 不是很完美的設計，但對於簡單的內容足以
// https://go.dev/play/p/l8SgFzmFp4o
func QuerySelector(htmlContent, tag string, filters ...func(elem *Element) bool) *Element {
	scanner := bufio.NewScanner(strings.NewReader(htmlContent))
	for scanner.Scan() {
		line := scanner.Text()

		elem := ParseElement(line)
		if elem == nil {
			continue
		}

		if elem.Tag != tag {
			continue
		}

		for _, filter := range filters {
			if filter(elem) {
				return elem
			}
		}
	}
	return nil
}

type Element struct {
	Tag string
	raw string
	// Parent *Element
	// Child  []*Element
}

func ParseElement(raw string) *Element {
	if len(raw) < 2 {
		return nil // 沒有開頭的'<'和結尾的'>'
	}

	// if raw[len(raw)-2:] != "/>" { // 缺點，對於<br>的類型不管用
	if raw[len(raw)-1:] != ">" {
		return nil
	}

	raw = strings.TrimSpace(raw)

	var tag string
	if firstBlankIdx := strings.Index(raw, " "); firstBlankIdx > 1 {
		tag = raw[1:firstBlankIdx] // <meta ...> 只取meta
	} else {
		tag = raw[1 : len(raw)-1] // <html>
	}

	return &Element{
		tag,
		raw, // 資料有可能會是<meta xxx></head>
	}
}

func (e *Element) GetAttr(name string) string {
	parts := strings.Split(e.raw, name+`="`) // <meta content=" />
	if len(parts) > 1 {
		endIdx := strings.Index(parts[1], `"`)
		if endIdx != -1 {
			return parts[1][:endIdx]
		}
	}
	return ""
}

func InsertMetaTag(htmlContent, metaTag string) string {
	headEndIdx := strings.Index(htmlContent, "</head>")
	if headEndIdx != -1 {
		return htmlContent[:headEndIdx] + metaTag + htmlContent[headEndIdx:]
	}
	return htmlContent
}
