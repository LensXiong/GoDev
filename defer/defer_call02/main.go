package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        defer fmt.Println(i, 1) // 4 3 2 1 0
    }

    // 闭包引用类型
    for i := 0; i < 5; i++ {
        defer func() {
            fmt.Println(i, 2) // 5 5 5 5 5
        }()
    }

    for i := 0; i < 5; i++ {
        defer func() {
            j := i
            // fmt.Println(j, 3) // 5 5 5 5 5
            // 5 0xc00001c0b0 5 0xc00001c0a8 3
            // 5 0xc00001c0d0 5 0xc00001c0a8 3
            // 5 0xc00001c0d8 5 0xc00001c0a8 3
            // 5 0xc00001c0e0 5 0xc00001c0a8 3
            // 5 0xc00001c0e8 5 0xc00001c0a8 3
            fmt.Println(j, &j, i, &i, 3)
        }()
    }

    for i := 0; i < 5; i++ {
        j := i
        defer fmt.Println(j, 4) // 4 3 2 1 0
    }

    for i := 0; i < 5; i++ {
        j := i
        defer func() {
            fmt.Println(j, 5) // 4 3 2 1 0
        }()
    }

    for i := 0; i < 5; i++ {
        defer func(j int) {
            fmt.Println(j, 6) // 4 3 2 1 0
        }(i)
    }

}
