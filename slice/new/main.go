package main

import "fmt"

// first argument to append must be slice; have *[]int
// 解析： 可以使用list:=make([]int,0) list类型为切片
// 或使用*list = append(*list, 1) list类型为指针
func main() {

    // new 和 make 的区别：
    // 二者都是内存的分配（堆上），但是make只用于slice、map以及channel的初始化（非零值）；
    // 而new用于类型的内存分配，并且内存置为零。
    //​ make返回的还是这三个引用类型本身；而new返回的是指向类型的指针
    list := new([]int)
    // list := make([]int,0)
    fmt.Println(list)  // &[]
    fmt.Println(*list) // []

    // *list = append(*list, 1)
    list = append(list, 1)

    fmt.Println(list)
}
