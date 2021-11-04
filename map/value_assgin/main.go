package main

import "fmt"

type Student struct {
    Name string
}

var list map[string]Student

// 结果：cannot assign to struct field list["student"].Name in map
// 解析：初始化的 map 是一个引用类型，只读不可写。
func main() {

    list = make(map[string]Student)
    fmt.Println(list)            // map[] 返回引用类型本身
    fmt.Println(list["student"]) // {}

    student := Student{"wangxiong"}

    list["student"] = student
    list["student"].Name = "wwxiong"

    fmt.Println(list["student"]) // {wangxiong}
}
