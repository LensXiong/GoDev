package main

import (
    "fmt"
    "reflect"
)

// 结果：invalid operation: sn1 == sn3 (mismatched types struct { age int; name string } and struct { name string; age int })
// invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)

// 解析：结构体比较规则
// ① 只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
// ② 结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice，则结构体不能用 == 比较，可以使用 reflect.DeepEqual 进行比较。
func main() {

    sn1 := struct {
        age  int
        name string
    }{age: 11, name: "qq"}

    sn2 := struct {
        age  int
        name string
    }{age: 11, name: "qq"}

    sn3 := struct {
        name string
        age  int
    }{age: 11, name: "qq"}

    if sn1 == sn2 {
        fmt.Println("sn1 == sn2")
    }

    // 结构体比较与属性的顺序有关
    if sn1 == sn3 {
        fmt.Println("sn1 == sn3")
    }

    sm1 := struct {
        age int
        m   map[string]string
    }{age: 11, m: map[string]string{"a": "1"}}

    sm2 := struct {
        age int
        m   map[string]string
    }{age: 11, m: map[string]string{"a": "1"}}

    // 结构体中的 map 需要使用 reflect.DeepEqual 进行比较
    if sm1 == sm2 {
        fmt.Println("sm1 == sm2")
    }

    if reflect.DeepEqual(sm1, sm2) {
        fmt.Println("sm1 == sm2")
    } else {
        fmt.Println("sm1 != sm2")
    }
}
