package main

import (
    "fmt"
    "time"
)

// 知识点1：map 和 slice 或者 Mutex、RWMutex 等 struct 类型不同，map 对象必须在使用之前初始化。
// 如果不初始化就直接赋值的话，会出现 panic 异常。

// 知识点2：从一个 nil 的 map 对象中获取值不会 panic，而是会得到零值。

// 知识点3：map 作为一个 struct 字段的时候，别忘记初始化。

// 结果：panic: assignment to entry in nil map
type Counter struct {
    Website      string
    Start        time.Time
    PageCounters map[string]int
}

func main() {
    var c Counter
    c.PageCounters = make(map[string]int)
    c.Website = "wwxiong.com"

    c.PageCounters["/"]++

    // map 使用前必须初始化
    var m1 map[int]int
    m1[100] = 100

    // 从一个 nil 的 map 对象中获取值不会 panic，而是会得到零值，所以下面的代码不会报错。
    var m map[int]int
    fmt.Println(m[100]) // 0
}
