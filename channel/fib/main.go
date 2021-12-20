package main

import "fmt"

func fibonaci(c, quit chan int) {
    x, y := 1, 1

    for {
        select {
        case c <- x:
            t := x
            x = y
            y = t + y
        case data := <-quit:
            fmt.Println("quit", data)
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)

    // sub go
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }

        quit <- 0
    }()

    // main go
    fibonaci(c, quit)
}

