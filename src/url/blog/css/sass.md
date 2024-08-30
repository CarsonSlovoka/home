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
