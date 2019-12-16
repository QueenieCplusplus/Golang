// 2019.11.17.0900, by Queenie
// 2019.11.18.1000, by Queenie
// 2019.11.19.0730, by Queenie

// 排程器本身是中介軟體，能操縱處理各種模組
// 排程器的介面宣告和實現程式碼所在套件名稱如下
package sched "webcrawler/scheduler"

// define other fields, which is alias of the module, loacted in the same package
// 網路爬蟲軟體設計架構中
// 其他程式套件的匯入方式
// import alias "module name"
// import alias "module path"
import  ipl "webcrawler/itempipeline"
import dl "webcrawler/downloader"
import mdw "webcrawler/middleware"
import base "webcrawler/base"
import anlz "webcrawler/analyzer"
import "time" //請另行加入
import "requestCache" //請另行加入
import "sync/atomic" //請另行加入
import "errors" //請另行加入
import "net/http" //請另行加入
import "base" //請另行加入

// define fields for Sched
channelLen uint32,
poolSize uint32,
crawlDepth uint32 // 爬取深度，第一次請求深度為 0
primDomain string // 主域名
// 加入模組方法
chanman mdw.ChannelManager
stopSign mdw.StopSign
dlpool dl.PageDownloaderPool 
analyzerPool anlz.AnalyzerPool
itemPipeline ipl.ItemPipeline
// 讀寫 running 欄位值時，設計為最小操作
// 即便是最小操作，頻繁地使用讀寫，可能造成負面影響，故盡量少使用。
running uint32 // 執行標記, 0 not do, 1 done, 2 stop
reqCache requestCache,
urlMap map[string]bool // 已請求的 url map

// define constant
// 為每一個排程器具和每一個處理模組都定義一常數
const (
	SCHECULER_CODE = "scheduler"
	DOWNLOADER_CODE = "downloader"
	ANALYZER_CODE = "analyzer"
	ITEMPIPELINE_CODE = "item_pipeline"
)

// define Start() 
// 當開啟方法執行完畢後，排程器就真正處于執行狀態了！
// 然而不用擔心一直執行的狀態，因為如下很多函數在各種環節中會埋下回應停止訊號的程式碼。
Start (
	// define fields
	channelLen uint32,
	poolSize uint32,
	crawlDepth uint32 // 爬取深度，第一次請求深度為 0,
	httpClientGenerator GenHttpClient,
	respParser []anlz.ParseResponse,
	itemProcesseors []ipl.ProcessItem,
	firstHttpReq *http.Request
) (err error) {

	if firstHttpReq == nil {
		return errors.New("the 1st Http Request is invalid.")
	}
	pd, err := getPrimDomain(firstHttpReq.Host)
	if err != nil {
		return err
	}

	sched.primDomain = pd // 替優先主域名欄位設定值以後
	firstReq := base.NewRequestf(firstHttpReq, 0)
	sched.reqCache.put(firstReq) // 把第一筆請求的參數值放入緩存記憶中

}

Stop (

)() {

	// 0. 尚未執行, 1.正在執行, 2. 停止執行
	// 當發現 running 代表尚未執行，或是停止執行時
	// 代碼讓呼叫方停止操作，所以 return false
	// 讓呼叫方知道停止操作並沒有被執行
	if atomic.LoadUint32(&sched.running) != 1 {
		return false
	}
	// 倘若 running 處於執行，則可以呼叫停止方法
	sched.stopSign.Sign()
	// 另外，也設計將其他執行中的通道管理員和請求快取這些執行元件也做關閉的動作。
	sched.chanman.Close()
	sched.reqCache.close()

	// running 是最小操作，對它的值進行變更
	// 並將 true 作為方法結果傳回
	atomic.StoreUint32(&sched.running, 2)
	return true
}

func (sched *queensScheduler) Running() bool {
	return atomic.Loadint32(&sched.runnung) == 1
}

// define funcs
// 此 func body 免對如上欄位初始化
// NewScheduler 是方法名稱
// Scheduler 是轉換的介面型別
// body 陳述句為 轉換的結果是 女王排程序器型別指標
// func0
func NewScheduler() Scheduler {
	return &queensScheduler{}
}

