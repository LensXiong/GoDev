package main

import (
    "fmt"
    "sync"
)

var a int
var wg sync.WaitGroup

// goroutine1
func main() {

    wg.Add(1)
    // 1 goroutine2
    go func() {
        a = 1 // 1.1
        defer wg.Done()
    }()

    wg.Wait()
    // 2 goroutine1
    if 0 == a { // 2.1
        fmt.Println(a) //2.2
    } else {
        fmt.Println(a)
    }
}
