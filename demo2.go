package main

import "fmt"

func main() {
	a := 10
	//一段一段输出，自动加换行
	fmt.Println("a = ", a)

	//格式化输出，把内容放在%d的位置
	fmt.Printf("a = %d \n", a)

	b := 20
	c := 30
	fmt.Printf("a = %d,b = %d,c = %d\n", a, b, c)
}

//  Printf 和 Println的区别
