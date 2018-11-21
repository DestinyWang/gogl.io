// 练习 4.5
// 编写一个就地处理函数, 用于去除 []string slice 中相邻的重复字符串
package main

import "fmt"

func distinct(slice []string) []string {
	flag := slice[0]
	j := len(slice)
	for i := 1; i < j; {
		if slice[i] == flag {               // 如果当前元素等于 flag
			copy(slice[i:], slice[i+1:])    // 将 i+1 后的所有元素 copy 到 i 后, 即最后一个元素会重复
			j--                             // copy 后, 删除尾部的多余元素
		} else {
			flag = slice[i]                 // 如果当前元素不等于 flag
			i++                             // 继续比较下一个元素
		}
	}
	return slice[:j]
}

func main() {
	slice := []string{"a", "b", "c", "d", "c", "d", "a", "a", "a", "destiny"}
	fmt.Printf("%v\n", distinct(slice)) // [a b c d c d a destiny]
}
