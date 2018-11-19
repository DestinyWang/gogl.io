package main

import "fmt"

// 输入一个 []string
// 对其中每个 string 进行判断, 如果是空串就跳过
// 最终返回一个所有元素都不是空串的 []string
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]              // 引用原始 slice 的新的零长度的 slice
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", data)            // ["one" "" "three"]
	fmt.Printf("%q\n", nonempty(data))  // ["one" "three"]
	fmt.Printf("%q\n", data)            // ["one" "three" "three"]
}
