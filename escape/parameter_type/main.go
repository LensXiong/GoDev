package main

import "fmt"

// 结果：# # command-line-arguments
// ./main.go:23:10: leaking param: s
// ./main.go:24:13: new(bar) does not escape
// ./main.go:25:11: func literal does not escape
// ./main.go:16:16: ... argument does not escape
// ./main.go:16:16: f escapes to heap

// 解析：这里的new(bar)会不会发生逃逸？其实没有发生逃逸。
// 而f escapes to heap 的逃逸是因为动态类型逃逸，
// fmt.Println(a …interface{})在编译期间很难确定其参数的具体类型，也能产生逃逸。

func main() {
    f := foo("golang")
    fmt.Println(f) // 这里的f是否会发生逃逸
}

type bar struct {
    s string
}

func foo(s string) bar {
    f := new(bar) // 这里的new(bar)会不会发生逃逸？
    defer func() {
        f = nil
    }()
    f.s = s
    return *f
}
