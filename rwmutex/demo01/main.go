package main

import (
    "fmt"
    "sync"
    "time"
)

var mu sync.RWMutex
var count int

// fatal error: all goroutines are asleep - deadlock!
// 解析：会产生死锁 panic，读写锁当有一个协程在等待写锁时，其他协程是不能获得读锁的。
// 而在 A 和 C 中同一个调用链中间需要让出读锁，让写锁优先获取，而 A 的读锁又要求 C 调用完成，因此死锁。
func main() {
    go A()
    time.Sleep(1 * time.Second)
    mu.Lock()
    defer mu.Unlock()
    count++
    fmt.Println(count)
}

func A() {
    mu.RLock()
    defer mu.RUnlock()
    B()
}

func B() {
    time.Sleep(2 * time.Second)
    C()
}

func C() {
    mu.RLock()
    defer mu.RUnlock()
}
