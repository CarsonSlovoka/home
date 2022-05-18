## USAGE

- `help` 查看可用的指令(目前只寫了`build`而已)
- 可以透過`build`指令建立實體的HTML文件

  > build -f -o=..\\docs\\

  注意-o不要加上雙引號，例如`-o="..\\docs\\"`會發生錯誤

  或只使用build剩下的全用預設(`-f=false, -o=..\\docs\\`)
  > build
- 使用`run`以本機當成server，能直接瀏覽網頁查看結果
- `quit` : 離開程式

## 開發注意事項

- 所有的html附檔名都請改用gohtml

## 如何發佈到gh-pages

改完之後，可以善用

> `git reset [sha1-id]`

回到gh-pages，這樣當前的檔案不會被更動<sup>還是維持目前的狀態，而不是checkout過去時該分支當時的狀態</sup>

而在把docs都更新並且提交之後，也是再利用同樣命令回到開發分支
