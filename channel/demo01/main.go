package main

import (
    "fmt"
    "runtime"
    "time"
)

// 结果： 一段时间后总是输出 #goroutines: 2
// 解析：因为 ch 未初始化，写和读都会阻塞，之后被第一个协程重新赋值，导致写的 ch 都阻塞。
func main() {
    var ch chan int
    go func() {
        ch = make(chan int, 1)
        ch <- 1
    }()
    go func(ch chan int) {
        time.Sleep(time.Second)
        <-ch
    }(ch)
    c := time.Tick(1 * time.Second)
    for range c {
        fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
    }
}
