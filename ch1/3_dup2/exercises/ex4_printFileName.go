package main

import (
	"bufio"
	"fmt"
	"os"
)

// 修改 dup2 程序, 输出出现重复行的文件的名称
func main() {
	counts := make(map[string]map[string]int)	// 创建一个类似 Map<String, Map<String, Integer>> 的结构, 用于存储文件名-每行内容-出现次数的映射关系
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			open, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stdout, "dup2: %v\n", err)
				continue
			}
			countLines(open, counts)
			open.Close()
		}
	}
	for name, linNoMap := range counts {			// 遍历每个文件
		for _, n := range linNoMap {				// 遍历每个文件中每行内容及其出现次数的映射关系
			if n > 1 {
				fmt.Printf("%s\n", name)		// 对于出现重复内容的文件, 输出其文件名
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	name := f.Name()
	for input.Scan() {
		counts[name][input.Text()]++
	}
}