// 练习 5.2
// 写一个函数, 用于统计 HTML 文档树内所有元素的个数
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var cnt = make(map[string]int)


func count(n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		cnt[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(c)
	}
	return cnt
}

func main() {
	if doc, err := html.Parse(os.Stdin); err == nil {
		counts := count(doc)
		for k, v := range counts {
			fmt.Printf("%s\t出现了: %d\t次\n", k, v)
		}
	} else {
		fmt.Fprintf(os.Stderr, "error: %s", err)
	}
	
}
