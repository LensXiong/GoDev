package main

import (
    "fmt"
    "sync"
)

// fatal error: all goroutines are asleep - deadlock!
// 解析：加锁后复制变量，会将锁的状态也复制，所以 mu1 其实是已经加锁状态，再加锁会死锁。

type MyMutex struct {
    count int
    sync.Mutex
}

func main() {
    var mu MyMutex
    mu.Lock()
    var mu2 = mu
    mu.count++
    mu.Unlock()
    mu2.Lock()
    mu2.count++
    mu2.Unlock()
    fmt.Println(mu.count, mu2.count)
}
