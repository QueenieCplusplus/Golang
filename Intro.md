# Go_RoadMap
{ Golang 總覽 }

1. Parallel & Concurrency, 『併行處理』 

    https://github.com/poupougo/Go_Parallel （理論）

2. Multithread & Mutex, 『多執行緒』

    https://github.com/poupougo/Go_Multi_Threads （理論）

3. go routine, 『go流程』 或稱 『go常式』

    https://github.com/poupougo/Go_Goroutine （理論）

4. Signal 『訊號』、Pipeline 『管線』、Socket 『通訊端』

    https://github.com/poupougo/Go_PassData_or_ShareMemory （理論）

5. Channel, 『通道』 （共享資源）

    https://github.com/poupougo/Go_ChannelManager (實作)

6. Signal, 『訊號』 (簡稱 Sign)

    https://github.com/poupougo/Go_Sign （實作）
    
7. Pipeline, 『管線』

    https://github.com/poupougo/Go_Pipeline （實作）

8. Scheduler 『排程器』

    https://github.com/poupougo/Go_Scheduler （實作）
    
9. EntityPool, 『實體池』 

    https://github.com/poupougo/Go_EntityPool （實作）
    
    
10. * 類型指標

     https://github.com/poupougo/Go_TypePointer

11. "*" (multicard) (語法)

    指標相關資料請詳：https://m.facebook.com/story.php?story_fbid=3660483613965420&substory_index=0&id=2730376213642836&ref=bookmarks&_ft_=mf_story_key.3660483613965420%3Atop_level_post_id.3660483613965420%3Atl_objid.3660483613965420%3Acontent_owner_id_new.2730376213642836%3Athrowback_story_fbid.3660483613965420%3Apage_id.2730376213642836%3Aphoto_id.3660483613965420%3Astory_location.4%3Astory_attachment_style.photo%3Apage_insights.%7B%222730376213642836%22%3A%7B%22page_id%22%3A2730376213642836%2C%22actor_id%22%3A2730376213642836%2C%22dm%22%3A%7B%22isShare%22%3A0%2C%22originalPostOwnerID%22%3A0%7D%2C%22psn%22%3A%22EntPhotoNodeBasedEdgeStory%22%2C%22post_context%22%3A%7B%22object_fbtype%22%3A22%2C%22publish_time%22%3A1574079870%2C%22story_name%22%3A%22EntPhotoNodeBasedEdgeStory%22%2C%22story_fbid%22%3A%5B3660483613965420%5D%7D%2C%22role%22%3A1%2C%22sl%22%3A4%2C%22targets%22%3A%5B%7B%22actor_id%22%3A2730376213642836%2C%22page_id%22%3A2730376213642836%2C%22post_id%22%3A3660483613965420%2C%22role%22%3A1%2C%22share_id%22%3A0%7D%5D%7D%7D&__tn__=%2As%2As-R

    因為 Golang 將 func 設定為一級階層，所以呼叫時可以搭配的指標符號 pointer。

12. "&" (and per se and) (語法)

    因為 Golang 將 func 設定為一級階層，所以呼叫時可以搭配的指標符號 pointer。
    
    項目 11 與 12 的語法範例如下：

	    // 導入工具包 別名 “工具包相對路徑”
	    package sched "webcrawler/scheduler"

	    // 方法名稱呼叫新的排程器 回傳排程器型別的皇后排程器位址裝載的物件
	    func NewScheduler() Scheduler {
		    return &queensScheduler{}
	    }

	    // 如下為尚須定義的方法呼叫名稱為 schedule() 其方法後為輸入參數 
	    // 方法前為 sched *queensScheduler
	    // 表示方法型別為一指標函數
	    // 皇后排程器的型別為排程器或其別名
	    // 皇后排程器左上為指標符號，是把可以開啟&queensScheduler{}的鑰匙。
	    func (sched *queensScheduler) schedule (/*TBD*/) {
		//TBD
	    } 
	    
	

13.  ^ 符號 (語法)

   https://github.com/poupougo/Go_bit_wise_ops

