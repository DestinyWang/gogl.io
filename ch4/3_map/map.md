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