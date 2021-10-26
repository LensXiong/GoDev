package main

import (
    "fmt"
    "sync"
)

var mu sync.Mutex
var chain string

// fatal error: all goroutines are asleep - deadlock!
// 解析：会产生死锁 panic ，因为 Mutex 是互斥锁。

func main() {
    chain = "main"
    A()
    fmt.Println(chain)
}

func A() {
    mu.Lock()
    defer mu.Unlock()
    chain = chain + " --> A"
    B()
}

func B() {
    chain = chain + " --> B"
    C()
}

func C() {
    mu.Lock()
    defer mu.Unlock()
    chain = chain + " --> C"
}
