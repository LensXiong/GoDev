package main

import "fmt"

// go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:5:10: a does not escape
// ./main.go:10:18: []string literal escapes to heap
// ./main.go:12:16: ... argument does not escape
// ./main.go:12:16: s escapes to heap
func foo(a []string) {
    return
}

func main() {
    s := []string{"wx"}
    foo(s)
    fmt.Println(s)
}
