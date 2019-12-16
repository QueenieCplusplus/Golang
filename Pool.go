// 2019.11.16 Noon, by Queenie 

import "fmt"
import "reflect"

// 使用通道 channel 作為池中實體的容器
container chan Entity 

// 池中實體的類型
etype reflect.Type 

// Golang 視 func() 為一等型別
// 池中實體產生的函數
getEntity func() Entity

// 此欄位忠實呈現實體池的容量
total uint32

// 如下為實體池的基礎結構
type queensPool struct {

	total uint32
	etype reflect.Type
	getEntity func() Entity
	container chan Entity

	// 初始化實體池
	func newPool(
		total uint32, 
		entityType reflect.Type,
		genEnity func() Entity) (Pool, error) {

			// 實體池容量
			if total == 0 {
				errMsg := fmt.Sprintf("the Pool can't be init! (total=%d)\n", total)
				return nil, errors.New(errMsg)
			}

			// 初始化實體池後，立刻填滿它
			size := int(total)
			container := make(chan Entity, size)

			for i := 0; i < size; i++ {
				newEntity := genEnity()

				if entityType != reflect.TypeOf(newEntity){

					errMsg := fmt.Sprintf("the Type of result of func called genEntity() is not %s\n", entityType)
					return nil, errors.New(errMsg)

				}

				container <- newEntity

			}

	}

}

// 製作實體池的實例
pool := &queensPool {

	total: total,
	etype: entityType,
	genEnity: genEnity,
	container: container
	
}

return pool, nil
 

