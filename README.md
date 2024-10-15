# 构建自己的 Redis项目
## 项目环境
系统： Windows 11 家庭中文版（22H2）<br />
开发软件：Visual Studio Code 1.81.1<br />
Go语言版本：go1.21.0 windows/amd64<br />
## 项目使用说明
在搭建好Go语言框架的情况下按以下步骤运行该程序：

1. ### 运行‘redis_from_newbing.go’文件，等待其显示<br />
        DAP server listening at: 127.0.0.1:7208
        Listening to port 6379...
2. ### 运行PowerShell，在其中输入以下指令：
        telnet localhost 6379
3. ### 待弹出名为‘Telnet localhost’的PowerShell窗口，则表示连接成功
4. ### 连接成功后，则可以通过以下命令与服务器进行交互：
        PING
        ECHO xxxx
        SET name data
        GET name
        DEL name
        QUIT
5. ### 在和服务器进行交互时，若给予了正确形式的命令进行交互，则服务器会回复正确响应。但是若命令错误或者参数数量错误时，服务器会根据错误类型进行报错。
