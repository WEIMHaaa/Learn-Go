package main

/**
 * @author        weimenghua
 * @time          2022-06-25 12:40
 * @description	  变量
 */

import "fmt"

func main() {
	var a string = "demo" //声明一个变量并初始化
	var b int             //没有初始化就为零值
	var c bool            //bool 零值为 false
	var d float64
	fmt.Printf(a, b, c, d)
}
