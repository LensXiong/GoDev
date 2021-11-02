package main

import (
    "fmt"
    "sync"
    "time"
)

var balance int

func Deposit(amount int) {
    balance = balance + amount
}
func Balance() int {
    return balance
}
// go run -race ./main.go
// 问题：向银行账户中存款问题。
// 解析：如果程序正确，那么最后的输出应该是 200000，但多次运行，结果可能是 198000、199000 或者其他的值。这个程序存在数据竞态。
// 这个问题的根本原因是 balance = balance + amount 这行代码在 CPU 上的执行操作不是原子的，有可能执行到一半的时候会被打断。

// 结果：200
// 解决方案：保证同一时间只能有一个 goroutine 来访问变量。
// ① 互斥锁。sync.Mutex
// ② 读写互斥锁。sync.RWMutex
// ③ once。 &sync.Once{}

func main() {
    once := &sync.Once{}
    for i := 0; i < 10; i++ {
        once.Do(func() {
            go func() {
                Deposit(100)
            }()

            go func() {
                Deposit(100)
            }()
        })
    }
    // 休眠一秒，让上面的 goroutine 执行完成
    time.Sleep(1 * time.Second)
    fmt.Println(Balance()) // 200
}
