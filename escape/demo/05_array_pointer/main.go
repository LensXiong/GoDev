package main

//  go build -gcflags '-m -l' ./main.go
// # command-line-arguments
// ./main.go:4:5: moved to heap: a
// ./main.go:5:19: []*int literal does not escape
func main() {
    a := 10
    data := []*int{nil}
    data[0] = &a
}
