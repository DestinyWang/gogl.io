package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int) // Unicode 字符数量
	var utfLen [utf8.UTFMax]int  // UTF-8 编码的长度
	invalid := 0
	
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()  // 解码 UTF-8 编码, 并返回三个值: 解码的字符, UTF-8 编码中字节的长度, 错误值
		if err == io.EOF {          // 唯一可能出现的错误是文件结束(EOF)
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 { // 如果出现不合法的 UTF-8 字符, 那么返回的字符就是 unicode.ReplacementChar, 并且长度是 1
			invalid++                               // 非法字符计数器自增 1
			continue
		}
		counts[r]++
		utfLen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utfLen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
// 大家好, 我是destiny!
//
// rune    count
// '好'    1
// ' '     1
// '是'    1
// 'y'     1
// '\n'    1
// '大'    1
// '家'    1
// 'd'     1
// 't'     1
// ','     1
// '我'    1
// 'i'     1
// '!'     1
// 'e'     1
// 's'     1
// 'n'     1
//
// len     count
// 1       11
// 2       0
// 3       5
