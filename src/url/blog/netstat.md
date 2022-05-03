# netstat

> netstat -p tcp
查看Proto為TCP的類型

> netstat -o
可以顯示PID

> netstat -n
Foreign Address改用`numerical`(數字)顯示

> netstat -n INTERVAL
多久顯示一次

> netstat -b
會顯示是哪一個應用程式所用到

> netstat -b
列出所有訊息列表(包含TCP, UDP, LISTENING, ESTABLISHED)

----

> netstat -ano | findstr :1234 | findstr ESTABLISHED
找尋埠號為1234且狀態為ESTABLISHED

----

| Name | Desc |
| ---- | ---- |
Proto | 有TCP, UDP兩種
Local Address | 表示本基地址
Foreign Address | 遠端地址，表示正在和那些機器連線
PID | Process ID
State |

State 有幾種狀態
- LISTENING: 表示處於監聽狀態(使用中，只是正在等待消息傳入)
- SYN_SENT:
  當要求與某機器連線時，就會是這個狀態，如果成功之後的狀態就會改為ESTABLISHED，

  因此通常這種訊息會很少被看到，如果您發現列表中有很多這種訊息，

  1. 有可能是訪問的網站不存在，或者線路不好
  2. 掃描軟體掃秒一個網段的機器也會常常出現
  3. 中毒！ 例如中了"衝擊波"，病毒發作時會掃描其它機器，這樣會有很多SYN_SENT出現

- SYN_RECV:接收到一個要求連線的主動連線封包。
- TIME_WAIT:等候對方回應的狀態。
- ESTABLISHED:連線已建立完成的狀態。
- FIN_WAIT1:該插槽服務(socket)已中斷，該連線正在斷線當中。
- FIN_WAIT2:該連線已掛斷，但正在等待對方主機回應斷線確認的封包。
