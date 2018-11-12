package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 一次读取整个输入到大块内存
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile 函数返回一个可以转化成字符串的字节 slice, 这样它可以被 strings.Split 分割
		data, err := ioutil.ReadFile(filename)							// 读取文件内容
		if err != nil {
			fmt.Fprintf(os.Stdout, "dup2: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {	// strings.Split 将文件中读到的内容以"换行符"进行分割
			// 对获取的次数直接+1, 如果为空则 0+1
			counts[line]++						// 再将每行进行重复统计
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
