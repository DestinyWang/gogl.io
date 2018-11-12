package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	s1 := "abc"
	s2 := "acc"
	fmt.Println(s1 == s2)
	fmt.Println(s1 < s2)
}
