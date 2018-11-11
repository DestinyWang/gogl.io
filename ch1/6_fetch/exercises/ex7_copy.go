// 函数 io.Copy(dst, src) 从 src 读, 并且写入 dst, 使用它代替 ioutil.ReadAll 来复制响应内容到 os.Stdout
// 这样不需要装下整个相应数据流的缓冲区, 并确保检查 io.Copy 返回的错误结果
package main

import (
	"fmt"
	"io"
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
		// b, err := ioutil.ReadAll(resp.Body)	// 读取整个响应并存入 b
		// Copy() 函数直接将 HTTP 响应的正文部分复制到标准输出中, 因此这一步会直接输出
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()						// 关闭 body 数据流来避免资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d\n", b)
	}
}