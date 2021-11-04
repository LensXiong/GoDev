package main

import "fmt"

func Foo(x interface{}) {
    fmt.Println(x) // <nil>
    if x == nil {
        fmt.Println("empty interface")
        return
    }
    fmt.Println("non-empty interface") // non-empty interface
}
func main() {
    var p *int = nil
    Foo(p)
}
