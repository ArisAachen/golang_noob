package main

import "fmt"

func main() {

	//数组
	array_fuc()

	//指针
	pointer_func()

	//结构体
	structure_func()

	//切片
	slice_func()

	//map
	map_func()

	//接口
	var inter_paramter Type_interface
	inter_paramter = Type_Implement{}
	inter_paramter.Print()

}

//数组变量
func array_fuc() {
	//var array_int[10] int

	fmt.Print("Start Print Array... \n")

	var array_int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < 10; i++ {
		fmt.Print(array_int[i], " ")
	}

	fmt.Print("\n")

}

//指针
func pointer_func() {
	var var_int int = 5
	var var_ptr *int = &var_int

	if var_ptr != nil {
		fmt.Print("var_int is at ", var_ptr)
	}

	fmt.Print("\n")

	//指针无法数增 vat_ptr++ error
}

//结构体
func structure_func() {

	type Type_book struct {
		title   string
		author  string
		book_id int
	}

	var one_book = Type_book{"Golang", "Aris", 100}

	fmt.Print(one_book)

	fmt.Print("\n")
}

//切片
func slice_func() {

	var_slice_one := [5]int{1, 2, 3, 4, 5}

	var_slice_two := make([]int, 5, 10)

	fmt.Print("Start Print Slice... \n")

	for k, v := range var_slice_one {
		fmt.Print(k, ":", v)
		fmt.Print("\n")
	}

	//打印长度
	fmt.Print("cap of var_slice_two is ", cap(var_slice_two), "\n")

	var_slice_two = append(var_slice_two, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	//打印长度 切片倍增
	fmt.Print("cap of var_slice_two is ", cap(var_slice_two), "\n")

}

func map_func() {
	id_name_map := make(map[int]string)

	id_name_map[1] = "Aris"
	id_name_map[2] = "Jack"
	id_name_map[5] = "John"

	fmt.Print("Start Print Map... \n")

	delete(id_name_map, 2)

	for k, v := range id_name_map {
		fmt.Print(k, ":", v, "\n")
	}

	//结构体map整体赋值
	/**
	map[int]struct
	map[1].item 修改error
	*/
}

type Type_interface interface {
	Print()
}

type Type_Implement struct {
}

func (implement Type_Implement) Print() {
	fmt.Print("This is implement")
}
