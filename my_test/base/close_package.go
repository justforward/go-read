package main

import "fmt"

func f(x int) func(int) int{
	g := func(y int) int{
		return x+y
	}
	// 返回闭包
	// g 是闭包的限定条件：上面的f()返回的g之所以称为闭包函数，是因为它是一个函数，且引用了不在它自己范围内的变量x，这个变量x是g所在作用域环境内的变量，因为x是未知、未赋值的自由变量。
	return g
}

func main() {
	// 将函数的返回结果"闭包"赋值给变量a
	a := f(3)
	// 调用存储在变量中的闭包函数
	res := a(5)
	fmt.Println(res)

	// 可以直接调用闭包
	// 因为闭包没有赋值给变量，所以它称为匿名闭包
	fmt.Println(f(3)(5))
}