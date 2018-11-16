// 练习 3.1
// 输入一个三位整数, 求出该数每个位上的数字之和.
package main

func nSum(n int) int {
	if n <= 0 {
		panic("number error")
	}
	cnt := 0
	for n > 0 {
		cnt += n % 10
		n /= 10
	}
	return cnt
}