package main

import (
    "fmt"
    "sync"
    "time"
)

var mu sync.RWMutex
var count int

// fatal error: all goroutines are asleep - deadlock!
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
