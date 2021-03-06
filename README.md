# Golang
11/16-11/22, 2019

plz see code.

Intro https://github.com/QueenieCplusplus/Golang/blob/master/Intro.md (KickOff)

structure https://github.com/QueenieCplusplus/Golang/blob/master/main.md (專案結構)

lifecycle https://github.com/QueenieCplusplus/Golang/blob/master/main.go (main)

---------------------------------------------

# 循序的通訊

管線與佇列都屬於循序的通訊，前者屬於單向半雙工。

MQ IPC: TBD (非 go 語言特點)

pipeline IPC https://github.com/QueenieCplusplus/Golang/blob/master/Pipeline.go (單向半雙工-管線)

       兩設備或兩端點同時送出訊號 emit signal，但不同時傳送 transmit，僅允許切換方向，只允許單向傳輸，所以還是有點像是肉眼看不到差異的非同步(非字面上同步的概念)。

half-duplex https://github.com/QueenieCplusplus/Networking/blob/master/half_duplex.md (半雙工的解釋)

---------------------------------------------

# 非同步的通訊

(增加結果的不確定性與資源競奪上鎖）

* 同系統處理共用資源 share memory 的 IPC 使用 signal notify IPC
* duplex 雙向全雙工的不同軟體系統間的溝通 Socket

       Analog Wave -> Modem Encoder to Signal-> Serailizer -> transmitter IO
       
       
          Socket Open -----------> Message Queue -----------> Socket Open
          

       Analog Wave<- Modem Decoder to Signal -> Serializer <- receiver IO

signal 

               //os/signal的 Notify 方法
              func Notify(c chan<- os.Signal, sig ...os.Signal)
              
>>>

                           核心
                  
                            ｜

                          虛擬機
                  
                            ｜

                           進程                
                
                          / ｜ \
                          
                       -   IPC   -
 
                     常式 常式 常式 常式 ...
                   
>>>

   https://github.com/QueenieCplusplus/Networking/blob/master/sig.md

   https://github.com/QueenieCplusplus/Golang/blob/master/sig.go

   https://github.com/QueenieCplusplus/Golang/blob/master/Sign.go
   
   https://github.com/QueenieCplusplus/Golang/blob/master/sig.md

channel https://github.com/QueenieCplusplus/Golang/blob/master/ChannelManager.go (單向-通道)

pool https://github.com/QueenieCplusplus/Golang/blob/master/Pool.go (實體池)

parallel https://github.com/QueenieCplusplus/Golang/blob/master/Parallel.md (平行執行)

threads https://github.com/QueenieCplusplus/Golang/blob/master/MultiThread.md

goroutine https://github.com/QueenieCplusplus/Golang/blob/master/Go_Routine.md

效能問題 https://github.com/QueenieCplusplus/Optimization/blob/master/README.md

--------------------------------------------
# 雙向通訊

socket: 雙向通道 (全雙工)

---------------------------------------------

迴圈與資料結構

loop https://github.com/QueenieCplusplus/Golang/blob/master/Range.go (range)

https://github.com/QueenieCplusplus/Golang/blob/master/Range.md

https://github.com/QueenieCplusplus/Golang/blob/master/TypePointer.md (Type)

map https://github.com/QueenieCplusplus/Golang/blob/master/Map.go (雜湊表)

https://github.com/QueenieCplusplus/Golang/blob/master/Map.md

slice https://github.com/QueenieCplusplus/Golang/blob/master/Slice.md (切片資料)

rune TBD

fallthru TBD

逐位元互斥運算 TBD

---------------------------------------------

Crawler 

https://github.com/QueenieCplusplus/Golang/blob/master/Scheduler.md (專案模組)

https://github.com/QueenieCplusplus/Golang/blob/master/Scheduler.go (cpu分配器)

https://github.com/QueenieCplusplus/Golang/blob/master/PageDownloader.go (網頁下載器)

https://github.com/QueenieCplusplus/Golang/blob/master/Analyzer.go （內容解析器）

---------------------------------------------

專案打包

Package https://github.com/QueenieCplusplus/Golang/blob/master/Pkg.md

Module Manager https://github.com/QueenieCplusplus/Golang/blob/master/GoMod.md


