package main

import "fmt"
import "time"
import "sync"

type group struct {
    gc chan bool
    tk *time.Ticker
    cap int
    mutex sync.Mutex
}

// 运行结果：
// start...
// exec goroutine product
// exec goroutine consumer
// all goroutine exec over
// exec return

func WaitGroup(timeOutRec int) *group{
    timeout     := time.Millisecond * time.Duration(timeOutRec)
    wg := group{
        gc   : make(chan bool),
        cap  :  0,
        tk   : time.NewTicker(timeout),
    }

    return &wg
}


func(w *group)Add(index int){
    w.mutex.Lock()
    w.cap++
    w.mutex.Unlock()

    go func(w *group,index int) {
        for i := 0;i<index;i++{
            fmt.Println("exec goroutine product")
            w.gc<- true
        }
    }(w,index)
}

func(w *group)Done(){
    <-w.gc
    fmt.Println("exec goroutine consumer")
    w.mutex.Lock()
    w.cap--
    w.mutex.Unlock()
}

func(w *group)Wait(){
    defer w.tk.Stop()
    for  {
        select {
        case <-w.tk.C:
            fmt.Println("time out exec over")
            return;
        default:
            w.mutex.Lock()
            if w.cap == 0 {
                fmt.Println("all goroutine exec over")
                return;
            }
            w.mutex.Unlock()
        }
    }
}

func main() {
    fmt.Println("start...")
    wg := WaitGroup(10)
    wg.Add(1)
    closer := func(wg *group) {
        wg.Done()
    }
    go closer(wg)

    wg.Wait()
    fmt.Println("exec return")
}