package main

import "fmt"

func main() {
    // 用 append 内置函数，可以对切片进行动态追加
    var slice3 []int = []int{100, 200, 300}
    // 通过append直接给slice3追加具体的元素
    slice3 = append(slice3, 400, 500, 600)
    fmt.Println("slice3", slice3) // 100, 200, 300,400, 500, 600

    //通过 append 将切片slice3追加给slice3
    slice3 = append(slice3, slice3...) // 100, 200, 300,400, 500, 600 100, 200, 300,400, 500, 600
    fmt.Println("slice3", slice3)
}
