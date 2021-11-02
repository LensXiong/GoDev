package main

import (
    "fmt"
    "sync"
)

// 问题：从已关闭的 channel 读取数据会怎么样？

// 解析：无论 channel 是有缓冲还是无缓冲的，从已经关闭的 channel 中读取数据是没有问题的，但读取完 channel 的数据后再读取会返回 false 和默认值0。

// 创建一个有缓冲的通道
// 运行结果：
// write
// close chan
// read
// 1
// true
// 2
// true
// 0
// false
// 0
// false

func main() {
    //	Channel: The channel's buffer is initialized with the specified
    //	buffer capacity. If zero, or the size is omitted, the channel is
    //	unbuffered.
    //  func make(t Type, size ...IntegerType) Type
    ch := make(chan int, 3)
    wg := sync.WaitGroup{}
    wg.Add(1)
    go func() {
        fmt.Println("write")
        for i := 1; i < 3; i++ { // 循环往 channel 中写
            ch <- i
        }
        close(ch) // 关闭 channel
        fmt.Println("close chan")
        wg.Done()
    }()
    wg.Wait()
    fmt.Println("read")
    for i := 1; i < 5; i++ { // 循环从channel中读
        v, ok := <-ch
        fmt.Println(v)
        fmt.Println(ok)
    }
}

// 创建一个没有缓冲的通道
// 运行结果：
// read
// write
// close chan
// 1
// true
// 0
// false
// 0
// false
//func main() {
//    ch := make(chan int)
//    go func() {
//        fmt.Println("write")
//        ch <- 1
//        close(ch)
//        fmt.Println("close chan")
//    }()
//
//    fmt.Println("read")
//    for i := 0; i < 3; i++ {
//        v, ok := <-ch
//        fmt.Println(v)
//        fmt.Println(ok)
//    }
//}
