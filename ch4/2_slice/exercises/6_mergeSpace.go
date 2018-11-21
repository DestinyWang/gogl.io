package main

import (
	"fmt"
	"unicode"
)

func mergeSpace(slice []byte) []byte {
	cnt := 0                        // 计数器代表当前连续空格的数量
	j := len(slice)
	for i := 0; i < j; {
		b := rune(slice[i])
		if unicode.IsSpace(b) {     // 如果当前是字符是空格, 计数器自增, 如果计数器大于 1, 将当前字符删除, 同时计数器重置
			cnt++
			if cnt > 1 {
				copy(slice[i:], slice[i+1:])
				cnt = 0             // 计数器必须重置而不是--, --代表全局总的空格数而非连续空格数
				j--
			} else {                // 如果不大于 1, 代表这是当前第一个空格, 继续查看下一个
				i ++
			}
		} else {
			i++
		}
	}
	return slice[:j]
}

func main() {
	slice := []byte{'a', 'b', ' ', ' ', 'c', ' '}
	fmt.Printf("%q\n", mergeSpace(slice))   // ['a', 'b', ' ', 'c', ' ']
}
