package main

import (
	"fmt"
	"os"
)

// 修改 echo 程序, 输出参数的索引和值
func main() {
	for i, s := range os.Args[1:] {
		fmt.Printf("index: %d, arg: %s\n", i, s)
	}
}
