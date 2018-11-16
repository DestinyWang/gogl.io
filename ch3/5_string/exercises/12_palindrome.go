// 练习 3.12
// 判断两个字符串是否是回文字符串, 即一个字符串中的全部字母都出现在另一个字符串
package exercises

import "bytes"

func PalindromeNormally(s1, s2 string) bool {
	charMap := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		charMap[s1[i]]++
	}
	
	flag := true
	for i := 0; i < len(s2); i++ {
		if charMap[s2[i]] == 0 {
			flag = false
		}
	}
	
	b := []byte(s2)
	for k := range charMap {
		if !bytes.Contains(b, []byte{k}) {
			flag = false
		}
	}
	return flag
}