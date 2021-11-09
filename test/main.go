package main

import (
    "fmt"
    "time"
)

func main() {
    //1
    quit := make(chan struct{})
    //2
    go func() {
        for {
            //2.1
            select {
            case v := <-quit: //2.1.1
                fmt.Println(v)
                fmt.Println("sub goroutine is over")
                return
            default: //2.1.2
                //dosomthing
                time.Sleep(time.Second)
                fmt.Println("sub goroutine do somthing")
            }
        }
    }()
    //3.dosomthing
    time.Sleep(time.Second * 3)
    //4.关闭通道quit
    fmt.Println("main gorutine start stop sub goroutine")
    close(quit)
    //5
    time.Sleep(time.Second * 10)
    fmt.Println("main gorutine is over")
}
