> JSON 是 JavaScript 值的 Unicode 编码, 这些值包括字符串, 数字, 布尔值, 数组和对象. JSON 是基本数据类型和符合数据类型的一种高效, 可读性强的表示方法.

JSON 对象用来编码 Go 中的 map(key 必须为 string 类型) 和结构体

| Go 基本类型 | JSON 实例 |
| --- | --- |
| bool | `true` |
| number | `-273.15` |
| string | `"My name is "\destiny\""` |
| array | `["Red", "Blue", "Black"]` |
| object | `{"name": "destiny", "age": 24, "skill": ["programming", "learning"]}` |

### 假设有一个程序需要手机电影的观看次数并提供推荐, Movie 类型和典型的元素列表都在下例中提供

[movie.go](https://github.com/DestinyWang/gopl.io/blob/master/ch4/5_JSON/examples/movie.go)

Marshal 使用 Go 结构体成员的名称作为 JSON 对象中字段的名称(通过反射的方式), 只有可导出的成员可以转化为 JSON 字段.

```go
type Movie struct {
    // Year 和 Color 后面的字符串字面量是成员的标签
    Title  string
    Year   int  `json:"released"`
    Color  bool `json:"color, omitempty"`
    Actors []string
}
```

> 上例中结构体成员 Year 对应地转化为 `released`, 另外 Color 转换为 `color`, 这是通过 `成员标签定义` 实现的, 成员标签定义是结构体成员在编译期间关联的一些元数据信息.
成员标签可以是任意字符串, 但习惯上使用由空格分开的标签键值对 `key: "value"` 组成, 因为标签值使用双引号括起来, 所以一般标签都是原生的字面量.
键 `json` 控制包 `encoding/json` 的行为, 同样起来的 `encoding/...` 包也遵循这个规则.

marshal 的逆操作将 JSON 字符串解码为 Go 数据结构, 这么操作叫做 `unmarshal`, 由 `json.Unmarshal` 实现.

```go
// 结构体唯一的成员是 Title, 通过合理定义 Go 的数据结构, 可以选择那部分 JSON 数据解码到结构体对象中, 哪些可以丢弃
var titles []struct {
    Title string
}

if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSON unmarshal failed: %s", err)
}
fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke}]"
```