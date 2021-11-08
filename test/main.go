package main

import (
    "fmt"
    "runtime"
    "time"
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

type Person struct {
    Friends []string
}

func main() {
    var ch chan int
    go func() {
        ch<- 1
        ch<- 2
    }()
    //go func() {
    //    <-ch
    //}()
    c := time.Tick(1 * time.Second)
    for range c {
        // NumGoroutine returns the number of goroutines that currently exist.
        fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
    }
    //var f1 []string // nil切片
    //json1, _ := json.Marshal(Person{Friends: f1})
    //fmt.Printf("%s\n", json1) // output：{"Friends": null}
    //
    //f2 := make([]string, 0) // non-nil空切片
    //json2, _ := json.Marshal(Person{Friends: f2})
    //fmt.Printf("%s\n", json2) // output: {"Friends": []}
}