// 真正在女王排程器開啟之後，欄位才有必要初始化
// 如下為匿名函式
// func1
defer func(){
	// 針對產生執行時期恐慌的狀況
	// 排程器能即時恢復及其對應的紀錄檔
	if p := recover(); p != nil {
		errMsg := fmt.Sprintf("fatal scheduler error occurs: %s\n, p")
		logger.Fatal(errMsg)
		err = errors.New(errMsg)
	}

	// 當排程器已經處於執行狀態時，則紀錄編號為 1，並且不應該再去回應後續 Start 方法的呼叫了
	// 編程設計上，則對尚未被開啟的排程序器被第一次呼叫之前，就要及時修改 running 這統計欄位的數值
	if atomic.LoadUint32(&sched.running) == 1 {
		return errors.New("the Scheduler has been started on ! \n")
	}
	atomic.StoreUint32(&sched.running, 1)

	// 如上 running 值設定為 1 之後，要對基礎資料型態參數均作數值檢查，如下
	// 一旦參數通過檢查，則指定給他對應的欄位。
	if channelLen == 0 {
		return errors.New("the Channel max length (capacity) shall over 0.\n")
	}
	sched.channelLen = channelLen

	if poolSize == 0 {
		return errors.New("the pool size shall over 0.\n")
	}
	sched.poolSize = poolSize
	sched.crawlDepth = crawlDepth

	// 參數指派到指定的欄位後，需要使用它們來初始化欄位。
	sched.chanman = generateChannelManager(sched.channelLen)

	if httpClientGenerator == nil {
		return errors.New("the Http Client Generator List is invalid.")
	}
	dlpool, err := 
		generatePageDownloaderPool(sched.poolSize, httpClientGenerator)
	
	if err != nil {
		errMsg := fmt.Sprintf("error occurs while getting pl pool: %s\n", err)
	    return errors.New(errMsg)
	} 
	sched.dlpool = dlpool

	analyzerPool, err := generateAnalyzerPool(sched.poolSize)
	if err != nil {
		errMsg := fmt.Sprintf("error occurs while getting analyzer pool: %s\n", err)
	    return errors.New(errMsg)
	}
	sched.analyzerPool = analyzerPool

	if itemProcesseors == nil {

	}
	// for in loop in Go
	for i, ip := range itemProcesseors {
		if ip == nil {
			return errors.New(fmt.Sprintf("the %d th item processor is invalid.", i))
		}

	}
	sched.itemPipeline = generateItemPipeline(itemProcesseors)

	if sched.stopSign == nil {
		sched.stopSign = mdw.NewStopSign()
	} else {
		sched.stopSign.Reset()
	}

	// 剩餘欄位分別為 reqCache urlMap primDomain 沒有被初始化
	// 則透過內建函數 make 為他們的值做出初始化即可
	sched.urlMap =  make(map[string]bool)

	// 執行環節
	sched.startDownloading() //下載時開啟通道
	sched.activateAnalyzers(respParsers) // 從回應通道接受回應並將其送給分析池中idle的分析器
	sched.openItemPipeline() // 使用通道中資料來啟動對應的處理模組，如上兩個執行行
	sched.schedule(10 * time.Millisecond) // 適切地請求快取中的請求向請求通道中傳送
}

// 開始下載方法既不接受參數，亦不回傳結果值
// 它的目的是：不斷地從請求通道中取得請求
// 並且交由空閒地網路下載器處理 dl, downloader
func (sched *queensScheduler) startDownloading() {
   go func(){
	   for {

		req, ok := <- sched.getReqChan()
		if !ok {

			break

		}
	   }
	   go sched.download()
   }
}

// 下載器
// 職責為下載與請求對應的網頁內容
// 透過某網頁下載器與目標網站進行互動
// 並且把互動結果（回應或是錯誤）發送予對應通道
// SCHECULER_CODE 為識別符號，為該方法所屬的排程器
// 該方法會呼叫網頁下載器的 Downoad 方法後，等待結果值
func download() {

	// 此 defer 述句目的是：拋出網頁下載器執行後可能造成的執行時期恐慌
	defer func() {
		if p := recover(); p != nil {
			errMsg := fmt.Sprintf("Fatal Download Error Occurs, which is %s\n", p)
			logger.Fatal(errMsg)
		}
	}()

	// Take() 方法目標是：取得 idle 的網頁下載器
	downloader, err := sched.dlpool.Take()
	if err != nil {
		errMsg := fmt.Sprintf("Downloader pool error occurs, which is %s", err)
		sched.sendError(errors.New(errMsg), SCHECULER_CODE)
	}

	defer func() {
		err := sched.dlpool.Return(downloader)
		if err != nil {
			errMsg := fmt.Sprintf("Downloader pool error occurs, which ias %s", err)
			sched.sendError(errors.New(errMsg), SCHECULER_CODE)
		}
	}()

	code := generateCode(DOWNLOADER_CODE, downloader.Id())
	respp, err := downloader.Download(req)
	if respp != nil {
		sched.sendResp(*respp, code)
	}
	if err != nil {
		sched.sendError(err, code)
	}
}

