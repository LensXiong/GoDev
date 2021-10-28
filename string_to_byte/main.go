package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

// 问题：字符串转成 byte 数组，会发生内存拷⻉吗? 有没有什么办法可以在字符串转成切片的时候不用发生拷⻉呢?

// 解析1：如果想要在底层转换二者，只需要把 StringHeader 的地址强转成 SliceHeader 就行。 go有个很强的包叫 unsafe 。
// 1. unsafe.Pointer(&a) 方法可以得到变量a的地址。
// 2. (*reflect.StringHeader)(unsafe.Pointer(&a)) 可以把字符串a转成底层结构的形式。
// 3. (*[]byte)(unsafe.Pointer(&ssh)) 可以把 ssh 底层结构体转成 byte 的切片的指针。
// 4. 再通过 * 转为指针指向的实际内容。

// 强制转换带来的安全问题 // b[0] = 10
// unexpected fault address 0x10ce277
// fatal error: fault
// [signal SIGBUS: bus error code=0x2 addr=0x10ce277 pc=0x10a8b21]
// 解析2：string 类型是不能改变的，也就是底层数据是不能更改的。
// 这里使用的是强转换的方式，那么 by 指向了 str 的底层数组，现在对这个数组中的元素进行更改，程序直接发生严重错误了，即使使用 defer+recover 也无法捕获。

// StringHeader 是字符串在 go 的底层结构。
// StringHeader is the runtime representation of a string.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
type StringHeader struct {
    Data uintptr
    Len  int
}

// SliceHeader 是切片在 go 的底层结构。
// SliceHeader is the runtime representation of a slice.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}

func main() {
    a := "wang xiong"
    ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
    b := *(*[]byte)(unsafe.Pointer(&ssh))
    // unexpected fault address 0x10ce277
    // fatal error: fault
    // [signal SIGBUS: bus error code=0x2 addr=0x10ce277 pc=0x10a8b21]
    // b[0] = 10 // 强制转换带来的安全问题
    fmt.Printf("%v", b) // [119 97 110 103 32 120 105 111 110 103]
}
