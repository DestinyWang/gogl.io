// Go 在 net 包下提供了一系列包用于通过互联网发送和接受信息, 使用底层的网络连接, 创建服务器
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {			// 遍历所有命令行参数
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
