package main

import (
    "fmt"
    "io"
    "os"
)
// 场景：在一个函数里，需要打开两个文件进行合并操作，合并完后，在函数执行完后关闭打开的文件句柄。

// defer 函数定义的时候，参数就已经复制进去了，之后真正执行 close() 函数的时候就刚好关闭的是正确的文件了。
// 如果不这样将 f 当成函数参数传递进去的话，最后两个语句关闭的就是同一个文件了，都是最后一个打开的文件。

func mergeFile() error {
    f, _ := os.Open("file1.txt")
    if f != nil {
        defer func(f io.Closer) {
            if err := f.Close(); err != nil {
                fmt.Printf("defer close file1.txt err %v\n", err)
            }
        }(f)
    }
    // ……
    f, _ = os.Open("file2.txt")
    if f != nil {
        defer func(f io.Closer) {
            if err := f.Close(); err != nil {
                fmt.Printf("defer close file2.txt err %v\n", err)
            }
        }(f)
    }
    return nil
}
