package main

import (
	"bufio"
	"fmt"
	"os"
)

// 可以从标准输入或者文件列表进行读取
// 以流式模式读取输入, 然后按需拆分为行
func main() {
	counts := make(map[string]int)		// 先初始化记录内容和对应次数的 map
	files := os.Args[1:]				// 从命令行的第 1 个参数起获取文件路径
	if len(files) == 0 {				// 如果命令行参数不存在
		countLines(os.Stdin, counts)	// 直接统计标准输入的内容
	} else {
		for _, arg := range files {		// 否则循环读取文件的内容(索引不需要使用因此使用 _ 代替)
			f, err := os.Open(arg)		// 尝试打开文件
			if err != nil {				// 如果打开失败, 输出错误信息
				fmt.Fprintf(os.Stdout, "dup2: %v\n", err)
				continue				// 继续处理下一个文件
			}
			countLines(f, counts)		// 统计当前文件的内容
			f.Close()					// 关闭文件
		}
	}
	for line, n := range counts {		// 遍历 map
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)	// 如果出现次数大于 1, 输出内容
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)		// 获取文件输入或标准输入
	for input.Scan() {					// 调用 input.Scan() 读取下一行
		counts[input.Text()]++			// 调用 input.Text() 获取内容, 并更新内容对应的出现次数
	}
}
