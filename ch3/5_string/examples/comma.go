package main

import "fmt"

// 接受一个表示整数的字符串, 从右侧开始每三个数字后面就插入一个逗号
func comma(s string) string {
	n := len(s)
	if n <= 3 {			// 若长度小于 3, 直接返回给上层
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]	// 将上层调用加上 "," 以及后三个字母
}

func main() {
	fmt.Println(comma("1234567"))
}