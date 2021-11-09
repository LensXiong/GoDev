package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

var count int32
var wg sync.WaitGroup //信号量
const ThreadNum = 1000

// goroutine1
func main() {
    // 1.信号
    wg.Add(ThreadNum)

    // 2. goroutine
    for i := 0; i < ThreadNum; i++ {
        go func() {
            //count++   // 2.1
            atomic.AddInt32(&count, 1)//2.1
            wg.Done() // 2.2
        }()
    }

    wg.Wait() // 3. 等待goroutine运行结束

    fmt.Println(count) // 4输出计数
}
