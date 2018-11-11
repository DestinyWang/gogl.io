// 和 fetch 一样获取 URL 的内容, 但它并发获取很多 URL 内容
// 于是这个进程使用的时间不超过耗时最长时间的获取任务, 而不是所有获取任务总的时间
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// main 函数在一个 goroutine 中执行, 然后 go 语句创建额外的 goroutine
// 当一个 goroutine 试图在一个通道上进行发送或接受操作时, 它会阻塞,
// 直到另一个 goroutine 试图进行接收或者发送操作才传递值, 并开始处理两个 goroutine
func main() {
	start := time.Now()
	ch := make(chan string)			// 使用 make 创建一个字符串通道
	for _, url := range os.Args[1:] {
		go fetch(url, ch)			// 启动一个 goroutine, 异步调用 fetch
	}
	for range os.Args[1:] {
		fmt.Println(<- ch)			// 从 ch 通道接收并输出汇总数据
	}
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

// 每一个 fetch 在通道 ch 上发送一个值 (ch <- fmt.Sprint()), main 函数接收他们 (<- ch)
// 由 main 函数来处理所有的输出确保了每个 goroutine 作为一个整体单元处理, 这样就避免了两个 goroutine 同时完成造成输出交织所带来的风险
func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)		// 发送到通道 ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)	// 获取 HTTP 响应正文的长度
	resp.Body.Close()				// 不要泄露资源
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, nbytes, url)	// 输出当前通道的运行时间, 汇总到通道 ch 中
}