package main

import (
    "fmt"
)

var a int

// goroutine1
func main() {

    // 1 goroutine2
    go func() {
        a = 1 // 1.1
    }()

    // 2 goroutine1
    if 0 == a { // 2.1
        fmt.Println(a) //2.2
    }
}
