//2019, 11/22, by Queenie
// map as a dictionary, including k/v pair

package main

import "fmt"

func main() {

	mp := map[uint]string{2: "Q", 0: "U", 1: "E", 8: "N"}
	var key uint
	for k := range mp {
		if k > key {
			key = k
			v := mp[k]
			fmt.Printf("%d: %s\n", k, v)
		}

	}

}
