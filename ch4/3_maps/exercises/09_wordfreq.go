// 练习 4.9
// 编写一个函数 wordfreq 来汇总输入文本文件中每个单词出现的次数吗在第一次调用 Scan 之前, 需要使用 input.Split(bufio.ScanWords) 来将文本行按照单词分割而不是行分割
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var words = make(map[string]int)
	
	input := bufio.NewScanner(os.Stdin) // 获取输入流
	input.Split(bufio.ScanWords)        // 将输入流按单词分割
	for input.Scan() {
		word := input.Text()
		words[word]++
	}
	
	for k, v := range words {
		fmt.Printf("%s\t出现的次数:\t%d\n", k, v)
	}
}
