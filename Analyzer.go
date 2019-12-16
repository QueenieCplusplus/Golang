// 2019, 11.19.14:00, by Queenie

import { "IDGenerator"
	 "ParseResponse"
	 "base"
}

// 分析器介面的宣告風格與網頁下載器的宣告風格很類似
type Analyzer interface {
	Id() uint32
	// 如下方程式為如上介面的簽名，也代表介面的功能
	Analyze(respParsers []ParseResponse, resp base.Response)([]base.Data, []error)
}

// 加強擴充性，如接收一重要參數，如下
type ParseResponse func(httpResp *http.Response, respDepth uint32)([]base.Data, []error)

// 實作分析器實例池的介面
// 建議實作於另一分頁
type AnalyzerPool interface {

	// 從實例池中取出一分析器
	Take()(Analyzer, error)
	
	// 將 idle 的分析器歸回給實例池子
	Return(analyzer Analyzer)error
	
	// 計算池子的總容量
	Total() uint32

	// 取得正在被使用的分析器數量
	Used() uint32

}

//TBD...(尚未完成)
