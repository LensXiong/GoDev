package main

import (
    "fmt"
    "sync"
)

var count int32
var wg sync.WaitGroup // 信号量
var lock sync.Mutex   // 互斥锁
const ThreadNum = 1000

// goroutine1
func main() {
    // 1.信号
    wg.Add(ThreadNum)

    // 2.goroutine
    for i := 0; i < ThreadNum; i++ {
        go func() {
            lock.Lock()   // 2.1
            count++       // 2.2
            lock.Unlock() // 2.3
            wg.Done()     // 2.4
        }()
    }

    wg.Wait() // 3.等待goroutine运行结束

    fmt.Println(count) // 4.输出计数
}
