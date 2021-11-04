package main

import (
    "fmt"
)

type People interface {
    Show()
}

type Student struct{}

func (stu *Student) Show() {}

func live() People {
    var stu *Student // <nil>
    fmt.Println(stu)
    return stu
}

func main() {
    if live() == nil {
        fmt.Println("AAAAAAAAAAA")
    } else {
        fmt.Println("BBBBBBBBBBB")
    }
}
