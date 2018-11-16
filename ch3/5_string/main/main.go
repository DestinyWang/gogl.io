package main

import (
	"fmt"
	"study/ch3/5_string/exercises"
)

func main() {
	fmt.Println(exercises.Comma("1234567"))
	fmt.Println(exercises.FloatComma(3.14159269453453626345))
	
	fmt.Println(exercises.PalindromeNormally("destiny", "destinyyyyy"))   // true
	fmt.Println(exercises.PalindromeNormally("destiny", "destnyyyyy"))    // false
	fmt.Println(exercises.PalindromeNormally("destin", "destiny"))        // false
}
