# Go_PassData_or_ShareMemory

多程序執行時，會遇到程序彼此間傳遞資料或是共用資料的情形，而不同情境與條件下適用的通訊方式互異：

pipeline (管線)

signal (訊號)

socket (通訊端)

-----------------------------------------------------------------

# IPC passing Data thru same Program

單一程式碼做傳遞資料的行程間通訊

* Pipeline, 管線 (平行)

是種半雙工(單向)的通訊方式，僅用於祖父子孫關係的程序間的通訊，應用於 shell。

---------------------------

* MQ, 訊息佇列 (循序, 非同步)

請詳見同作者的 nodejs2019 此 repository.

Links: https://github.com/nodejs2019/message_queue （本作者親筆稿）

QMQ: http://zguide.zeromq.org/page:all#header-3 （官網)

-----------------------------------------------------------------

# IPC using Share Memory

共用資源的行程間通訊

* Signal, 訊號 (循序, 非同步)

它的本質是軟體模擬硬體的中斷 breakdown 機制，訊號目的是通知某程序遇到某事件發生了！例如

SIGINIT

SIGQUIT

SIGKILL

作業系統中，以上訊號都由正整數代表，即訊號的編號。
在 Linux 和 Unix 平台中，指令為如下，用以檢視支援的訊號指令：

    $kill -l
    
輸出結果：有可靠訊號與不可靠訊號，目前 Unix 僅支援不可靠訊號。
https://github.com/poupougo/Go_PassData_or_ShareMemory/blob/master/Signal.png
    
         1) SIGHUP	 2) SIGINT	 3) SIGQUIT	 4) SIGILL
         5) SIGTRAP	 6) SIGABRT	 7) SIGEMT	 8) SIGFPE
         9) SIGKILL	10) SIGBUS	11) SIGSEGV	12) SIGSYS
        13) SIGPIPE	14) SIGALRM	15) SIGTERM	16) SIGURG
        17) SIGSTOP	18) SIGTSTP	19) SIGCONT	20) SIGCHLD
        21) SIGTTIN	22) SIGTTOU	23) SIGIO	24) SIGXCPU
        25) SIGXFSZ	26) SIGVTALRM	27) SIGPROF	28) SIGWINCH
        29) SIGINFO	30) SIGUSR1	31) SIGUSR2

身為程式開發者，我們可以自訂當處理程序接收到這些訊號的反應和處理方式，需要執行某操作等等，可運用標準函式庫 os/signal API 處理之。

如下函數實作，將順道使用 go 的 channel 通道，很有趣喔～

https://github.com/poupougo/Go_PassData_or_ShareMemory/blob/master/sig.go

(1)資料型別宣告

    type Signal interface {

        String() string
        Signal()

    }
    
(2)使用 os/signal 套件的 Notify 方法宣告方程式

    func Notify(c chan<- os.Signal, sig ...os.Signal)
    
註解：

第一個參數表示只能向該 channel 放入 signal ，不能從該通道中取出訊號。
運用的約束符號 <- 就是代表如上意義，其中 c 代表通道類型值。

第二個參數代表著可變長的參數，參數 sig 代表的參數值包含我們希望自行 handle 的所有 signals。

(3) 建立兩參數值，等下步驟 (4) 的 signal.Notify 方法會用到。

    sigRecv := make(chan os.Signal, 1)
    sigs := []os.Signal{syscall.SIGINIT, syscall.SIGQUIT}
    
(4) 呼叫方法
    
    signal.Notify(sigRecv, sigs...)
    
(5) 迴圈輸出

    for sig := range sigRec {
         fmt.Print("we have a signal now, it is %s\n", sig)
    }
    
-----------------------------------------------------------------

# IPC passing Data via Internet

網際網路的行程間通訊

* Socket, 通訊端

屬於雙向全雙工的通訊方式，常見於登入免重複驗證的網路程式碼。

Links: https://github.com/nodejs2019/socket （本作者親筆稿）

https://github.com/nodejs2019/nodeApp_SocketIO (實作 )

-----------------------------------------------------------------









