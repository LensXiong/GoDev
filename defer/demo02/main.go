package main

import "fmt"

func f1() (r int) {
    defer func(r *int) {
        *r = *r + 5
    }(&r)
    return 1
}
func main() {
    fmt.Println(f1()) // 6
}
