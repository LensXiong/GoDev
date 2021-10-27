package main

import "fmt"

type A struct {
    s string
}

// 逃逸场景一： 方法内返回局部变量指针。
// 局部变量原本应该在栈中分配，在栈中回收。但是由于返回时被外部引用，因此其生命周期大于栈，则溢出。

// 结果： go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:10:10: leaking param: s
// ./main.go:11:13: new(A) escapes to heap
// ./main.go:17:14: a.s + " world" does not escape
// ./main.go:18:12: b + "!" escapes to heap
// ./main.go:19:16: ... argument does not escape
// ./main.go:19:16: c escapes to heap

func foo(s string) *A {
    // new(A) escapes to heap
    a := new(A)
    a.s = s
    return a // 返回局部变量a
}
func main() {
    // new(A) escapes to heap
    a := foo("hello")
    // a.s + " world" does not escape
    b := a.s + " world"
    // b + "!" escapes to heap
    // b 变量没有逃逸，因为它只在方法内存在，会在方法结束时被回收。
    c := b + "!"
    // c escapes to heap
    fmt.Println(c) // hello world!
}
