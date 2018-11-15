package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	//buf.WriteByte('[')	// 在追加 ASCII 字符的时候, 可以使用 WriteByte, 而追加 UTF-8 编码的时候, 最好使用 WriteRune
	//buf.WriteString("[")	// 也可以直接追加字符串 "["
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	//buf.WriteString("]")
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))	// [1, 2, 3]
}