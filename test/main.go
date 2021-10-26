package main

import (
    "fmt"
)

type Student struct {
    Age int
}

func main() {
    kv := map[string]Student{"wangxiong": {Age: 21}}
    // cannot assign to struct field kv["wangxiong"].Age in map
    // kv["wangxiong"].Age = 22
    s := []Student{{Age: 21}}
    s[0].Age = 22
    fmt.Println(kv, s)
}
