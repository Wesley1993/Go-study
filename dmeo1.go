package main // 必须有一个main包

import "fmt"

func main() {
	// 赋值前，必须先声明变量
	var a int
	fmt.Println("a = ",a)

	// := 自动推导类型，先声明变量，再给变量赋值
	b := 20
	fmt.Println("b = ",b)
}

