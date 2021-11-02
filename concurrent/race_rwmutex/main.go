package main

import (
    "fmt"
    "sync"
    "time"
)

// 读写锁，这个锁读与读是不互斥的，简单来说就是这个锁可以保证同时只能有一个 goroutine 在写入。
// 如果有 goroutine 在写入，其他的 goroutine 既不能读，也不能写，但允许多个 goroutine 同时来读。
var mu sync.RWMutex

var balance int

func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    balance = balance + amount
}
func Balance() int {
    return balance
}

// 问题：向银行账户中存款问题。
// 解析：如果程序正确，那么最后的输出应该是 200000，但多次运行，结果可能是 198000、199000 或者其他的值。这个程序存在数据竞态。
// 这个问题的根本原因是 balance = balance + amount 这行代码在 CPU 上的执行操作不是原子的，有可能执行到一半的时候会被打断。

// 结果：200000
// 解决方案：保证同一时间只能有一个 goroutine 来访问变量。
// ① 互斥锁。sync.Mutex
// ② 读写互斥锁。sync.RWMutex
// ③ once。 &sync.Once{}

func main() {
    for i := 0; i < 1000; i++ {
        go func() {
            Deposit(100)
        }()

        go func() {
            Deposit(100)
        }()
    }
    // 休眠一秒，让上面的 goroutine 执行完成
    time.Sleep(1 * time.Second)
    fmt.Println(Balance())
}
