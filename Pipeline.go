// 2019.11.16.1500, by Queenie
// pipeline 管線實作
// 需要注入 ProcessItem 及其同一套件中的介面 ItemPipeline
// 需要注入 fmt, base, sync/atomic
import "ProcessItem"
import "ItemPipeline"
import {atomic} from "sync"


// 宣告定義型別如下
type queensPipeline struct {

	itemProcessors []ProcessItem // 此欄位代表此結構中持有許多項目處理器
	failFast bool // 表示處理是否需要快速失敗的標示位

	// 以下四個欄位是為滿足統計需要而生成
	processing uint64 // 正在處理項目的數量
	processed uint64 // 已經被處理項目的數量
	sent uint64 
	accepted uint64 

	// 此函數接受一個 []ProcessItem 類型的參數
	// 回傳一個 ItemPipline 類型的結果
	// 函數 body 檢查了輸入進來的參數值
	// panic() 為執行時期恐慌，請注意其用法
	// 請注意 range 用法
	func NewItemPipeline(itemProcessors []ProcessItem) ItemPipeline {

		if itemProcessors == nil {
			panic(errors.New(fmt.Sprintln("the item Processor List is Invalid.")))
		}
		
		innerItemProcessors := make([]ProcessItem, 0)
		for i, itpr := range itemProcessors {
			if itpr == nil {
				panic(errors.New(fmt.Sprintf("item processor [%d]\n", i)))
			}
			innerItemProcessors = append (innerItemProcessors, itpr)
		}

		return &queensPipeline{itemProcessors: innerItemProcessors}

	}
    
}

// 撰寫方法
// 此指標方法擁有功能：
// 1. 檢查 item 值的有效性，忽略無效值的處理
// 2. 依次呼叫 itemProcessors 欄位中的項目處理器對有效項目進行處理，並且依照 failFast 欄位值控制其流程。
// 3. 收集處理過程中發生錯誤，並且將對應的值作為結果回傳
// 4. 處理過程中，對應的欄位 sent accept processing processed 將分別紀錄，滿足統計需求。
// 函數 body 中對輸入值進行細節檢查
// 因為 base.item 型別容許 nil，故額外做處理
// 請注意  _, 用法
// atomic 其他用法
// 請注意 ⌃ 的用法
// 透過呼叫 sync/atomic 套件方法，保障遞增遞減的操作

func(itpr *queensPipeline) Send(item base.item) []error {

	atomic.AddUint64(&itpr.sent, 1)
	atomic.AddUint64(&itpr.processing, 1)

	errs := make([]error, 0)
	if item == nil {
		errs = append(errs, errors.New("the item is invalid."))
		return errs
	
	atomic.AddUint64(&itpr.accepted, 1)

	var currentItem base.Item = item

	defer atomic.AddUint64(＆itpr.processed, ⌃uint64(0))
	for _, itemProcessors := range itpr.itemProcessors {

		ProcessedItem, err := itemProcessors(currentItem)

		if err != nil {
			errs = append(errs, err)
			if itpr.failFast {
				break
			}
		}

		if processedItem != nil {
			currentItem = processedItem
		}

	}

	// 統計需要
	counts := make([]uint64, 3)
	counts[0] = atomic.LoadUint64(&itpr.sent)
	counts[1] = atomic.LoadUint64(&itpr.accepted)
	counts[2] = atomic.LoadUint64(&itpr.processed)
}

func(itpr * queensPipeline) Sum() string {

	var sum = "failFast: %v, processing: %d," + 
	"sent: %d, accepted: %d, processed: %d, processing: %d"

	counts :itpr.Count()

	sum := fmt.Sprintf(sum, itpr.failFast, len(itpr.itemProcessors), 
		counts[0], counts[1], counts[2], itpr.processing())

	return sum

}


	
