# DOIS 上传项目逻辑说明

## 对应类型

- 1:护理白板
- 2:床旁屏
- 3:门旁屏
- 4:护士站主机
- 5:走廊屏
- 6:安卓白板
- 7:移动护理
- 8:移动查房
- 9:输液监控
- 10:医生白板
- 101:启动器
- 102:WebApp
- 103:音视频
- 104:人脸识别
- 105:语音转文字
- 106:WS推送
- 107:浏览器内核
- 108:门旁屏（安卓）
- 109:走廊屏（安卓）
- 110:叫号大屏（安卓）
- 111:叫号门旁屏（安卓）
- 112:输液监控（安卓）
- 113:移动查房（安卓）
- 114:移动护理（安卓）
- 201:嵌入Web
- 301: 后台app
- 302：大屏app
- 401:移动推车

## zip类型

- "1", "3", "5", "6", "7", "8", "9", "10", "201", "401"
- 先将对应类型的目录内容删除，zip文件直接解压缩到对应目录


## zip(exe)类型

- "301", "302"
- 先将对应类型的目录内容删除，zip文件直接解压缩到对应目录，执行gomd5.exe，生成有文件md5值的version.json文件

## apk类型

"2", "4", "101", "102", "103", "104", "110", "111", "112", "113", "114"
这些类型为一些基础apk，主apk可能需要用到，放在/upgrade/android包下，并且需要循环遍历android包下的所有文件夹目录找到app.json。如果app.json里面存在对应的名称，如call，则将call对应的md5值进行修改保存
"103":"call",
"104", "face",
"105", "tts",
"106", "websocket"
"107", "xwalk"
其他主apk则放在/upgrade/android下对应的文件夹，并重命名为app.apk，并修改当前文件夹下的app.json里面名为app的md5值

### 对应的包路径

- 1:/upgrade/windows/com.chindeo.nwb
- 2:/upgrade/android/com.chindeo.bed.app
- 3:/upgrade/web/com.chindeo.bfmpp
- 4:/upgrade/android/com.chindeo.nursehost
- 5:/upgrade/web/com.chindeo.zlp
- 6:/upgrade/web/com.chindeo.nwb
- 7:/upgrade/web/com.chindeo.ydhl
- 8:/upgrade/web/com.chindeo.ydcf
- 9:/upgrade/web/com.chindeo.syjk
- 10:/upgrade/windows/com.chindeo.dwb
- 11:/upgrade/web/com.chindeo.bfmpp-s
- 101:/upgrade/android/com.chindeo.launcher.app
- 102:/upgrade/android/com.chindeo.webapp
- 103:/upgrade/android
- 104:/upgrade/android
- 105:/upgrade/android
- 106:/upgrade/android
- 107:/upgrade/android
- 108:/upgrade/android/com.ktcp.launcher.bfmpp
- 109:/upgrade/android/com.ktcp.launcher.zlp
- 110:/upgrade/android/com.ktcp.launcher.jhdp
- 111:/upgrade/android/com.ktcp.launcher.jhmpp
- 112:/upgrade/android/com.ktcp.launcher.syjk
- 113:/upgrade/android/com.ktcp.launcher.ydcf
- 114:/upgrade/android/com.ktcp.launcher.ydhl
- 201:/html/multipage
- 301:/upgrade/windows/Admin
- 302:/upgrade/windows/App
- 401:/upgrade/windows/com.chindeo.wrv

### 推送消息更新

#### 消息体

```json
{"category":"upgrade","msgType":"text","sendType":1,"sendUserType":"0","toUser":"nis","msg":""}
```

##### 哪个更改过，apk基础包会发送多个，推送给对应的主题

app/nis // 护理白板
app/bis // 床旁屏
app/wrv // 移动推车
app/nws // 护士站主机
app/dps // 门旁屏
app/wcs // 走廊屏
app/bdps // 叫号大屏
app/cdps // 叫号门旁屏
app/mcs // 移动护理
app/mwrs // 移动查房
app/tms // 输液监控
app/dwb // 医生白板

