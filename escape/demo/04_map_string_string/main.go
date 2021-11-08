package main

// go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:4:17: make(map[string][]string) does not escape
// ./main.go:5:27: []string literal escapes to heap
func main() {
    data := make(map[string][]string)
    data["key"] = []string{"value"}
}
