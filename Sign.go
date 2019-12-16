// 2019.11.16.1300, by Queenie 

import "signal"
import "sync/RWmutex"

// Signal

// 表示訊號是否已經發出
signed bool

// 處理計數的字典
dealCountMap map[string]uint32 

// 讀寫鎖
rwmutex sync.RWmutex 

// 用來初始化 "停止訊號" 的函數
func NewStopSign() StopSign {

	ss := &queensStopSign {
		dealCountMap: make(map[string]uint32)
	}

	return ss

}

// 宣告 sign 的方法

func (ss *queensStopSign) Sign() bool {
	
	ss.rwmutex.Lock()
	defer ss.rwmutex.Unlock()

	if ss.signed {
		return false
	}

	ss.signed = true
	return true

}

// 無需加入前置和後置的檢查與操作
// 因此沒有互斥鎖和互斥執行的需求
func (ss *queensStopSign) Signed() bool {
	return ss.signed
}

// 處理訊號的函數，主要操作 dealCountMap 欄位
// 影響讀寫鎖
// 邏輯設計上：若停止訊號尚未發出，則後續操作將被忽略，因為處理計數不準確
func (ss *queensStopSign) Deal(code string) {
	ss.rwmutex.Lock()
	defer ss.rwmutex.Unlock()

	if !ss.signed {
		return
	}

	if _, ok := ss.dealCountMap[code]; !ok {

		ss.dealCountMap[code] = 1

	} else {

		ss.dealCountMap[code] += 1

	}
}

// 重置 "停止訊號"
// 讓已發出轉為尚未發出的狀態
// 捨棄 dealCountMap 欄位的所有值
// 綁定新值，利用歸零所有計數
func (ss *queensStopSign) Reset() {
	ss.rwmutex.Lock()
	defer ss.rwmutex.Unlock()

	ss.signed = false
	ss.dealCountMap = make(map[string]uint32)
}
