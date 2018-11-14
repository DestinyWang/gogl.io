package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))						// 13
	fmt.Println(utf8.RuneCountInString(s))	// 9

	/*
	0	H
	1	e
	2	l
	3	l
	4	o
	5	,
	6
	7	世
	10	界
	 */
	for i := 0; i < len(s); {
		// DecodeRuneInString 返回一个 r(文字符号本身) 和 一个值(表示 r 按 UTF-8 编码所占用的字节数)
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size									// 用当前文字符号的 UTF-8 编码所占用的字节数更新下标
	}
}
