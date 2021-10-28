package main

import (
    "reflect"
    "unsafe"
)

// 结果：
// unexpected fault address 0x109d9ff
// fatal error: fault
// [signal SIGBUS: bus error code=0x2 addr=0x109d9ff pc=0x107ee5c]

// 解析：string 类型是不能改变的，也就是底层数据是不能更改的。
// 这里使用的是强转换的方式，那么 by 指向了 str 的底层数组，现在对这个数组中的元素进行更改，程序直接发生严重错误了，即使使用 defer+recover 也无法捕获。


func stringToSliceByteTmp(s string) []byte {
    str := (*reflect.StringHeader)(unsafe.Pointer(&s))
    ret := reflect.SliceHeader{Data: str.Data, Len: str.Len, Cap: str.Len}
    return *(*[]byte)(unsafe.Pointer(&ret))
}

func main() {
    str := "hello"
    by := stringToSliceByteTmp(str)
    by[0] = 'H'
}
