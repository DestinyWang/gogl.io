package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

// 常量
const (
	whiteIndex = 0 // 画板中的第一种颜色
	blackIndex = 1 // 画板中的第二种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {		// 如果命令行中带有 "web", 则以 web 服务器的形式启动
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的 x 振荡器变化的个数
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布包含 [-size .. +size]
		nframes = 64    // 动画中的帧数
		delay   = 8     // 以 10ms 为单位的帧间延迟
	)

	freq := rand.Float64() * 3.0 			// y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}		// GIF 是一个结构体
	phase := 0.0 							// phase difference
	for i := 0; i < nframes; i++ {			// 外层循环 64 个迭代, 每个迭代产生一个动画帧
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)	// 创建一个 201 * 201 的画板
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {	// 内层循环通过设置一些像素为黑色产生一个新的图像
			x := math.Sin(t)							// x 轴是正弦函数
			y := math.Sin(t*freq + phase)				// y 轴也是正弦化的,
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)	// 每个内层循环通过设置一些色素为黑色产生一个新的图像
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)			// 然后追加到 anim 的帧列表中
	}
	gif.EncodeAll(out, &anim)						// 最后写入输出流 out
}
