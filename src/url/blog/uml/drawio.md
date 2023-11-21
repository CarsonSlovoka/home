---
{
  "title": "DrawIO教學",
  "tags": [ "uml", "tutorial" ],
  "layout": "blog/blog.base.gohtml",
  "cTime": "2022-11-20T19:00:00+08:00"
}
---

# DrawIO教學

您可以直接上官網

> https://www.drawio.com/

點擊Start Now按鈕即可開始

它有提供一些保存方式

例如:

- google drive
- github
- gitlab
- ...
- local (本機)

如果你不想要讓它有訪問一些你私有項目的權限，例如Github, google drive等項目，

那就選擇用`本機`，選擇本機的時候，它會要求可以擁有`檔案編輯`的權限

這個你可以在chrome設定之中找到(隱私權和安全性)，你可以自由勾選`檔案編輯`看是要{詢問、允許、封鎖}等等

這種技術，其實是透過瀏覽器本身就能辦到，類似以下的javascript代碼

```html
<button>要求寫入文件的權限</button>
<script>
  // 當用戶點擊按鈕時觸發
  document.querySelector('button').addEventListener('click', async () => {
    // 請求用戶選擇一個文件
    try {
      const fileHandle = await window.showOpenFilePicker()
      // 現在我們有了文件的句柄，我們可以請求讀取或寫入權限
      if (await fileHandle[0].queryPermission({ mode: 'readwrite' }) !== 'granted') {
        const permissionGranted = await fileHandle[0].requestPermission({ mode: 'readwrite' })
        if (permissionGranted !== 'granted') {
          throw new Error('We do not have permission to edit this file.')
        }
      }

      // 讀取文件
      const file = await fileHandle[0].getFile() // file 屬於blob的一種

      const contents = await file.text()
      console.log(contents); // 現在您可以看到文件的內容

      // 假設我們想要寫入新的內容
      const writable = await fileHandle[0].createWritable()
      await writable.write('New contents')
      await writable.close()
    } catch (err) {
      console.error(err.message)
    }
  })
</script>
```

## 設定

在繪圖之中，可以對`頁面尺寸`進行調整，如果要列印，建議用成A4的尺寸(210*297mm)

## 常用操作

### 替換物件

直接抓取您要的物件，拖曳到想取代的物件即可

### 範本

可以善用旁邊的`範本`，把相關的物件拖曳進去，即可成為範本

範本可以導出，要點選`編輯`，之後可以看到有`導出`與`導入`的按鈕(都為xml格式)

如果您是用網頁版本，範本應該是用local storage等等相關的技術緩存下來，才會使得每次刷新，都還是能看到範本

但是以防萬一，建議還是要將您的範本導出，避免遺失。

### Container/Group

成為一個容器，您可以把相關的內容拖曳進去到容器之中，容器可以允許摺疊

### 圖層(layer)

快速鍵<kbd>Ctrl+Shift+L</kbd>可以開啟圖層

選中物件之後，點擊旁邊的`...`可以選擇這些物件要套用在哪些圖層

圖層排名最前面的項目，會優先顯示在頂層，例如

```
圖層-foo
圖層-bar
...
圖層N
```

假設`圖層-foo`的項目與`圖層-bar`有重疊，那麼`圖層-foo`會呈現在bar之上，因為它排名比較前面

### 標籤

<kbd>Ctrl+K</kbd>與圖層差不多，只是它是針對物件，如果你想對某些物件做隱藏或顯示的開關，這也是一種方式

## 快捷鍵

Edit相關

| 快速鍵                    | 說明 |
|------------------------| ---- |
| <kbd>Ctrl</kbd>        | 對著圖形按住Ctrl不放，即可複製
| <kbd>Ctrl+D</kbd>      | 複製同樣的內容在旁邊
| <kbd>Delete</kbd> | 刪除該物件(相關的連線不會刪除)
| <kbd>Ctrl+Delete</kbd> | 刪除物件(包含相關的連線等等都會一併刪除)

Canvas

| 快速鍵              | 說明 |
|------------------| ---- |
| <kbd>Ctrl+G</kbd> | Group 其實就是將該物件的屬性，變成Container(您可以在該物件的屬性中看到這個屬性會被打勾)

View

| 快速鍵              | 說明 |
|------------------| ---- |
| <kbd>Ctrl+J</kbd> | Fit page(可以觀看到完整一頁資料)
