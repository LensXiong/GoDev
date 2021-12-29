package main

import (
	"fmt"
	"sync"
)

// writeFile recover Err writeFile panic

// 写文件
func writeFile(wg *sync.WaitGroup) {
	defer wg.Done()
	// defer wg.Add(-1)
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			fmt.Println("writeFile recover Err", recoverErr)
		}
	}()
	panic("writeFile panic")
}

// 打包文件
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go writeFile(&wg)
	}
	wg.Wait()
}
