package main

import "fmt"

var g_int int = 5

func main() {

	//定义模式: var name type
	var out_string string = "This is system"

	//多返回值 _占位符   := 就地初始化
	_, second_, third_ := three_func()

	//可变参数
	fmt.Print(second_, third_)

	//判断语句
	if third_ == 3 {
		fmt.Print("third == 3 \n")
	}

	//const string
	const const_str string = "This is const string"

	fmt.Print("\n")

	fmt.Print(out_string)

	//输出iota
	fmt.Print("\n Out_iota:\n")
	out_iota()

	//条件语句判断
	out_condition(1, 2)

	//匿名函数
	anonym()

	//闭包引用,由于引用变量存在于堆上，不会被释放
	f_closure := closure_func(1)
	fmt.Print("the sum One is ", f_closure(2, 3), "\n")
	fmt.Print("the sum Two is ", f_closure(2, 3), "\n")
	//闭包副本
	g_closure := f_closure
	fmt.Print("the sum Three is ", g_closure(2, 3), "\n")
	//闭包副本
	g_closure = closure_func(1)
	fmt.Print("the sum Four is ", g_closure(2, 3), "\n")

}

//多返回值
func three_func() (string, string, int) {

	first_, second_, third_ := "First"+"Parameter", "Second", 3

	return first_, second_, third_
}

//iota用法
func out_iota() {
	const (
		a = 1
		b
		c = iota
		d
		e
		f = "f"
		g = "g"
		h
		i = iota
		j
		k
	)

	fmt.Print(a, b, c, d, e, f, g, h, i, j, k)
}

//条件语句 if switch
func out_condition(first_int, second_int int) {
	first_str := "first condtion match"
	second_str := "second condtion match"
	the_same_str := "third condition match"

	fmt.Print("\n")

	if first_int > second_int {
		fmt.Print(first_str)
	} else if first_int < second_int {
		fmt.Print(second_str)
	} else {
		fmt.Print(the_same_str)
	}

	fmt.Print("\n")

	switch {
	case first_int > second_int:
		fmt.Print(first_str)
	case first_int < second_int:
		fmt.Print(second_str)
	default:
		fmt.Print(the_same_str)
	}

}

//匿名函数
var anonym = func() {

	fmt.Print("\n")
	fmt.Print("This is Anonym")
}

//闭包函数
/**
 * @biref 闭包对外部变量为直接引用，外部引用量直接被分配到堆上,可能会导致内存泄漏
 */
func closure_func(main_int int) func(closure_first, closure_second int) int {
	var sum = main_int
	return func(closure_first, closure_second int) int {
		sum += (closure_first + closure_second)
		return sum
	}
}
