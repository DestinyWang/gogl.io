package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]        // slice 仍有增长空间, 扩展 slice 内容
	} else {
		zcap := zlen        // slice 已无空间, 为它分配一个新的底层数组, 为了分摊线性复杂性, 容量扩展一倍
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)          // 内置 copy 函数
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	// 0,	cap=1	[0]
	// 1,	cap=2	[0 1]
	// 2,	cap=4	[0 1 2]
	// 3,	cap=4	[0 1 2 3]
	// 4,	cap=8	[0 1 2 3 4]
	// 5,	cap=8	[0 1 2 3 4 5]
	// 6,	cap=8	[0 1 2 3 4 5 6]
	// 7,	cap=8	[0 1 2 3 4 5 6 7]
	// 8,	cap=16	[0 1 2 3 4 5 6 7 8]
	// 9,	cap=16	[0 1 2 3 4 5 6 7 8 9]
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d,\tcap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
