# Go_Scheduler
排程器的實作

排程器也是中介軟體，掌管多執行緒，其最小操作稱為 running，具有處理調度所有模組的功能，並且與模組間有一道專門傳遞訊息的中介軟體，稱為通道管理員，而模組們都有多個實例 instance or entity，太多的關係，歸屬在池子中，稱為 Pool。

另外，排程器的資源釋放需要呼叫回應關閉訊號的函數，稱為 stopSign()。

本代碼需要先載入其他模組

    import mdw "webcrawler/middleware"
    import  ipl "webcrawler/itempipeline"
    import dl "webcrawler/downloader"
    import mdw "webcrawler/middleware"
    import base "webcrawler/base"
    import anlz "webcrawler/analyzer"
    
其模組代碼如下：

＊mdl/channelManager
https://github.com/poupougo/Go_ChannelManager

＊mdl/entityPool
https://github.com/poupougo/Go_EntityPool

＊mdl/stopSign
https://github.com/poupougo/Go_Sign/blob/master/Sign.go

＊downloader
https://github.com/poupougo/Go_PageDownloader

＊analyzer
https://github.com/poupougo/Go_Analyzer

＊itempipeline
https://github.com/poupougo/Go_Pipeline

