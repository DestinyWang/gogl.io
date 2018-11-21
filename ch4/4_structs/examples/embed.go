package main

import "fmt"

type Point struct {
	X, Y    int
}

type Circle struct {
	Point
	Radius  int
}

type Wheel struct {
	Circle
	Spokes  int
}

func main() {
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	w2 := Wheel{
		Circle: Circle{
			Point: Point{
				X:8,
				Y:8,
			},
			Radius:5,
		},
		Spokes:20,
	}
	
	w2.X = 9
	
	// # 似的 Printf 的格式化符号 %v 以类似 Go 语法的方式输出对象, 包含了成员变量的名称
	fmt.Printf("%#v\n", w1) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
	fmt.Printf("%#v\n", w2) // main.Wheel{Circle:main.Circle{Point:main.Point{X:9, Y:8}, Radius:5}, Spokes:20}
}
