# golang_noob

//这是学习golang的基本说明
//具体示例代码看.go文件

##go 基本数据类型与初始化

bool byte int int8 int16 ... uintptr

float float64

复数 complex64 complex128

字符 rune string 

错误 error

初始化
var parameter_name type = value

paramater := value  自动语义判断

var v = complex(2.1,3)      a:=real(v) 实部   b:=image(v) 虚部

数组(不可变长)

var arr_name [count] type 

array_int := [...]int{1,2,3,4} 自动长度

遍历数组
for k,v := range arr_name {

}

//长度
length := len(arr_name)


切片  slice  可变长数组
本质上是一个结构
维护了一个unsafe.pointer  len cap

//var_name := make([]type,len,cap)
var_slice_one := make([]int,10,15)

var_slice_two := make([]int,2,2)

复制最小长度
copy(var_slice_two,var_slice_one)

//添加元素
var_slice_two.append(...)

//长度
len(var_slice_two)

//容量
cap(var_slice_two) 

扩容方式类似于vector cap的倍增

map map[key_type]value_type

var_map := map[int]string{1:"Aris",2:"Jack"}

var_map[1] = "Change"

//遍历
for k,v := range var_map{
    //删除元素
    delete(var_map,k)
}

map不能修改item元素的值
type User struct{
    name string
    age int
}

User_map := map[int]User{1:{"Aris",25}}

//User_map[0].age = 19 Eroor

aris := User_map[1]
aris.int = 26
User_map[0] = aris


