// 联系 3.11
// 增强 comma 函数的功能, 让其正确处理浮点数, 以及带有可选正负号的数字
package main

import (
	"strconv"
)

// 先将 float64 转换为 string, 然后就可以按原有的逻辑格式化为字符串
func floatComma(f float64) string {
	// strconv.FormatFloat 参数列表:
	// f float64:   需要格式化的浮点数
	// fmt byte:    b - 无小数部分
	//              e - 科学计数法，例如 -1234.456e+78
	//              E - 科学计数法，例如 -1234.456E+78
	//              f - 有小数点而无指数，例如 123.456
	//              g - 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
	//              G - 根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出
	// prec int:    有效数字(对 fmt = 'b' 无效)
	// bitSize:     整数取值范围
	s := strconv.FormatFloat(f, 'f', -1, 64)
	if f > 0 {          // 对正负值做判断
		s = "+" + s
	}
	return comma(s)
}