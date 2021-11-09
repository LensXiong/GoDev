package main

import (
    "fmt"
    "sync"
)

// 结果：fatal error: all goroutines are asleep - deadlock!

// 使用 vet 工具发现 Mutex 复制使用问题。
// go vet ./main.go
// # command-line-arguments
// ./main.go:31:9: call of foo copies lock value: command-line-arguments.Counter
// ./main.go:35:12: foo passes lock by value: command-line-arguments.Counter

// 解析：Mutex 是一个有状态的对象，它的 state 字段记录这个锁的状态。
// 如果你要复制一个已经加锁的 Mutex 给一个新的变量，那么新的刚初始化的变量居然被加锁了，这显然不符合你的期望。
// 因为你期望的是一个零值的 Mutex。关键是在并发环境下，你根本不知道要复制的 Mutex 状态是什么。
// 因为要复制的 Mutex 是由其它 goroutine 并发访问的，状态可能总是在变化。

type Counter struct {
    sync.Mutex
    Count int
}

func main() {
    var c Counter
    c.Lock()
    defer c.Unlock()
    c.Count++
    foo(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter) {
    c.Lock()
    defer c.Unlock()
    fmt.Println("in foo")
}
