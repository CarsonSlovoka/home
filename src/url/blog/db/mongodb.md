# MongoDB

## Install

需要安裝:

- mongoD.exe: server運行環境
- mongosh.exe: shell環境
- mongodb-Compass.exe (可選): 這個提供良好的可視化結果，建議可以啟動`mongodb\bin\Install-Compass.ps1`(powershell)來安裝它

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

### MongoSh安裝

> https://www.mongodb.com/try/download/shell

```yaml
Version: 1.5.4 # 38.8MB
Platform: Windows 64-bit (8.1+)
Package: zip
```

同樣也把它裡面的bin資料夾加到環境變數

測試: 打開cmd，執行mongosh.exe

### mongodb-Compass安裝

有兩種安裝方法:

1. ~~透過`mongodb\bin\Install-Compass.ps1`來進行安裝~~ 不建議
2. 直接到官網下載 (推薦用這個方法)

#### ~~Install-Compass.ps1~~

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

### 環境變數

我會建議先新增一個屬於MongoDB的環境變數之後再把此變數加到Path去，而不是直接開啟Path增加

這樣日後如果有要異動只需要更改MongoDB的變數內容即可

```yaml
# SETX 可以永久更改
# /M表示寫在Machine之中 ，有多個路徑可以用;隔開
SETX /M Mongo "%ProgramFiles%\Mongo;%programFiles%\Mongo\mongodb-win32-x86_64-windows-6.0.1\bin;%ProgramFiles%\Mongo\mongosh-1.5.4-win32-x64\bin;"

  # 放置變數到PATH之後
SETX /M PATH "%PATH%;%Mongo%"
# 如果太長會被截段，建議還是手動加)
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

### 進階用法

首次您使用，他預設都是在本機127.0.0.1，所以你可以先利用單機模式去設定admin等等相關用戶的權限，

完成之後再開始把`bindIp`加上去，再重啟動。

實際上您還可以傳遞一些參數給mongoD.exe，如下:

> mongod.exe --dbpath "C:\\data\db" --logpath "C:\data\db\log.log" --port 30000 --bind_ip 127.0.0.1 --directoryperdb

指令還有很多，我建議寫在設定檔之後只要執行:

> mongod.exe --config "C:\xxx\mongo.conf"

設定檔的內容

```yaml
systemLog:
  destination: file
  path: "C:\\data\\db\\log.log"
  logAppend: true
storage:
  dbPath: "C:\\data\\db"
  directoryPerDB: true
net:
  port: 27017
  bindIp: 127.0.0.1, 192.123.123.101 # 有多個ip可以用,隔開就好
processManagement:
  windowsService:
    serviceName: MongoDB27017
    displayName: MongoDB27017
```

您也可以只設定部分資訊就好

```yaml
mongod.exe --config "C:\xxx\mongo.yaml" # 或者您也可以把設定檔放在cmd一打開的路徑，這樣就不需要再額外寫path給它
```

```yaml
net:
  port: 17823
  bindIp: 127.0.0.1, 192.168.17.1 # 有多個ip可以用,隔開  # 放 "乙太網路卡 乙太網路" 的ipv4
security: # 等同 mongod.exe --auth
  authorization: enabled # 有了這個就必須要用身分登入才行，您可以use到指定的db再建立user，指令類似右方: db.createUser({user: "guest", pwd: "guest", roles: ["read"]})
processManagement:
  windowsService:
    serviceName: CarsonDB17823
    displayName: CarsonDB
