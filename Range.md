# Go_range = Iterator

for 述句可以攜帶一個 range 子句，進一步對陣列或是切片的值中每個元素，抑或是對字串中每個字元，也可對字典值中每個 k/v pair 鍵值對，做反覆運算。

適用情境

       []T
       
       String Char
       
       {
         k/v,
         k1/v1,
         k2/v2
       }

>>>
src file

    package main

    import "fmt"

    var i, d int

    func main() {

      ints := []int{1, 2, 3, 4, 5}
      for i, d := range ints {
        fmt.Printf("%d: %d\n", i, d)
      }

    }
    
>>>
execute cli $go run <src file> in terminal
    
    $cd desktop
    $ls
    $go run <src file name with extension name>
    ➜  desktop git:(master) ✗ go run Range.go
    0: 1
    1: 2
    2: 3
    3: 4
    4: 5



