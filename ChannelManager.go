// 2019.11.16.0830, by Queenie.

// 需要注入 time 和 fmt 套件。

import "time"
import "fmt"

// 實作通道管理員。

var chan1 chan int
var chanLength int = 18
var interval time.duration = 1500 * time.millisecond

chan1 = make(chan int, chanLength)

go func() {

	for i := 0; i < chanLength, i++ {

		if > 0 && i % 3 == 0 {
			fmt.Println("reset channel.")
			chan1 = make(chan1, chanLength)
		}
		
		fmt.Println("send element %d ...\n", i)
		chan1 <- i 
		time.Sleep(interval)

	}

	fmt.Println("Closing Channel one.")
	close(chan1)

}
