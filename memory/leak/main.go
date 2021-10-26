package main

import "fmt"

type query func(string) string

// 此代码有严重的内存泄漏问题，出错的位置是 go fn(i) 。
// 代码执行后会启动 4 个协程，但是因为 ch 是非缓冲的，只可能有一个协程写入成功。而其他三个协程会一直在后台等待写入。
func exec(name string, vs ...query) string {
    ch := make(chan string)
    fn := func(i int) {
        ch <- vs[i](name)
    }
    for i, _ := range vs {
        go fn(i)
    }
    return <-ch
}

func main() {
    ret := exec("111", func(n string) string {
        return n + "func1"
    }, func(n string) string {
        return n + "func2"
    }, func(n string) string {
        return n + "func3"
    }, func(n string) string {
        return n + "func4"
    })
    fmt.Println(ret)
}
