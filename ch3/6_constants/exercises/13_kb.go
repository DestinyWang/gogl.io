// 联系 3.13
// 用尽可能简介的方法声明从 KB, MB 直到 YB 的常量
package  exercises

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	YB
)