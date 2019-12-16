# Go_Goroutine
go 流程，Go 語言的亮點。

go 語言中，goroutine 相對應 go 敘述平行處理執行的匿名或命名函式。 go 敘述就是關鍵字 go 和函式的統稱，其實它沒有那麼神奇，只是在執行時期向系統遞交需要平行處理的執行任務而已。

執行時期系統收到此呼叫時，會先行檢查函式和參數的合法性，並且從本機 P （處理器）的自由 G 列表和排程器的自由 G 列表取得可用的 G，若沒有取得，則新增之。

# G (平行封裝) 的生命週期

每當我們呼叫 goroutine 時，執行時期系統透過 G 封裝 goroutine 前，會先對 G 初始化，初始化狀態就是 Grunnable 的生命週期階段。

                            Gidle

                            Grunnable

                            Grunning

                            Gsyscall

                            Gwaiting

                            Gdead

# Runtime Scheduler, 排程器

runtime.sched 代表 排程。

排程器所需要的功能是計時器與監測通知，所以其定義的欄位包含：

                            gcwaiting

                            stopwaiting

                            sysmonwait

                            sysmonnot
