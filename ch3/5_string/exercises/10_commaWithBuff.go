// 练习 3.10
// 编写一个非递归的 comma 函数, 运用 bytes.Buffer, 而不是简单的字符串拼接
package exercises

import (
	"bytes"
)

func Comma(s string) string {
	var buff bytes.Buffer
	l := len(s)
	for i := 0; i < l; i++ {
		buff.WriteByte(s[i])
		if (l-i-1)%3 == 0 && i != l-1 {
			buff.WriteByte(',')
		}
	}
	return buff.String()
}

