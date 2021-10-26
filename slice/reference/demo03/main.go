package main

import "fmt"

func test(slice []int) {
    // 改变实参
    slice[0] = 0
}

func main() {
    var slice = []int{1, 2, 3, 4}
    fmt.Println("", slice) // [1 2 3 4]
    test(slice)
    fmt.Println("", slice) // [0 2 3 4]
}
