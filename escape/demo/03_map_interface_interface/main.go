package main

//  go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:4:17: make(map[interface {}]interface {}) does not escape
// ./main.go:5:9: 100 escapes to heap
// ./main.go:5:15: 200 escapes to heap
func main() {
    data := make(map[interface{}]interface{})
    data[100] = 200
}
