package main

import "fmt"

func main() {
	ints := make([]int, 0)
	for i := 0; i < 10; i++ {
		ints = append(ints, i)
		fmt.Printf("%d,\tcap=%d\t%v\n", i, cap(ints), ints)
	}
}
