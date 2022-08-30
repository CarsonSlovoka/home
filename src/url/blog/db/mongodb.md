# MongoDB

## Install

需要安裝:

- mongoD: server運行環境
- mongosh: shell環境
- Compass (可選): 這個提供良好的可視化結果，建議可以啟動`mongodb\bin\Install-Compass.ps1`(powershell)來安裝它

### mongoD安裝

至官網

> https://www.mongodb.com

在上方的頁籤找尋

```
Products -> Community Edtion -> Community Server
```

之後會來到此頁面
> https://www.mongodb.com/try/download/community

選擇`MongoDB Community Server`

在右方會有下拉選單，選擇您想要下載的版本以及平台

```yaml
Version: 6.0.1 # 498MB
Platform: Windows
Package: zip (不需要選msi)
```

下載完之後解壓，放到ProgramFile去

把裡面的bin資料夾加入到環境變數

- mongod: database server

如果您選擇的是6.0以上的版本，[它把mongo.exe移除改用mongosh來取代](https://www.mongodb.com/docs/manual/release-notes/6.0-compatibility/#legacy-mongo-shell-removed)
> The mongo shell is removed from MongoDB 6.0. The replacement is mongosh.

```yaml
...
DBException in initAndListen, terminating","attr":{"error":"NonExistentPath: Data directory C:\\data\\db\\ not found. Create the missing directory or specify another path using (1) the --dbpath command line option, or (2) by adding the 'storage.dbPath' option in the configuration file."}}
...
```

會告知您找不到該路徑`C:\\data\db`

所以要去新增這些資料夾給它

至此您已經安裝好了mongod

### Monghsh安裝

> https://www.mongodb.com/try/download/shell

```yaml
Version: 1.5.4 # 38.8MB
Platform: Windows 64-bit (8.1+)
Package: zip
```

同樣也把它裡面的bin資料夾加到環境變數

測試: 打開cmd，執行mongosh.exe

### Compass安裝

有兩種安裝方法:

1. 透過`mongodb\bin\Install-Compass.ps1`來進行安裝
2. 直接到官網下載 (推薦用這個方法)

#### Install-Compass.ps1

```yaml
打開powershell # 以系統管理員身分
  # 注意如果路徑有空白，記得用 "
cd "mongodb\bin" # mongoD解壓之後的bin資料夾
Install-Compass.ps1
```

如果它報錯，那麼可能是您設定的關係，

執行`Get-ExecutionPolicy`，查看當前您的Policy

```yaml
Get-ExecutionPolicy -List # 如果您需要查看所有Scope各個的Policy為何，則使用這個指令
Get-ExecutionPolicy
AllSigned # 那麼表示如果您要運行所有的ps1腳本，該腳本都必須經過簽屬，不然就會報錯

Set-ExecutionPolicy Restricted # 設定為限制模式，禁止所有腳本
Set-ExecutionPolicy RemoteSigned -Force # 可以改成這個，之後再改回, 其中的-Force可以避免再出現確認訊息
# Set-ExecutionPolicy AllSigned -Scope Process -Force # 您也可以使用-Scope來指定要變更的Scope
```

#### 官網下載Compass

> https://www.mongodb.com/try/download/compass

```yaml
Version: 1.32.6 (Stable)
Platform: Windows 64-bit(7+)
Package: exe
```

完成之後，雙擊這個執行檔就可以執行Compass了

在首次執行時為有一些隱私設定需要設定，建議只勾前兩項

- [x] Enable Automatic Updates
- [x] Enable Geographic Visualization
- [ ] Enable Crash Reports (傳送錯誤報告)
- [ ] Enable Usage Statistics (啟動分析(會匿名傳送您的使用習慣))
- [ ] Give Product Feedback

URI設定

您要先啟用mongod.exe 執行server之後他才抓的到

```yaml
mongodb://localhost:27017 # 區網
mongodb://127.0.0.1:27017 # 純本機
```

### 教學影片

- [msi版本](https://youtu.be/Ph1Z97X6xno)
- [zip版本](https://youtu.be/nI6brMJdO1o)

## USAGE 使用

不論您是要用compass來檢視或者要啟動shell都要先啟動server
> mongod.exe

### Shell

一定要先啟用server: `mongod.exe`

否則當您運行: `mongosh.exe`

會遇到:

> MongoNetworkError: connect ECONNREFUSED 127.0.0.1:27017

正常啟動會出現類似以下的訊息

```yaml
Current Mongosh Log ID: abcdefghijaskdfmadsf0fbf
Connecting to: mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.5.4 # port預設都是27017
Using MongoDB: 6.0.1
Using Mongosh: 1.5.4
```

如果您不用了，想要關閉server可以運行

> db.shutdownServer()

就可以把server關閉

## Driver

go對它有非常好的支持，可以參考以下網址

> https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.10.1/mongo

## [指令](https://www.mongodb.com/docs/manual/reference/mongo-shell/#basic-shell-javascript-operations)

| command                                               | desc | return message |
|-------------------------------------------------------|----| ---- |
show dbs                                              | 注意! 如果資料庫是空的，會看不見
use mydb                                              | 切換到mydb，如果不存在，會自動建立 | switched to db mydb 告訴您目前已經切換過去
db.books.insertOne({"name":"my mongodb book"})        | 新增一筆資料到collection名稱為books去(如果該collection不存在會自動建立) | WriteResult({"nInserted": 1}) # 告訴您結果: 有一筆資料被插入
db.myCollection.insertOne({"name":"my mongodb book"}) |
db.books.find()                                       | 列出collection名稱為books的所有document
db.collection.drop()                                  | 刪除collection
.exit                                                 | 退出shell
