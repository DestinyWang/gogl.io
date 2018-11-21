// 练习4.4 
// 编写一个函数 rotate, 实现一次遍历就可以完成元素旋转
package main

import "fmt"

// 样例: 12345, 3 就是将 12345 的向左旋转 3 位 => 45123
func rotate(s []int, l int) []int {
	s = reverse(s, 0, l - 1)
	s = reverse(s, l, len(s) - 1)
	return reverse(s, 0, len(s) - 1)
}

// 12345, 3
// 123, 45
// 321, 54
// 45124
func reverse(s []int, low, high int) []int {
	for i, j := low, high; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", rotate(s, 3))
}
