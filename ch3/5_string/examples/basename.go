// 用于展示文件本身的名字, 出去路径名和后缀
// e.g., a/b/c.go => c
package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	// 保留最后一个 / 之后的全部内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 保留最后一个 . 之前的全部内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename1(s string) string {
	// 取最后一个子串, 如果没取到, 返回 -1
	i := strings.LastIndex(s, "/")
	s = s[i+1:]
	if dot := strings.LastIndex(s, "."); dot != -1 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename1("a/b/c.go"))
}