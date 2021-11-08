package main

//  go build -gcflags '-m -l' ./main.go
//  #command-line-arguments
// ./main.go:4:26: []interface {} literal does not escape
// ./main.go:4:27: 100 does not escape
// ./main.go:4:32: 200 does not escape
// ./main.go:5:13: 100 escapes to heap
func main() {
    data := []interface{}{100, 200}
    data[0] = 100
}
