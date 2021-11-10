package main

import (
    "fmt"
    "sync"
    "time"
)

// 重入导致死锁
// 运行结果：fatal error: all goroutines are asleep - deadlock!

// 当一个 writer 请求锁的时候，如果已经有一些活跃的 reader，它会等待这些活跃的 reader 完成，才有可能获取到锁。
// 但是，如果之后活跃的 reader 再依赖新的 reader 的话，这些新的 reader 就会等待 writer 释放锁之后才能继续执行。
// 这就形成了一个环形依赖: writer 依赖活跃的 reader -> 活跃的 reader 依赖新来的 reader -> 新来的 reader 依赖 writer。

// 解析：factorial 方法是一个递归计算阶乘的方法，用它来模拟 reader。
// 为了更容易地制造出死锁场景，在这里加上了 sleep 的调用，延缓逻辑的执行。
// 这个方法会调用读锁(第 46 行)，在第 53 行递归地调用此方法，每次调用都会产生一次读锁的调用。
// 所以可以不断地产生读锁的调用，而且必须等到新请求的读锁释放，这个读锁才能释放。
// 同时，我们使用另一个 goroutine 去调用 Lock 方法，来实现 writer。这个 writer 会等待 200 毫秒后才会调用 Lock。
// 这样在调用 Lock 的时候，factorial 方法还在执行中不断调用 RLock。
// 这两个 goroutine 互相持有锁并等待，谁也不会退让一步，
// 满足了“writer 依赖活跃的 reader -> 活跃的 reader 依赖新来的 reader -> 新来的 reader 依赖 writer”的死锁条件，所以就导致了死锁的产生。
// 使用读写锁最需要注意的一点就是尽量避免重入，重入带来的死锁非常隐蔽，而且难以诊断。

func main() {
    var mu sync.RWMutex

    // writer，稍微等待，然后制造一个调用Lock的场景。
    go func() {
        time.Sleep(200 * time.Millisecond)
        mu.Lock()
        fmt.Println("Lock")
        time.Sleep(100 * time.Millisecond)
        mu.Unlock()
        fmt.Println("Unlock")
    }()

    go func() {
        factorial(&mu, 10) // 计算10的阶乘, 10!
    }()
    select {}
}

// 递归调用计算阶乘（模拟 reader）
func factorial(m *sync.RWMutex, n int) int {
    if n < 1 { // 阶乘退出条件
        return 0
    }
    fmt.Println("RLock")
    m.RLock()
    defer func() {
        fmt.Println("RUnlock")
        m.RUnlock()
    }()
    time.Sleep(100 * time.Millisecond)
    // 不断地产生读锁的调用，而且必须等到新请求的读锁释放，这个读锁才能释放。
    return factorial(m, n-1) * n // 递归调用
}
