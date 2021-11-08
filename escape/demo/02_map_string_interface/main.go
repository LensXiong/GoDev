package main

// go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:4:17: make(map[string]interface {}) does not escape
// ./main.go:5:17: 200 escapes to heap
func main() {
    data := make(map[string]interface{})
    data["key"] = 200
}
