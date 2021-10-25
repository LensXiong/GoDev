package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    // GOMAXPROCS sets the maximum number of CPUs that can be executing
    runtime.GOMAXPROCS(1)
    // A WaitGroup waits for a collection of goroutines to finish.
    wg := sync.WaitGroup{}
    // Add adds delta, which may be negative, to the WaitGroup counter.
    wg.Add(20)
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println("i: ", i)
            wg.Done()
        }()
    }
    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Println("i: ", i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}