14. for range (語法)

   https://github.com/poupougo/Go_range/blob/master/README.md

    範例：
    
		go func() {
			sched.itemPipeline.SetFailFast(true)
			code := ITEMPIPELINE_CODE
			for item := range sched.getItemChan() {
				go func(item base.item) {
					//...
				}(item)
			}
		}()
		
     另一範例：
     
		// 把傳入參數值的 item 即項目發送給項目處理管線
		// 檢查變數 errs 的值，在必要時刻將錯誤值一個個發送給錯誤通道
		errs := sched.itemPipeline.Send(item)
		if errs != nil {
			for _, err := range errs {
				sched.sendError(err, code)
			}
		}  

15. ... (語法)

    （TBD）
    
16. ; (語法)

    常用在判斷式 if 子句，代表 &&，亦可省略。

    範例：
    
		// 建築防線
		defer func() {
			if p := recover(); p != nil {
				errMsg := fmt.Sprintf("fatal Item Processing Error occurs, which is %s", p)
				logger.Fatal(errMsg)
			}
		}()
	
17. 方程式的表示方式

    func(){}()
    
    func()(){}
    
    範例：
    
            // 轉換型別的過程
	    // assign 脈絡為 pointer 的值 -> *int 這指標類型的值 -> 變數容器 vPointer
	    vptr := (*int)(pointer)
    
    另一範例：
    
	    type Analyzer interface {
	    
			Id() uint32
		
			// 如下方程式為如上介面的簽名，也代表介面的功能
			Analyze(respParsers []ParseResponse, resp base.Response)([]base.Data, []error)
		
	      }
	      
18. _, 隱匿欄位 (語法)

    _, 為空白識別符號，相當於此值可以扔掉不用，此判別只需要其它參數值。
    
    範例：
    
        // 倘若要求目標 url 已經被處理過，則設計它成為不合要求的。 
        // _, 為空白識別符號，相當於此值可以扔掉不用，此判別只需要用第二個值
        if _, ok := sched.urlMap[reqUrl.String()]; ok {
            logger.Warnf("plz ignore the request, since the target url (requestUrl=%s) is repeated.", reqUrl)
            return false
        }
    另一範例：
    
        // 判斷 Host name 是否與 PrimDomain 值相等
        // 倘若不相等，則判斷為不合要求
        // _ 為空白識別符號，用來被扔掉的值，因為在此判斷式中無需用到
        if pd, _ := getPrimDomain(httpReq.Host); pd != sched.primDomain {
            logger.Warnf("plz ignore the request, since the target host '%s' is the Prim Domain", 
            httpReq.Host, sched.primDomain, reqUrl)
            return false
        }


# Grammer

a. go 是併發的關鍵字，放置於方程式命名之前。

b. defer 與 js 的 await 或是 kotlin 的 lateinit 相同，可節省 cpu 資源避免浪費。

c. goto 為 jump to，即跳過此 block。

d. chan 為通道的型別。

e. select 與其他程式語言的 switch 相同，body 通常搭配 case 。

f. running 為最小操作，其範例如下 

	StopScheduler()(){
	
		if atomic.LoadUint32(&sched.running) != 1 {
		return false
	        }
	        //...
		// running 是最小操作，對它的值進行變更
		// 並將 true 作為方法結果傳回
		atomic.StoreUint32(&sched.running, 2)
		return true
		
	}

	func (sched *queensScheduler) Running() bool {
	
		return atomic.Loadint32(&sched.runnung) == 1
		
	}
	
g. gothrough 控制流程的其中一個特殊的關鍵字用法

   https://github.com/poupougo/Go_fallthrough
	
h. Slice 切片資料結構的解析

   https://github.com/poupougo/Go_Slice

i. GOPATH|GOMOD 工具包管理工具的使用

   https://github.com/poupougo/Go_GoPath

# Package

 (1）time => 在同步的語言中，用計數器搭配在旁紀錄其 log 是明智的設計。

 (2) fmt => 與輸出有關，可以 print len 也可以 print function
 
 (3) sync/atomic => 與統計載入當下的值至相應欄位有關。
