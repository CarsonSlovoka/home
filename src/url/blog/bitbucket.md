---
{
  "title": "Bitbucket異動紀錄",
  "tags": [ "git", "hosting" ],
  "layout": "blog.base.gohtml",
  "meta": {
    "description": "..."
  },
  "draft": false,
  "date": "2022-09-09T00:00:00+08:00",
  "lastMod": "2022-09-09T00:00:00+08:00"
}
---


# Bitbucket

最近<sup>2022/04/02</sup>有了異動，可能會收到以下錯誤

```
Bitbucket Cloud recently stopped supporting account passwords for Git authentication.
remote: See our community post for more details: https://atlassian.community/t5/x/x/ba-p/1948231
remote: App passwords are recommended for most use cases and can be created in your Personal settings:
remote: https://bitbucket.org/account/settings/app-passwords/
fatal: Authentication failed for 'https://bitbucket.org/xxx/....git/
```

要至

> https://bitbucket.org/account/settings/app-passwords/

設定好權限(勾以下幾種應該就很足夠了)

```
Repositories
  Read
  Write
  Admin
  Delete
Pull requests
  Read
  Write
```
完成之後，他會彈出一個對話框，裡面的密碼可以複製起來。

這樣做的好處是，你可以在密碼中看到每個設定最後被**存取的時間點**，再加上通常使用者也不太會去記密碼，

因此如果忘記密碼就是重新設定，也就等同使用者**要常常更換密碼**，變相的增加安全性。
