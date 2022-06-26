package main

import "fmt"

/**
 * @author      weimenghua
 * @time        2022-06-26 11:48
 * @description	不定长参数
 */

func myfunc(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}
