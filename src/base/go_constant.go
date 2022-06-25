package main

/**
 * @author      weimenghua
 * @time        2022-06-25 12:52
 * @description 常量：常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
 */

import (
	"fmt"
	_ "fmt"
)

func main() {
	const a bool = true
	const b int = 1
	const c float32 = 1.1
	const d float64 = 1.1
	const e complex128 = complex(1, 2)
	const f string = "字符串"

	fmt.Println(a, b, c, d, e, f)
}