// 此函數目的是：從通道管理員處取得請求通道
func (sched *queensScheduler) getReqChan chan base.Request {
	reqChan, err := sched.chanman.reqChan()
	if err != nil {
		panic(err)
	}
	return reqChan
}

func (sched *queensScheduler) sendResp(resp base.Response, code string) bool {

	if sched.stopSign.Signed() {
		sched.stopSign.Deal(code)
		return false
	}

	sched.getRespChan() <- resp
	retrun true

}

func (sched *queensScheduler) sendError(err error, code string) bool {
	if err == nil {
		return false
	}

	codePrefix := parseCode(code)[0]
	var errorType base.ErrorType
	switch codePrefix {

	case DOWNLOADER_CODE:
		errorType = base.DOWNLOADER_ERROR
	
	case ANALYZER_CODE:
		errorType = base.ANALYZER_ERROR
	
	case ITEMPIPELINE_CODE:
		errorType = base.ITEM_PRPCESSOR_ERROR
	}

	// 變數 cError 代表了我們將要發送 base.CrawlerError型別的值給錯誤通道
	// 發送前，需要檢查停止訊號
	cError := base.NewCrawlerError(errorType, err.Error())
	if sched.stopSign.Singed() {
		sched.stopSign.Deal(code)
		return false
	}
	// 若停止訊號處於尚未發出的狀態，則可以放心同步地發送錯誤迅行
	go func() {
		sched.getErrorChan <- cError
	}
}

