package main

import (
    "fmt"
    "unsafe"
)

// 运行结果：
// int:8
// int8:1
// int16:2
// int32:4
// int64:8
// slice:24

// 知识点：为什么slice会占24byte？
/*
type slice struct {
	array unsafe.Pointer // 指向底层数组的指针
	len   int // 切片的长度
	cap   int // 切片的长度
}

type Pointer *ArbitraryType
type ArbitraryType int
*/

func main() {
    var a int
    var b int8
    var c int16
    var d int32
    var e int64
    slice := make([]int, 0)
    slice = append(slice, 1)
    fmt.Printf("int:%d\nint8:%d\nint16:%d\nint32:%d\nint64:%d\n",
        unsafe.Sizeof(a),
        unsafe.Sizeof(b),
        unsafe.Sizeof(c),
        unsafe.Sizeof(d),
        unsafe.Sizeof(e))
    fmt.Printf("slice:%d", unsafe.Sizeof(slice))
}
