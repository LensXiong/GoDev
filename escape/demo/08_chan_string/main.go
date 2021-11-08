package main

// go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:6:18: []string literal escapes to heap
// ./main.go:8:8: func literal escapes to heap
func main() {
    ch := make(chan []string)

    s := []string{"wx"}

    go func() {
        ch <- s
    }()
}
