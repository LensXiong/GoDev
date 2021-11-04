package main

import "fmt"

type Student struct {
    Name string
}

func main() {

    list := make(map[string]*Student)

    student := Student{"wangxiong"}

    list["student"] = &student
    list["student"].Name = "wwxiong"

    fmt.Println(list["student"]) // &{wwxiong}
}
