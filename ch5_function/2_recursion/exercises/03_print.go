// 练习 5.3
// 编写一个函数, 用于输出 HTML 文档树中所有文本节点的内容, 但不包括 <script> <style> 等元素, 因为这些内容在 Web 浏览器中是不可见的
package main

import (
	"fmt"
	"golang.org/x/net/html"
)

func print(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		fmt.Println(n.Attr)
	}
}

func main() {
	
}
