# 概述
本项目用于针对`《Go程序设计语言》`一书的以下几项内容:
1. 书中概念的讲解(以 `.md` 的形式)
2. 示例代码以及本人简单的思路讲解(以代码注释的形式)
3. 课后习题(同样给出思路)
4. 部分重要概念的单独描述(以穿插在 `章节(ch*)` 目录下的 `.md` 形式文件为主)

注释的使用方式:

| 注释方式 | 使用场景 |
| :-: | :-: |
| 行首注释 | 对本行代码内容进行说明 |
| 行尾注释 | 对程序运行流程进行说明 |

示例:

```go
// 一次读取整个输入到大块内存
func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        // ReadFile 函数返回一个可以转化成字符串的字节 slice, 这样它可以被 strings.Split 分割
        data, err := ioutil.ReadFile(filename)							// 读取文件内容
        if err != nil {
            fmt.Fprintf(os.Stdout, "dup2: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {	// strings.Split 将文件中读到的内容以"换行符"进行分割
            // 对获取的次数直接+1, 如果为空则 0+1
            counts[line]++						// 再将每行进行重复统计
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```

# 勘误表
