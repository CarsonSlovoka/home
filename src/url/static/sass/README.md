## [sass change log](https://pub.dev/packages/sass/versions/1.50.1/changelog)

## [bootstrap專案](https://github.com/twbs/bootstrap)

- [檔案分類](https://github.com/twbs/bootstrap/tree/da54101/site/content/docs/5.2): 告訴您文件夾結構大致的分法
- [bootstrap/scss/](https://github.com/twbs/bootstrap/tree/da541014cb9caf830b77fa8993da0465b7210aac/scss) 此文件夾的scss就是組成bootstrap.css的內容
- [bootstrap/site/assets/scss/](https://github.com/twbs/bootstrap/tree/e12e080/site/assets/scss): 這個文件夾的scss就是組成github page (docs)所渲染出來的css

## SASS

### [Sass的種類](https://sass-lang.com/documentation/values/functions)

- [Dart Sass](https://sass-lang.com/dart-sass): 建議用這個版本，其他兩個Sass很多東西都有棄用的東西跑出來，LibSass甚至有很多問題。
- ~~[LibSass](https://sass-lang.com/libsass)~~
- ~~[Ruby Sass](https://sass-lang.com/ruby-sass)~~

### SASS基礎知識
- [sass-list-functions](https://kittygiraudel.com/2013/08/08/advanced-sass-list-functions/)
- [@use](https://sass-lang.com/documentation/at-rules/use) 這個會告訴您如何把文件分開寫
  - `@use "myDir/xxx" as *`: 直接嵌入 (方便，但名稱可能產生衝突)
  - `@use "myDir/xxx" as c`: 取別名
- 純引入用的文件名稱，開頭加上「`_`」
- `$-`[私有變數](https://sass-lang.com/documentation/at-rules/use#private-members), `$`公變數
- [index Files](https://sass-lang.com/documentation/at-rules/use#index-files) 對於每一個資料夾內的`_index.sass`這是一個特殊的檔案，允許您直接`@use '該資料夾名稱'`即可調用到該檔案，類似`init.go`
- [!default](https://sass-lang.com/documentation/variables#default-values) <sup>[實際範例參考](https://stackoverflow.com/a/72301550/9935654)</sup>這是sass特有的產物，當編譯到此變數時，如果該變數還是`null`或者沒有定義，就會使用當前`!default`的數值來設定它，否則就使用已經編譯的變數當成此數值
- `//` 單行註解不會產出在目標css內, `/* */` 多行註解則**會**
- `/* */` 結尾`*/`的欄位置不能和開始的欄位置相同
