> 在 Go 中, map 是散列表的应用, 它是一个拥有键值对元素的无序集合, map 的类型是 `map[k]v`, 其中 k 和 v 是字典的键和值对应的数据类型, map 中所有的键都拥有相同的数据类型, 同时所有的值也都拥有相同的数据类型, 但是键的类型和值的类型不一定相同, 键的类型 k, 必须是可以通过 `==` 来进行比较的数据类型, 所以 map 可以检测某一个键是否已经存在, 虽然浮点型是可以比较的, 但由于存在误差, 大多数场景中并不是一个好主意.


```go
// 内置函数 make 可以用来创建一个 map
ages := make(map[string]int)    // 创建一个从 string 映射到 int 的 map

// 也可以使用 map 的字面量来新建一个带初始化键值对的 map
ages := map[string]int {
    "destiny":  24,
    "camery":   24,
}

// map 的元素访问也是通过下标的方式
fmt.Println(ages["destiny"])    // 24

// 可以使用内置函数 delete 来从字典充根据键删除一个元素
delete(ages, "destiny")         // 删除元素 ages["destiny"]

// 可以使用 range 来遍历 map 中的所有元素
for k, v := range ages {
    fmt.Println("%s\t%d\n", k, v)
}
```

但是 map 不是一个变量, 不可以获取它的地址

```go
_ = &map["destiny"]     // 编译错误, 无法获取 map 元素的地址
```

map 中元素的迭代顺序是不固定的, 不同的实现方式会使用不同的散列算法, 得到不同的元素顺序, 实践中, 我们认为这种顺序是随机的, 从一个元素开始到后一个元素, 依次执行, 在 GO 的设计中, 这是有意为之的, 这样可以似的程序在不同的散列算法实现下变得健壮.

可以通过下面代码查看 map 的遍历顺序

[sequence.go](https://github.com/DestinyWang/gogl.io/blob/master/ch4/3_map/examples/sequence.go)

通过下标的方式访问 map 时, 总会有值, 如果键在 map 中, 可以得到键对应的值; 否则会得到 map 对应值类型的 `零值`, 如果需要知道 map 是否真的存在这个 key, 那么就需要多接收一个返回值

```go
age, ok := ages["dst"]
if ok {
    /// ~ 表示存在该 key, age 即为 map 中的值
} else {
    /// ~ 表示不存在该 key, age 为 int 型的 零值
}
```

> 有时候我们需要一个 map 并且需求它的键是 slice, 但是因为 map 的 key 必须是可以比较的, 所以这个功能无法直接实现, 但我们可以分成两步:
第一步, 定义一个帮助函数 k, 将每个键都映射到字符串, 当且仅当 x 和 y 相等的时候, 我们才认为 k(x) == k(y);
第二步, 可以创建一个 key 为 string 的 map, 在每个 key 被访问的时候, 调用这个帮助函数.

下面的例子通过一个字符串序列表示 map 来记录 add 函数被调用的次数, 帮助函数用 `fmt.Sprintf` 来将一个字符串 slice 转换为一个适合做 map 的 key 的字符串, 使用 `%q` 来格式化 slice 并记录每个字符串的边界.

```go
var m = make(map[string]int)

// 将 []string 转换为格式化的 string
func k(list []string) string {
    return fmt.Sprintf("%q", list)
}

// 以被格式化的 []string 为 key 进行自增, 来统计调用次数
func add(list []string) {
    m[k(list)]++
}

func Count(list []string) int {
    return m[k(list)]
}
```

同样的方法适用于任何不可直接比较的 key 类型, 不仅仅局限于 slice, 甚至有的时候不希望通过 `==` 来比较相等性, 而是一种自定义的比较方法, 例如字符串不区分大小写的比较, 同样 k(x) 的类型不一定的字符串类型, 任何能够想得到的比较结果的可比较类型都可以.