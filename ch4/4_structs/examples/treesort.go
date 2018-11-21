package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// 就地排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
		appendValues(values[:0], root)
	}
}

// appendValues 将元素按照顺序追加到 values 里面, 然后返回结果 false
func appendValues(values []int, t *tree) []int {
	// 中序遍历将元素添加到 values 中
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// 等价于返回 &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	t := tree{
		value:1,
		left:nil,
		right:nil,
	}
	fmt.Println(t)
}
