package html_test

import (
	"carson.io/pkg/html"
	"fmt"
)

func ExampleQuerySelector() {
	htmlContent := `
        <html>
        <head>
			<meta name="qoo" content="abc">
        </head>
        <body>
            <!-- 頁面內容 -->
        </body>
        </html>
    `

	htmlContent = html.InsertMetaTag(htmlContent, `<meta name="md5-hash" content="abc">`)

	elem := html.QuerySelector(htmlContent, "meta", func(elem *html.Element) bool {
		if elem.GetAttr("name") == "md5-hash" {
			return true
		}
		return false
	})
	if elem != nil {
		fmt.Println(elem.GetAttr("content"))
	}

	// Output:
	// abc
}
