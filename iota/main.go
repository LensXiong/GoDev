package main

import "fmt"

// 解析：iota 是 golang 语言的常量计数器，只能在常量的表达式中使用。
// iota 在 const 关键字出现时将被重置为0(const 内部的第一行之前)。
// const 中每新增一行常量声明将使 iota 计数一次( iota 可理解为 const 语句块中的行索引)。
const (
    a     = iota
    b     = iota
    name1 = "wangxiong1"
    c     = iota
    d     = iota
)
const (
    name2 = "wangxiong2"
    e     = iota
    f     = iota
)

// 0 1 3 4 1 2
func main() {
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
    fmt.Println(f)
}
