# discord

## [搜尋](https://support.discord.com/hc/en-us/articles/115000468588)

建議熟悉discord之後可以把介面改成英文，這樣關鍵字才可以用英文去打，否則關鍵字需要打本地化的結果才會有作用

- `從: Foo#5776 之後: 2022-06-01 之前: 2022-06-09 有: 連結`
- `from: Foo#5776 after: 2022-06-01 before: 2022-06-09 has: link`

## [WebHooks](https://support.discord.com/hc/zh-tw/articles/228383668-%E4%BD%BF%E7%94%A8%E7%B6%B2%E7%B5%A1%E9%89%A4%E6%89%8B-Webhooks-)

它允許您當發生什麼事情時，傳送通知。

例如您可以到discord的伺服器設定中找到Webhook，接著產生，設定要在哪一個channel發送消息，接著複製webhook的網址(`https://discord.com/api/webhooks/<xxxID>/<token>`)

然後您可以找有支援webhooks的服務器，把該網址填入，舉Github為例:

> https://github.com/{username}/{repoName}/settings/hooks
>
> (username和repoName會成自己的)
>
>
> Payload URL: https://discord.com/api/webhooks/{xxxID}/{token}/github   (❗注意: 最後要加上`/github`)
>
> content type: application/json

就可以設定webhook，可以指派是所有動作都要發送通知還是只有push的時候才傳送通知等等

選擇好之後可以隨便push東西上去github，您到該channel應該就會看到有推播了

----

它的原理白話來說當您做了該動作，它就post請求，發送到該網址去，使得應用程式收到該消息，再呈現出來。


## [開發人員](https://discord.com/developers/applications)

### OAuth2

OAuth是一個開發標準([Open Standard](https://en.wikipedia.org/wiki/Open_standard))用來處理有關「授權」(Authorization)讓**應用程式**(Application)代表**資源擁有者**(Resource Owner)能訪問其資源。

它有別於傳統只靠帳密來驗證，OAuth更加的靈活，

它給的是Token，您可以透過更換它，來阻擋應用程式再繼續訪問，

同時也能再詳細的指派，該令牌能訪問的能力(有全部的權限、還是只能做否些事情而已)

您如果有興趣，可以讀一下[RFC6749](https://datatracker.ietf.org/doc/html/rfc6749)全部內容不多，也可以只看前半部了解一些名詞就行了

#### 名詞介紹

為了方便理解，我們舉例，想讓某個應用程式(AP)能具有訪問該使用者discord資料的能力

| Name | Desc  |
| ---- |-------|
Resource Owner | 資源的擁有者----使用者
Client | 想要取得受保護資料的「應用程式」
Authorization Server | 驗證Resource Owner的身分，並能分發Access Token給應用程式(Client)的伺服器。對應例子中的Discord，他會要求你輸入帳號密碼來核對您的身分，並且您可以透過它來產生Access Token
Resource Server | 存放使用者資料的伺服器。 以例子為例是指: Discord用哪一個伺服器來存放您的資料 (不一定是自己，也可以再委外出去)
Access Token | 應用程式要拿這一個Token，去向Resource Server取得同意被使用的資料
Authorization Grant | 同意應用程式去做某些事、取得一個範圍內的資源。
Redirect URI(Callback URL) | 在驗證完使用者身份並獲得授權同意後，把使用者導向哪一個URL(通常是應用程式的某一個路徑)
Scopes | 在discord中，有兩樣能勾選. 1. `bot` 2. `applications.commands` 您也許有些納悶，兩個有什麼差，bot沒意外是要選的(不然您要怎麼產生機器人)，至於applications.commands指的是能不能用`/`來指派要做什麼事情，例如您可以自訂`/`某樣東西要跑出怎樣的清單出來

#### 設定

- General: 這個是只通用的設定。
- URL Generator: 客製化。您可以不理會General的設定，直接來到此頁面，勾選您

## FAQ

- [如何更換channel的圖標](https://support.discord.com/hc/en-us/community/posts/360040862772-Channel-Icons) :
  1. 支付25美元成為開發人員(一次,永久)
  2. 如果您的伺服器變成公開的(community) 也可以獲得一些圖標(Announcements 📢, rules 📜、 Stage)
- ID種類:
  - Server(Guild)
  - Channel
  - Message (每一個訊息都有其專屬的ID)
- How to find Discord IDs
  - Go to User Settings > Advanced > enable Developer Mode.
  - Next, simply right-click your profile picture and select `Copy ID` to copy your User ID.
  - To find a Server(Guild), Channel, or Message ID, right-click on the server/ channel name, or message, and select `Copy ID`.
- 訊息是永久保存嗎: 是
- 我要怎麼刪除訊息
  - 點擊個別的訊息，或者按下Shift再點刪除，可以不會出現是否要刪除地確認框 (在早期有批次刪除的選項，但後來好像對discord伺服器造成太大的負擔，所以刪除了)
  - 如果您想刪除某頻道的所有訊息:
    1. 找處理訊息相關的機器人，用機器人幫你刪除
    2. clone頻道，再把舊的刪除
