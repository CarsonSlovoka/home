---
{
  "title": "sass",
  "tags": [ "sass", "scss", "css" ],
  "layout": "blog/blog.base.gohtml",
  "cTime": "2024-08-30",
  "mTime": "2024-08-30"
}
---

# Sass

## Q&A

### @use與@forward有何不同

```scss
// _colors.scss
$primary: blue;
$secondary: green;

// _forward.scss
@forward 'colors';

// _use.scss
@use 'colors';

// main.scss
@use 'forward';
@use 'use';

.example {
  color: $primary; // 可以直接使用，因為它是通過 @forward 轉發的
  background-color: use.$secondary; // 需要使用命名空間，因為它是通過 @use 導入的
}
```

### `%`的作用

`%` 是一個佔位符，例如

```sass
// 定義佔位符選擇器
%base-button
  padding: 10px 15px
  border: none
  border-radius: 5px
  cursor: pointer

// 使用佔位符選擇器
.primary-button
  @extend %base-button
  background-color: blue
  color: white

.secondary-button
  @extend %base-button
  background-color: gray
  color: black

// 定義一個普通類（不使用佔位符）
.unused-class
  margin: 10px
```

output
```css
/* 用佔位符的部分會被單獨抓出來 */
.primary-button, .secondary-button {
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.primary-button {
  background-color: blue;
  color: white;
}

.secondary-button {
  background-color: gray;
  color: black;
}

.unused-class {
  margin: 10px;
}
```
