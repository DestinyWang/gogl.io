// 修改 fetch 程序, 添加一个 http:// 前缀(加入该 URL 参数确实协议前缀), 可能会用到 strings.HasPrefix
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {			// 遍历所有命令行参数
		prefix := "http://"
		if !strings.HasPrefix(url, prefix) {	// 如果传入的 URL 不包含指定前缀
			url += prefix + url					// 补上该前缀
		}
		resp, err := http.Get(url)				// 产生一个 HTTP 请求去访问指定 url, 将访问结果保存在 resp 中
		// 此时 resp 的 body 域包含一个可读取的数据流
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)		// 读取整个响应并存入 b
		resp.Body.Close()						// 关闭 body 数据流来避免资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}
