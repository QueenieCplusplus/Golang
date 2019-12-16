# Go_map
字典，鍵值對。

>>>
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
    
 >>>
 
     $go run <src file with extension name>
     
 >>>
 output
 
    ➜  desktop git:(master) ✗ go run Map.go
    1: E
    8: N
