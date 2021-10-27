package main

import (
    "errors"
    "fmt"
)

// 结果：
// <nil>
// defer2 error
// <nil>

func f1() (r int) {
    defer func(r *int) {
        *r = *r + 5
    }(&r)
    return 1
}

func defer1() {
    var err error
    defer fmt.Println(err)
    err = errors.New("defer1 error")
    return
}

func defer2() {
    var err error
    // 闭包对 err 的引用
    defer func() {
        fmt.Println(err)
    }()
    err = errors.New("defer2 error")
    return
}

func defer3() {
    var err error
    defer func(err error) {
        fmt.Println(err)
    }(err)
    err = errors.New("defer3 error")
    return
}

func main() {
    defer1()
    defer2()
    defer3()
}
