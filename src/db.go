package main

import (
	bytes2 "carson.io/pkg/bytes"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type DB struct {
	// Contents 用來表示每篇blog檔案(md)的主要訊息
	Contents []*Content
}

var db DB

type Content struct {
	Filepath string
	FrontMatter
}

type FrontMatter struct {
	Title  string
	Tags   []string
	Layout string // 使用`tmpl/底下的哪一個gohtml檔案`當作layout

	// todo <meta>
	Meta struct {
		Keywords    string // <meta name="keywords" content="..."> // 目前大多的瀏覽器不再使用它來排名，但還是可以用來定義頁面的關鍵詞
		Description string // <meta name="description" content="...">

		// Robots // <meta name="robots" content="noindex, nofollow"> index vs noindex: 希不希望搜尋引擎索引此頁面; follow vs nofollow: 要不要進一步的檢索該頁面所連出去連結
		Robots struct {
			NoIndex  bool // true=>搜尋引擎會檢索此頁面
			NoFollow bool // true=>會進一步的檢索頁面所連出去的連結
		}
	}

	// Draft todo: 如果為草稿, build的時候，就不會導出此頁面
	Draft bool

	CreateTime  time.Time `json:"date"`
	LastModTime time.Time `json:"lastMod"`
}

func init() {
	targetDir := "./url/blog"
	err := filepath.Walk(targetDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil { // <-- 這個要補上，如果targetDir這個路徑錯誤，這個err會是該錯誤
			panic(err)
		}

		if info.IsDir() {
			/*
				// if info.Name() == "api" { // 如果不想要他掃描某些目錄可以用這招, 不過這種方法可能要注意 {xxx/api, api/} 這兩個都算這種情況
				if path == "xxx/api" { // 建議用這個比較明確
					return filepath.SkipDir // 當錯誤是SkipDir他會跳過，所以之後的檢查還是會繼續 https://github.com/golang/go/blob/8fb9565832e6dbacaaa057ffabc251a9341f8d23/src/path/filepath/path.go#L495-L510
				}
			*/
			return nil
		}

		for _, suffix := range []string{".md"} {
			if filepath.Ext(path) == suffix {
				var (
					b      []byte
					frontM *FrontMatter
				)
				b, err = os.ReadFile(path)
				if err != nil {
					panic(err)
				}
				frontM, _, err = bytes2.GetFrontMatter[FrontMatter](b, false)
				if err != nil {
					panic(err)
				}
				if frontM == nil {
					return nil
				}
				content := &Content{
					path,
					*frontM,
				}
				db.Contents = append(db.Contents, content)
				if err != nil {
					panic(err)
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
