# Go_Slice

切片的資料結構

定義：
可以說是對陣列 [] 的一種包裝形式，切片是針對其底層陣列中某連續片段的描述符號。

效用：
切片類型為了實現針對其底層陣列中某連續片段的操作提供比陣列更為好用的介面，Gopher 使用切片的比例甚至比陣列來得高。

特色：
切片長度是可變的！

類型標記方式：

    // 元素類型為 T 的切片類型
    []T
    
    // 可將匿名結構類型作為切片類型的元素類型
    []struct{ name string, phoneNumber int}
    
值的標記方式：

    // 切片值的表示和陣列值的表現如出一徹。
    []string{"Go", "Java", "PHP", "C++", "C", "Python", "JS"}
  
// TBD...(尚未完成)


