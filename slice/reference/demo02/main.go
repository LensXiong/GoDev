package main

import "fmt"

// golang 切片是引用类型，所以在传递时，遵守引用传递机制。
func main() {
    var slice []int
    var arr [5]int = [...]int{1, 2, 3, 4, 5}
    slice = arr[:]
    var slice2 = slice
    slice2[0] = 0
    fmt.Println("slice2", slice2) // [0 2 3 4 5]
    fmt.Println("slice", slice)   // [0 2 3 4 5]
    fmt.Println("arr", arr)       // [0 2 3 4 5]
}
