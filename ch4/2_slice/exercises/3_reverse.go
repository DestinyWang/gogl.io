// 练习 4.3
// 重写函数 reverse, 使用数组指针作为参数而不是 slice
package main

import "fmt"

func Reverse(ptr *[6]int) {
	for i, j := 0, 5; i < j; i, j = i + 1, j - 1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func main() {
	a1 := [6]int{1, 2, 3, 4, 5, 6}
	a2 := [6]int{10, 20, 30, 40, 50, 60}
	Reverse(&a1)
	Reverse(&a2)
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
}
