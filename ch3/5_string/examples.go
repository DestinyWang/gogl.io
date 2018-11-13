package main

import "fmt"

const GoUsage = `Go is a tool for managing Go source Code.
Usage:\n
	gp command [arguments]\t\n\\
...`

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	s1 := "abc"
	s2 := "acc"
	fmt.Println(s1 == s2)
	fmt.Println(s1 < s2)

	fmt.Println(GoUsage)
}
