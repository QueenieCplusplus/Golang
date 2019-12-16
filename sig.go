//(1)資料型別宣告
type Signal interface {
    String() string
    Signal()
}

//(2)使用 os/signal 套件的 Notify 方法宣告方程式
func Notify(c chan<- os.Signal, sig ...os.Signal)

/*註解：
第一個參數表示只能向該 channel 放入 signal ，不能從該通道中取出訊號。 運用的約束符號 <- 就是代表如上意義，其中 c 代表通道類型值。
第二個參數代表著可變長的參數，參數 sig 代表的參數值包含我們希望自行 handle 的所有 signals。*/


//(3) 建立兩參數值，等下步驟 (4) 的 signal.Notify 方法會用到。
sigRecv := make(chan os.Signal, 1)
sigs := []os.Signal{syscall.SIGINIT, syscall.SIGQUIT}

//(4) 呼叫方法
signal.Notify(sigRecv, sigs...)

//(5) 迴圈輸出
for sig := range sigRec {
     fmt.Print("we have a signal now, it is %s\n", sig)
}