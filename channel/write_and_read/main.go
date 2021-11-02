package main

import "fmt"

// 结果：fatal error: all goroutines are asleep - deadlock!

// 解析：不可以在同一个 goroutine 中既读又写，否则将会死锁。
func main() {
    ch := make(chan int)

    ch <- 2
    x := <-ch
    fmt.Println(x)
}


// 结果：
// after write
// after read: 2

// 解析：两个 goroutine 中使用无缓冲的channel，则读写互为阻塞。
// 即双方代码的执行都会阻塞在 <-ch 和 ch <- 处，直到双方读写完成在 ch 中的传递，各自继续向下执行。
func main1() {
    ch := make(chan int)

    go func() {
        ch <- 2
        fmt.Println("after write")
    }()

    x := <-ch
    fmt.Println("after read:", x)
}
