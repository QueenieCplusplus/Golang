// 2019, 11/22, by Queenie
// for range
package main

import "fmt"

var i, d int

func main() {

	ints := []int{1, 2, 3, 4, 5}
	for i, d := range ints {
		fmt.Printf("%d: %d\n", i, d)
	}

}
