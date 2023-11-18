package bytes

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

func GetFrontMatter[T any](src []byte, needRemain bool) (frontMatter *T, remain []byte, err error) {
	reader := bufio.NewReader(bytes.NewReader(src))

	// 有frontMatter的處理
	if len(src) >= 3 && string(src[:3]) == "---" {
		// https://go.dev/play/p/6cAAncvYKL2
		// https://github.com/CarsonSlovoka/go-make-gif/blob/6715c84eb4ae16289ee9ab4612e8c1aa87a870d1/src/make-gifex/main.go#L159-L184
		var (
			frontMatterFound bool
			writer           strings.Builder // strings.Builder可以有效地用於連續地組合和構建字串，特別是在需要動態添加字串內容時非常方便
			checkCount       = 0             // 確定有結尾的字段
			line             []byte

			part     []byte
			isPrefix bool
		)
		for {
			part, isPrefix, err = reader.ReadLine()
			if err != nil && err != io.EOF {
				return nil, nil, err
			}

			line = append(line, part...)

			if isPrefix { // 表示這一列，非常的長，沒辦法讀完，所以要繼續在讀，才是完整的資料
				continue
			}

			lineStr := string(line)

			if strings.TrimSpace(lineStr) == "---" {
				if frontMatterFound {
					// 遇到第二個 ---，代表結束，停止找尋
					checkCount++
					break
				} else {
					// 找到第一個 ---，開始記錄資料
					frontMatterFound = true
					checkCount++
				}
			} else if frontMatterFound {
				// 將內容寫入結果變數
				writer.WriteString(lineStr)
				writer.WriteString("\n") // 換行符號
			}

			line = nil // 清空緩存
		}

		// 表示有找到開始字段---和結尾字段---，即為一個完整的frontMatter
		if checkCount == 2 {
			if err = json.Unmarshal([]byte(writer.String()), &frontMatter); err != nil {
				return nil, nil, err
			}
		}
	}

	if needRemain { // 如果只是想要獲取到frontMatter的資訊，不需要取得到剩下的資料
		remain, err = io.ReadAll(reader) // 繼續讀取剩下的資料
	}

	return frontMatter, remain, nil
}
