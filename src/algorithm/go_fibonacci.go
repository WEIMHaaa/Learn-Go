package main

import "fmt"

/**
 * @author      weimenghua
 * @time        2022-06-27 15:08
 * @description 斐波那契
 */

func demo(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return demo(n-1) + demo(n-2)
}

func main() {
	n := 10
	num := demo(n)
	fmt.Println(n, num)
}
