package main

import (
    "fmt"
    "os"
    "time"
)

// 解析：defer 配合 recover。
// panic 会停掉当前正在执行的程序，不只是当前协程。
// 在这之前，它会有序地执行完当前协程 defer 列表里的语句，其它协程里挂的 defer 语句不作保证。
// 因此，我们经常在 defer 里挂一个 recover 语句，防止程序直接挂掉，这起到了 try...catch 的效果。
// 注意，recover()函数只在 defer 的上下文中才有效（且只有通过在 defer 中用匿名函数调用才有效），直接调用的话，只会返回 nil。

// 结果：
// end of main function
// defer main
// defer here
// defer caller
// recover success. err:  should set user env.

func main() {
    defer fmt.Println("defer main")
    var user = os.Getenv("USER_")
    go func() {
        defer func() {
            fmt.Println("defer caller")
            if err := recover(); err != nil {
                fmt.Println("recover success. err: ", err)
            }
        }()
        func() {
            defer func() {
                fmt.Println("defer here")
            }()
            if user == "" {
                // panic 最终会被 recover 捕获到。这样的处理方式在一个 http server的主流程常常会被用到。¬
                panic("should set user env.")
            }
            // 此处不会执行
            fmt.Println("after panic")
        }()
    }()
    time.Sleep(100)
    fmt.Println("end of main function")
}
