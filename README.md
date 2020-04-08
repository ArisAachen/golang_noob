# golang_noob

//这是学习golang的基本说明
//具体示例代码看.go文件

#go 基本数据类型与初始化

bool byte int int8 int16 ... uintptr

float float64

复数 complex64 complex128

字符 rune string 

错误 error

初始化
var parameter_name type = value

paramater := value  自动语义判断

var v = complex(2.1,3)      a:=real(v) 实部   b:=image(v) 虚部


##array
数组(不可变长)
var arr_name [count] type 
array_int := [...]int{1,2,3,4} 自动长度
遍历数组
for k,v := range arr_name {

}
//长度
length := len(arr_name)


##slice
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

##map
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

#语句
##if语句
if x := func_exec();x < y{

} else if x < z{

}else {

}

##switch
多值swicth
switch i := "x";j{
    case "x","y":
        code_body

    case "x","z":
        code_body

    default:
       code_body
}

语句switch
switch {
    case x > y:
        code_body

    case x < y:
        code_body

    default:
        code_body
}

##for语句
for _,v := range array{

}
for i:=0;i<10;i++{

}


##Label goto




#func函数

多返回值
func func_name(parameter_list) (return list){
    func_body
}
func swap(a,b int)(int,int){
    return b,a
}

引用传递和值传递
func add_ptr(add_int int){
    *add_int = *add_int + 1
}

不定参数
func sum(arrary_int ...int) int{
    sum:=0
    for _,v := range arrary_int {
        sum += v
    }
    return sum
}
func main(){
    var_slice := []int{1,2,3,4,5}
    sum(var_slice...)
}

##函数签名
type func_type (int,int) int
func add(a_int,b_int int) int{
    return a_int + b_int
}
func do(f func_type,a_int,b_int int)int{
    return f(a_int,b_int)
}
func main(){
    fmt.Print(do(add,1,2))
}

函数是first-class

匿名函数
var sum = func(a,b int) int{
    return a + b
}

##defer延迟调用,保证释放资源
func copy_file(dst,file_path string) (w int,err error){
    file_str,err := os.Open(file_path)
    if err != nil{
        return 
    }

    defer file_str.Close()

    dst,error := os.Create(dst)

    if(error != nil){
        return
    }

    defer dst.Close()

    w,error = io.Copy(dst,file_str)

    return
}

##闭包
导致函数内局部变量逃逸到堆，不会被GC自动释放
可以访问函数内部变量
多次访问得到的是一个父函数的副本
多次访问函数对象，闭包函数共享外部引用
func closure_fuc(num_int int) func(add_int int)int{
    var sum = 0
    return func(add_int int)int{
        return sum += add_int
    }
}


##panic 和recover
//捕获异常
defer func(){
    fmt.Println("defer inner")
    println(recover())
}()

func test(){
    panic("test panic")
}

##Error 
type error interface{
    Error()string
}
是个接口，只要实现Error()string


#类型系统
命名类型
type Student struct{
    name string
    id   int
}
未命名类型 map slice array
a := struct{
    name string
    id   int
}{"Aris",25}

//self_map自定义底层是map,可以赋值和for
type self_map map[string]int
type grand_map self_map

func test_map(){
    father_map := map[id]string{"Aris":1}
    var son_map self_map = father_map
    for k,v := range son_map{

    }

    //底层虽然相同，但没有未命名类型
    var grand_son grand_map = son_map  //error 
    var grand_son grand_map = (self_map)son_map //强制类型转换
}

##类型方法
type newType oldType newType为命名类型，继承操作集合

//不常用的初始化
get_ptr = new(newType)
get_ptr = newType{}

//类型item
type Element struct{
    id int
    next,pre *Element   //指针
    list *list 
    Value interface{}  //接口
    *bool             //匿名字段
}

func (t Type_name)MethodName(parameter_list)(return_list){
    method_body
}
t是接收者  MethodName(t Type_name,parameter_list)(return_list){
    method_body
}

type Slice_Int []int
func (slice_int Slice_Int)Sum()int{
    sum := 0
    for _,v := range slice_int {
        sum += i
    }
}
类型方法规则:
1.非命名类型不能自定义方法，命名类型可以
2.方法的定义必须和类型的定义在同一个包
3.大写开头的方法能被包外访问，否则不能
4.自定义类型泵调用原有类型方法，但支持的运算可以被继承
