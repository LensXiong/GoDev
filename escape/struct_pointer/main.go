package main

import "fmt"

// 结果：# command-line-arguments
// ./main.go:22:10: leaking param: s
// ./main.go:23:13: new(bar) escapes to heap
// ./main.go:24:11: func literal does not escape
// ./main.go:15:16: ... argument does not escape

// 解析：Go 可以返回局部变量指针，这其实是一个典型的变量逃逸案例。
// 虽然在函数 foo() 内部 f 为局部变量，其值通过函数返回值返回，f 本身为一指针，其指向的内存地址不会是栈而是堆，这就是典型的逃逸案例。

func main() {
    f := foo("golang")
    fmt.Println(f)
}

type bar struct {
    s string
}

func foo(s string) *bar {
    f := new(bar) // 这里的new(bar)会不会发生逃逸？
    defer func() {
        f = nil
    }()
    f.s = s
    return f
}
