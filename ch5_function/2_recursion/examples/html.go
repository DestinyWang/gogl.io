package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

type Attribute struct {
	Key, Val        string
}

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

//func Parse(r io.Reader) (*Node, error)

// 将 n 节点中的每个链接添加到结果中
// 遍历 HTML 树上的所有结果, 从 <a href='...'> 中得到 href 属性的内容, 将获取到的内容添加到 []string, 最后返回这个 slice
func visit(linkes []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {    // 如果当前 Node 类型是 Element, 并且 Node 标签是 a
		for _, a := range n.Attr {                      // 遍历当前 Node 的所有属性
			if a.Key == "href" {
				linkes = append(linkes, a.Val)          // 将 key 为 href 的属性添加到 links 中
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {    // 遍历所有的 child 节点
		linkes = visit(linkes, c)                           // 都进行相同的操作
	}
	return linkes
}

func main() {
	// html.Parse 从给定的阅读器返回 HTML 的解析树
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
