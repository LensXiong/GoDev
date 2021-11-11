package main

import "fmt"

// 知识点：使用时需要注意的要点。
// 在 Go 中，map[key]函数返回结果可以是一个值，也可以是两个值，这是容易让人迷惑的地方。
// 如果获取一个不存在的 key 对应的值时，会返回零值。
// 为了区分真正的零值和 key 不存在这两种情况，可以根据第二个返回值来区分。

// 结果：a=0; b=0
// a=0, existed: true; b=0, existed: false

func main() {
    var m = make(map[string]int)
    m["a"] = 0
    fmt.Printf("a=%d; b=%d\n", m["a"], m["b"])
    av, aExisted := m["a"]
    bv, bExisted := m["b"]
    fmt.Printf("a=%d, existed: %t; b=%d, existed: %t\n", av, aExisted, bv, bExisted)
}
