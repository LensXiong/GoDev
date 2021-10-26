package main

import (
    "fmt"
    "math/rand"
    "sync"
)
// 编程题：写代码实现两个 goroutine ，最终输出五个随机数。
// 其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。

// 解析：
// ① goroutine 在 golang 中是非阻塞的。
// ② channel 无缓冲情况下，读写都是阻塞的，且可以用 for 循环来读取数据，当管道关闭后， for 退出。
// ③ golang 中有专用的 select case 语法从管道读取数据。

func main() {
    writeChannel := make(chan int)
    wg := sync.WaitGroup{}
    wg.Add(2)
    count := 5
    go func() {
        defer wg.Done()
        for i := 0; i < count; i++ {
            writeChannel <- rand.Intn(99)
        }
        close(writeChannel)
    }()

    go func() {
        defer wg.Done()
        for i := range writeChannel {
            fmt.Println("i=", i)
        }
    }()
    wg.Wait()
}
