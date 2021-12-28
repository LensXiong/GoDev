package main

import (
	"fmt"
	"sync"
)

// writeFile recover Err writeFile panic

// 写文件
func writeFile(firstTestScript string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			fmt.Println("writeFile recover Err", recoverErr)
		}
	}()
	panic("writeFile panic")
}

// 打包文件
func main() {
	testScriptSlice := []string{"111", "222"}
	var wg sync.WaitGroup
	for _, v := range testScriptSlice {
		wg.Add(1)
		go writeFile(v, &wg)
	}
	wg.Wait()
}
