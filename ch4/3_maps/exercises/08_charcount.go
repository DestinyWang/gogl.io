// 练习 4.8
// 修改 charcount 的代码来统计字母, 数字和其他在 Unicode 分类中的字符数量, 可以使用函数 unicode.IsLetter 等
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
	
	lCnt := 0       // 字母的数量
	nCnt := 0       // 数字的数量
	oCnt := 0       // 其他字符的数量
	
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
		
		if unicode.IsLetter(r) {
			lCnt++
		} else if unicode.IsNumber(r) {
			nCnt++
		} else {
			oCnt++
		}
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
	
	fmt.Printf("字母数量: %d\n", lCnt)
	fmt.Printf("数字数量: %d\n", nCnt)
	fmt.Printf("其他字符数量: %d\n", oCnt)
}
