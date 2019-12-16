// 2019, 11/19.13:30, by Queenie
// PageDownLoader 
// the func called Download(type base.Request)(*base.Response, error)
// above metioned function shows the PageDownLoader's signature & functionality

import IdGenerator // 補實作，筆者將近期提供源碼

type PageDownloader interface {

	Id() uint32
	Download(base.Request) (*base.Response, error)

}

type PageDownloaderPool interface {

	// 從實例池子中取出一個網頁下載器
	Take() (PageDownLoader, error)

	// 把使用完畢的網頁下載器歸還給實例池子
	Return(dl PageDownloader) error

	// 獲得實例池的總容量
	// 總容量是可變的，所以不同時刻呼叫該方法，得到的數值可能都不相同。
	// 故，最好設計池子是固定容量的，策略上較佳。
	Total() uint32

	// 獲得正在被使用的網頁下載器數量
	Used() uint32

}


// type IdGenerator interface {
// 	// 獲得一個型別為 uint32 的 ID
// // 此 Id 有 2 的 23次方種變化，約莫 42 億種不同的數字組合
// 	GetUint32() uint32
// }

//TBD...(尚未完成)
