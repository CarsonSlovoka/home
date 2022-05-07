# iphone作業系統更新

## 關閉自動備份

1. 查找itunes.exe的位置: 打開powershell，輸入
    > gcm itunes.exe
2. 關閉自動備份:
   這邊的範例假設位於ProgramFiles下，可自行替換成您的iTunes.exe路徑

    ```
    %ProgramFiles%\iTunes\iTunes.exe /setPrefInt DeviceBackupsDisabled 1
    ```

    如果要再開啟備份，把後面的1改成0即可。

    正常輸入完畢之後沒有任何訊息，以不會啟動iTunes

## iTunes未知錯誤4000

如果您發現更新itunes的時候，它會要求你輸入密碼，然後輸入完畢之後itunes又自動關閉，您要再次點擊啟動

或者遇到未知錯誤4000時，就可以考慮「`停用解鎖密碼`」，停用的方法如下

```
設定 -> FaceID與密碼/TochID與密碼 -> 關閉密碼
```

等更新完畢之後再啟用即可(指紋可能要重壓)

## itunes[下載檔案的位置](https://apple.stackexchange.com/a/49406/403361)

更新檔通常都好幾GB，而且也不會自動砍掉，有需要了話要自己去以下位子刪除

> %userprofile%\AppData\Local\Packages\AppleInc.iTunes_nzyj5cx40ttqa\LocalCache\Roaming\Apple Computer\iTunes\iPhone Software Updates
