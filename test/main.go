package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    defer mu.Unlock()
    fmt.Println("hello world!")
}
