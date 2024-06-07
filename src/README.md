## USAGE

- `help` 查看[可用的指令](https://github.com/CarsonSlovoka/CarsonSlovoka.github.io/blob/071430aa50868fb8010bdd6948abcabdef0a92c1/src/cmd.go#L66-L145)
- 可以透過`build`指令建立實體的HTML文件

  ```yaml
  build -f -o=..\\docs\\ # -f表示目錄存在還是會強制建立
  build -f -o=..\\docs\\ -forceBuildAll # 不進行比較，強制建立所有東西
  ```

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

- [所有的html附檔名都請改用gohtml](https://github.com/CarsonSlovoka/CarsonSlovoka.github.io/blob/461cbcc4889d35542f222eec54102c7b5992e373/src/main.go#L188-L194)

  這是因為超連結的設定也是使用正常的html，如果把html只當作一般頁面，src依然是寫.html，它就會不曉得到底是該.gohtml還是.html

## 如何發佈到gh-pages

改完之後，可以善用

> `git reset [sha1-id]`

回到gh-pages，這樣當前的檔案不會被更動<sup>還是維持目前的狀態，而不是checkout過去時該分支當時的狀態</sup>

而在把docs都更新並且提交之後，也是再利用同樣命令回到開發分支
