package main

import "fmt"

//  go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:5:10: a does not escape
// ./main.go:10:5: moved to heap: data
// ./main.go:13:16: ... argument does not escape
// ./main.go:13:16: data escapes to heap
func foo(a *int) {
    return
}

func main() {
    data := 10
    f := foo
    f(&data)
    fmt.Println(data)
}
