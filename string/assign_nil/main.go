package main

import (
    "fmt"
)

// 字符串一旦赋值了，字符串就不能修改了:在 Go 中字符串是不可变的。
func main() {
    // cannot use nil as type string in assignment
    var x string = nil
    // invalid operation: x == nil (mismatched types string and nil)
    if x == nil {
        x = "default"
    }
    fmt.Println(x)
}
