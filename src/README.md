## USAGE

- `help` 查看可用的指令(目前只寫了`build`而已)
- 可以透過`build`指令建立實體的HTML文件

  > build -f -o=..\\docs\\

  注意-o不要加上雙引號，例如`-o="..\\docs\\"`會發生錯誤

  或只使用build剩下的全用預設(`-f=false, -o=..\\docs\\`)
  > build
- 使用`run`以本機當成server，能直接瀏覽網頁查看結果
- `quit` : 離開程式

## 開發相關

### build簡介

複製[src/url](https://github.com/CarsonSlovoka/CarsonSlovoka.github.io/tree/fe36034/src/url)的內容<sup>會排除不需要的，如`sass`, `pkg`</sup>到輸出資料夾

其中遇到`gohtml`的檔案，會使用`text/template`進行渲染

> 以[src/url/tmpl](https://github.com/CarsonSlovoka/CarsonSlovoka.github.io/tree/fe36034/src/url/tmpl)<sup>含其所有**子**資料夾</sup>
> 當作`ParseFiles`的內容(當然該內容本身也必須涵蓋當前的目標檔案本身)

### 注意事項

- 所有的html附檔名都請改用gohtml

## 如何發佈到gh-pages

改完之後，可以善用

> `git reset [sha1-id]`

回到gh-pages，這樣當前的檔案不會被更動<sup>還是維持目前的狀態，而不是checkout過去時該分支當時的狀態</sup>

而在把docs都更新並且提交之後，也是再利用同樣命令回到開發分支
