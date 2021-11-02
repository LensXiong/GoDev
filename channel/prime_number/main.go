package main

import (
    "fmt"
)

func Processor(seq <-chan int, wait chan struct{}, level int) {
    go func() {
        prime, ok := <-seq
        if !ok {
            close(wait)
            return
        }
        fmt.Printf("[%d]: %d\n", level, prime)
        out := make(chan int)
        Processor(out, wait, level+1)
        for num := range seq {
            if num%prime != 0 {
                out <- num
            }
        }
        close(out)
    }()
}

func main() {
    origin, wait := make(chan int), make(chan struct{})
    Processor(origin, wait, 1)
    for num := 2; num < 10; num++ {
        origin <- num
    }
    close(origin)
    <-wait
}
