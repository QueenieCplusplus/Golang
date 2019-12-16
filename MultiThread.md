# Go_Multi_Threads
單一核心內多執行緒的平行處理。多執行緒與多處理程序並不相同，請勿會錯意。在 Go 語言，這些角色和支援的環境，彼此有代號定義著。

# P, G, M plays their roles

* P, Processor 處理器也稱為本機

上與機器環境是一對一關係，下與程式碼是一對多關係。

* G, Goroutine 被 Go 常式封裝的程式碼

Goroutine 稱為 go 流程，請詳 Links: https://github.com/poupougo/Go_Goroutine

* M, Machine 機器（可能是虛擬機）

runtime system 代表每次執行時期，每一個 M 的 lifecycle 中，僅有一 KSE 核心空間提供支援。

# Relationship amongs P, G, M

配上核心 KSE ， 角色間的階層關係如下：

                        KSE
                      
                      (一對一)

                        M
                      
                      (一對一)

                        P
                      
                      (一對多)

                        G

由上清晰可見，G 終究對應回 M 還是一對一關係，所以需要對其所需要的 CPU 時間和 Memory 資源做排程管理，而 G 能否執行，跟 P 能否提供時脈與資源給 G，至關重要。

P 的狀態如下：

                            Pidle

                            Prunning

                            Psyscall

                            Pgcstop

                            Pdead