```

## Driver

go對它有非常好的支持，可以參考以下網址

> https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.10.1/mongo

## [指令](https://www.mongodb.com/docs/manual/reference/mongo-shell/#basic-shell-javascript-operations)

如果您是在compass中使用mongoSh, 打上關鍵字，他其實就會提是有您可能想輸入的內容

| command                                               | desc | return message |
|-------------------------------------------------------|----| ---- |
| show dbs                                              | 注意! 如果資料庫是空的，會看不見
| show users                                            | 查看當前的db有哪些使用者
| use mydb                                              | 切換到mydb，如果不存在，會自動建立 | switched to db mydb 告訴您目前已經切換過去
| db.books.insertOne({"name":"my mongodb book"})        | 新增一筆資料到collection名稱為books去(如果該collection不存在會自動建立) | WriteResult({"nInserted": 1}) # 告訴您結果: 有一筆資料被插入
| db.myCollection.insertOne({"name":"my mongodb book"}) |
| db.books.find()                                       | 列出collection名稱為books的所有document
| db.collection.drop()                                  | 刪除collection
| db.getRoles()                                         |
| show roles                                            |
| show collections                                      |
| db.createRole()                                       |
| .exit                                                 | 退出shell

## 新建使用者

- [Manage Users and Roles](https://www.mongodb.com/docs/manual/tutorial/manage-users-and-roles/)
- [createUser](https://www.mongodb.com/docs/manual/reference/method/db.createUser/)

使用mongosh.exe

```yaml
show users # 查看當前的db有哪些使用者
```

```js
// use admin
db.createUser(
  {
    user: "yourUserName",
    pwd: "yourPassword",
    roles: [{role: "yourRoleName", db: "yourDatabaseName"}]
  }
)
```

範例:

```js
db.createUser(
  {
    user: "Carson xxx",
    pwd: "xxx",
    roles: [
      {role: "dbAdminAnyDatabase", db: "admin"},
      "readWriteAnyDatabase" // 當您只寫一個其實就代表 [role: "readWriteAnyDatabase", db: "<當前您use的db名稱>"]
    ]
  }
)

// 更多範例
// use admin // 把用戶資料儲存在admin
db.createUser(
  {
    user: "Carson",
    pwd: passwordPrompt(),
    // userAdminAnyDatabase: 可以看到db
    // dbAdminAnyDatabase: 可以看到db.document的名稱，但不能訪問
    roles: ["root"] // 給管理員超級權限
  }
)
```

```js
db.createUser(
  {
    user: "guest",
    pwd: "guest",
    roles: [
      {role: "read", db: "ocr"}
    ]
  })
```
> 除了建議超級用戶外，不建議直接用createUser，建議先建立想要的role之後再用此role去創建

移除

```js
db.dropUser("guest")
```

### [role](https://www.mongodb.com/docs/manual/reference/built-in-roles/#all-database-roles)

- Read：只能讀
- readWrite：讀、寫都可
- dbAdmin：允許用戶在指定數據庫中執行管理函數，如索引創建、刪除，查看統計或訪問system.profile **不能調整role**
- userAdmin：允許用戶向system.users集合寫入，可以找指定數據庫裡創建、刪除和**管理用戶**(調整role)
- dbOwner:
- clusterAdmin：只在admin數據庫中可用，賦予用戶所有分片和複製集相關函數的管理權限。
- clusterManager

----

- readAnyDatabase：只在admin數據庫中可用，賦予用戶所有數據庫的讀權限
- readWriteAnyDatabase：只在admin數據庫中可用，賦予用戶所有數據庫的讀寫權限
- userAdminAnyDatabase：只在admin數據庫中可用，賦予用戶所有數據庫的userAdmin權限
- dbAdminAnyDatabase：只在admin數據庫中可用，賦予用戶所有數據庫的dbAdmin權限。

----

- root：只在admin數據庫中可用。超級賬號，超級權限

```js
db.adminCommand({
  createRole: "myClusterwideAdmin",
  privileges: [
    {resource: {db: "config", collection: ""}, actions: ["find", "update", "insert", "remove"]}, // targetDB: config, target: collection: 空白表示 所有
    {resource: {db: "users", collection: "usersCollection"}, actions: ["update", "insert", "remove"]},
    {resource: {db: "", collection: ""}, actions: ["find"]}
  ],
  roles: [ // 表示此role會繼承的項目
    {role: "read", db: "admin"} // parent role name, parent db name
  ],
  writeConcern: {w: "majority", wtimeout: 5000}
})
```

刪除role
```js
db.runCommand({
  dropRole: "myClusterwideAdmin",
})
```

創建腳色使用自定義的role
```js
db.createUser({user: "test" , pwd: "123", roles: [  "myClusterwideAdmin" ]})
```

## 使用身分進行連線 (Connection)

> mongodb://username:password@123.123.123.123:27017/myDBName