// 啟動分析器
// 此方程式會緊跟著 startDownloading 方法後呼叫
// 使用 go 常式執行所有其中的程式，他們將以盡可能最快的速度從回應通道接收回應
// 然後平行處理之
// 參數 respParser 代表意義：分析器所需的回應解析函數的序列
// resp 參數則代表：需要被分析的回應
// defer() 作用意義：做為當運行時如果對系統拋出錯誤，避免產生執行時期恐慌因而影響整體運行流程。
// defer 的設計算是設計架構中的防線。
func (sched *queensScheduler) activateAnalyzers (respParsers []anlz.ParseResponse） {

	code := generateCode(ANALYZER_CODE, analyzer.Id())
	dataList, errs := analyzer.Analyze(respParsers, resp)
	if dataList != nil {
		for _, data := range dataList {

			if data == nil {
				continue
			}

			switch d := data.(type) {

				case *base.Request:
					sched.saveReqToCache(*d, code)
				case *base.Item:
					sched.sendItem(*d, code)

				default: 
					errMsg := fmt.Sprintf()
					sched.sendError(errors.New(errMsg), code)

			}

		}

	}

	go func(){

		for{
			resp, ok := <- sched.getRespChan()
			if !ok {
				break
			}
			go sched.analyze(respParsers, resp)
		}
	}()

	// 從排程器中持有的分析器池中取出分析器
	// 此一操作可能會被 block 阻塞，直到池中有可用的分析器為止。
	// 此 defer 防線能確保分析回應之後即時將該分析器歸還到分析池
	defer func() {

		if p := recover(); p != nil {
			errMsg := fmt.Sprintf("Fatal Analysis Error occurs, which is %s\n",p)
			logger.Fatal(errMsg)

		}
	}()

	analyzer, err := sched.analyzerPool.Take()
	if err != nil {
		errMsg := fmt.Sprintf("Analyzer pool error occurs, which is %s", err)
		sched.sendError(errors.New(errMsg), SCHECULER_CODE)
		return
	}

	defer func() {
		err := sched.analyzerPool.Return(analyzer)
		if err != nil {
			errMsg := fmt.Sprintf("Analyzer pool error occurs, which is %s", err)
			sched.sendError(errors.New(errMsg), SCHECULER_CODE)
		}

	}()
	
}

// 將請求儲存到快取
// 若請求值（型別為 *http.Request）為 nil，則該請求非法(Invalid)。
// 若 http 請求的 URL 欄位(net/url 套件中 URL 型別)為 nil，則包含它的請求是非法的。
// 若 URL 的 Schema 欄位值非 http，則該請求非法。
// 為避免後續錯誤，把尚未支援的協定過濾掉
func (sched *queensScheduler) saveReqToCache (req base.Request code string) bool {

	httpReq := req.httpReq()
	if httpReq == nil {
		logger.Warnln("plz ignore the request, since the target http req is invalid.")
		return false
	}

	reqUrl := httpReq.URL
	if reqUrl == nil {
		logger.Warnln("plz ignore the request, since the target url is invalid.")
		return false
	}
	
	if string.ToLower(reqUrl.Schema) != "http" {
		logger.Warnf("plz ignore the request, since the target url schema '%s' is not http.\n", reqUrl.Schema)
		return false
	}

	// 倘若要求目標 url 已經被處理過，則設計它成為不合要求的。 
	// _, 為空白識別符號，相當於此值可以扔掉不用，此判別只需要用第二個值
	if _, ok := sched.urlMap[reqUrl.String()]; ok {
		logger.Warnf("plz ignore the request, since the target url (requestUrl= '%s') is repeated.", reqUrl)
		return false
	}

	// 判斷 Host name 是否與 PrimDomain 值相等
	// 倘若不相等，則判斷為不合要求
	// _ 為空白識別符號，用來被扔掉的值，因為在此判斷式中無需用到
	if pd, _ := getPrimDomain(httpReq.Host); pd != sched.primDomain {
		logger.Warnf("plz ignore the request, since the target host '%s' is the Prim Domain", 
		httpReq.Host, sched.primDomain, reqUrl)
		return false
	}

}

// 開啟項目處理管線
// 排程器僅有項目處理管線實例，而無需對應的 pool
// 藉此，有利紀錄與項目處理有關的 count 計數
func (sched *queensScheduler) openItemPipeline () {

	go func() {
		sched.itemPipeline.SetFailFast(true)
		code := ITEMPIPELINE_CODE
		for item := range sched.getItemChan() {
			go func(item base.item) {
				//...
			}(item)
		}
	}()

	// 建築防線
	defer func() {
		if p := recover(); p != nil {
			errMsg := fmt.Sprintf("fatal Item Processing Error occurs, which is %s", p)
			logger.Fatal(errMsg)
		}
	}()

	// 把傳入參數值的 item 即項目發送給項目處理管線
	// 檢查變數 errs 的值，在必要時刻將錯誤值一個個發送給錯誤通道
	errs := sched.itemPipeline.Send(item)
	if errs != nil {
		for _, err := range errs {
			sched.sendError(err, code)
		}
	}  
	
}

// 排程器將適切地將快取記憶體中的請求搬運到請求通道
// 需要使用 time 套件，因為要將搬運時間與計數器綁在一起
// 排程器的參數 interval 值會決定 for 程式區塊中城市的執行時間間隔
// 所以代碼 body 只需要做兩件事：繼續搬運以及即時回應停止訊號
// 適當有兩層面意義：
// 1. 根據請求通道中的 idle 位置決定搬運請求的數量
// 2. 根據快取記憶體中儲存的請求數量來決定搬運請求的數量
func (sched *queensScheduler) schedule (interval time.Duration) {
	go func() {

		for {

			//取得請求通道中的剩餘位置
			//通道的容量減去實際包含的元素值數量等於空位數量
			remainder := cap(sched.getReqChan()) - len(sched.getReqChan())

			//根據剩餘空位，循環地搬運快取中的請求，直到快取空了，方才停止搬運
			var temp *base.Request
			for remainder > 0 {
				temp = sched.reqCache.get()
				if temp == nil {
					break
				}
				// 實際搬運快取中的請求至請求通道中的剩餘位置
				sched.getReqChan() <- *temp
				remainder--
			}

			// 當搬運完成時，在指定間隔時間等待一會兒
			time.Sleep(interval)

		}

        // 此程式碼代表回應“停止訊號”的策略
		if sched.stopSign.Signed() {
			sched.stopSign.Deal(SCHECULER_CODE)
			return
		}
	}()
}

func (sched *queensScheduler) sendItem (/*TBD*/) {
	//TBD
}