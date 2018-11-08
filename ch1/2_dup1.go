package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)		// 创建一个 key 为 int, value 为 string 的 map
	input := bufio.NewScanner(os.Stdin)	// 获取标准输入
	for input.Scan() {					// 调用 input.Scan() 读取下一行
		counts[input.Text()]++			// 调用 input.Text() 获取内容, 并更新内容对应的出现次数
	}

	fmt.Println()
	for k, v := range counts {			// 遍历 map 并获取每个元素的内容和次数
		if v > 1 {
			fmt.Printf("%d, %s\n", v, k)	// 如果出现超过一次, 就打印出来
		}
	}
}
