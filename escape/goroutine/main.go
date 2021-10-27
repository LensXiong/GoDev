package main

import (
    "fmt"
    "runtime"
    "time"
)

// 结果：10 10 10 10 10 10 10 10 10 10

// 分析：go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:11:9: moved to heap: i
// ./main.go:12:12: func literal escapes to heap
// ./main.go:13:22: ... argument does not escape
// ./main.go:13:22: i escapes to heap
// ./main.go:13:26: " " escapes to heap

// 解析：
// 什么是内存逃逸？本该分配到栈上的变量，跑到了堆上，这就导致了内存逃逸。
// 如果变量从栈逃逸到堆，会怎样？栈是高地址到低地址，栈上的变量，函数结束后变量会跟着回收掉，不会有额外性能的开销。
// 变量从栈逃逸到堆上，如果要回收掉，需要进行 gc，那么 gc 一定会带来额外的性能开销。
// 编程语言不断优化 gc 算法，主要目的都是为了减少 gc 带来的额外性能开销，变量一旦逃逸会导致性能开销变大。

func foo() {
    runtime.GOMAXPROCS(1)
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Print(i, " ") // 10 10 10 10 10 10 10 10 10 10
        }()
        //go func(m int) {
        //   fmt.Print(m, " ") // 9 0 1 2 3 4 5 6 7 8
        //}(i)
    }
    runtime.Gosched()
    time.Sleep(time.Second)
}

func main() {
    foo()
}
