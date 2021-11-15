package main

import "fmt"

// go tool compile -S main.go | grep CALL

// 0x0042 00066 (main.go:7)        CALL    runtime.makeslice(SB)
// 0x006d 00109 (main.go:8)        CALL    runtime.growslice(SB)
// 0x00a4 00164 (main.go:9)        CALL    runtime.convTslice(SB)
// 0x00c0 00192 (main.go:9)        CALL    runtime.convT64(SB)
// 0x00d8 00216 (main.go:9)        CALL    runtime.convT64(SB)

func main() {
    slice := make([]int, 0)
    slice = append(slice, 1)
    fmt.Println(slice, len(slice), cap(slice))
}
