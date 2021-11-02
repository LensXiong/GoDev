package main

import "fmt"

// 结果：main func end
// recover: defer panic

func main() {
    go func() {
        defer func() {
            if e := recover(); e != nil {
                fmt.Printf("recover: %v", e)
            }
        }()
        panic("defer panic")
    }()

    fmt.Println("main func end")
}
