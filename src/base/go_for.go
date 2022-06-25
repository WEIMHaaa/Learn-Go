package main

/**
 * @author      weimenghua
 * @time        2022-06-25 12:52
 * @description for语句
 */

import "fmt"

func main() {
	arr := [...]int{1, 3, 5}

	//传统for循环遍历
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	//for...range循环遍历
	for i, v := range arr {
		fmt.Println(i, v)
	}

	//for true {
	//	fmt.Print("这是无线循环 \n")
	//}
}
