# golang_noob

//这是学习golang的基本说明
//具体示例代码看.go文件

# go 基本数据类型与初始化

bool byte int int8 int16 ... uintptr

float float64

复数 complex64 complex128

字符 rune string 

错误 error

初始化
var parameter_name type = value

paramater := value  自动语义判断

var v = complex(2.1,3)      a:=real(v) 实部   b:=image(v) 虚部


## array
数组(不可变长)
var arr_name [count] type 
array_int := [...]int{1,2,3,4} 自动长度
遍历数组
``` go
for k,v := range arr_name {

}
//长度
length := len(arr_name)
``` 

## slice
切片  slice  可变长数组
本质上是一个结构
维护了一个unsafe.pointer  len cap
``` go
//var_name := make([]type,len,cap)
var_slice_one := make([]int,10,15)
var_slice_two := make([]int,2,2)
```
复制最小长度
``` go
copy(var_slice_two,var_slice_one)
//添加元素
var_slice_two.append(...)
//长度
len(var_slice_two)
//容量
cap(var_slice_two) 
```
扩容方式类似于vector cap的倍增

## map
``` go
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
```

# 语句
## if语句
``` go
if x := func_exec();x < y{

} else if x < z{

}else {

}
```

## switch
多值swicth
``` go
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
```

## for语句
``` go
for _,v := range array{

}
for i:=0;i<10;i++{

}
```

## Label goto




#func函数

多返回值
``` go
func func_name(parameter_list) (return list){
    func_body
}
func swap(a,b int)(int,int){
    return b,a
}
```
引用传递和值传递
``` go
func add_ptr(add_int int){
    *add_int = *add_int + 1
}
```

不定参数
``` go
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
```

## 函数签名
``` go
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
```

函数是first-class

匿名函数
``` go
var sum = func(a,b int) int{
    return a + b
}
```

## defer延迟调用,保证释放资源
``` go
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
```
## 闭包
导致函数内局部变量逃逸到堆，不会被GC自动释放
可以访问函数内部变量
多次访问得到的是一个父函数的副本
多次访问函数对象，闭包函数共享外部引用
``` go
func closure_fuc(num_int int) func(add_int int)int{
    var sum = 0
    return func(add_int int)int{
        return sum += add_int
    }
}
```

## panic 和recover
``` go
//捕获异常
defer func(){
    fmt.Println("defer inner")
    println(recover())
}()

func test(){
    panic("test panic")
}
```
## Error 
``` go
type error interface{
    Error()string
}
```
是个接口，只要实现Error()string


# 类型系统
命名类型
``` go
type Student struct{
    name string
    id   int
}
```
未命名类型 map slice array
``` go
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
```
## 类型方法
type newType oldType newType为命名类型，继承操作集合
``` go
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
```
类型方法规则:
1.非命名类型不能自定义方法，命名类型可以
2.方法的定义必须和类型的定义在同一个包
3.大写开头的方法能被包外访问，否则不能
4.自定义类型泵调用原有类型方法，但支持的运算可以被继承

``` go
//方法调用
type Add_Obj struct{
    var sum int = 0
}
func (self_add Add_Obj) add(num_int int)int{
    sum  += num_int
    return sum
}
```

## 类型嵌套   组合(类继承)
``` go
struct Inner{
    inner_int int
}

struct Outter{
    Inner
    outter_int int
}

此时Outter outter  
outter.inner_int = outter.Inner.inner_int
同样可以Inner inner{5}
Outter{
    Inner:inner
    outter_int : 5
}

方法重写
func (inner Inner)Print(){
    fmt.Print(inner.inner_int)
}

func (outter Outter)Print(){
    fmt.Print(outter.outter_int)
}
outter.Print()优先从外层向内层寻找
```

组合方法集规则:
1.若类型包含匿名字段S，则包含匿名字段S的方法集
2.若类型包含匿名字段S*,则包含匿名字段S*和S的方法集
3.T*方法集包含是S和S*的方法集

#函数类型
函数字面量类型  有名函数(匿名函数赋值)  匿名函数
``` go
var func_known = func (num_int int) int{} = func func_known (num_int) int{} //有名函数
func (num_int int) int{}  //匿名函数
```
函数命名类型  type func_name func(parameter_list) (return_list)

//例子
``` go
func add(num_one,num_two int) int{
    return num_one + num_two
}

type ADD func(num_first,num_second int)int

func main(){
    var add_local ADD = add

    println(add_local(3,4))
}
```


# 接口
空接口  interface{}
``` go
//接口声明
type Reader interface{
    Read(p []byte)(n int,err error)
}
type Writer interface{
    Write(p []byte)(n int,err error)
}
type ReaderAndWriter interface{
    Reader
    Writer
}

//超集实现
type IOReader struct{}
func (ioreader IOReader)Read(p []byte)(n int,err error){
    file,error := os.OpenFile("note.txt",os.O_RDWR|os.O_CREATE)

    if(error != nil){
        return 0,error
    }
    file.Close()
    return 1,nil
}

//编译时期静态检查
func main(){
    var reader Reader;
    reader = IOReader{}
    reader.Read(...)
}
```

## 类型判断和接口类型查询

### 类型判断
接口断言的语法表现
直接赋值模式   o := i.(Type_name)
``` go
comma,ok表达式   if o,ok := i.(Type_name);ok{

}
type Integer interface{
    Add(add_num int)int
    Minus(min_num int)int
}
type Integer_Max interface{
    Integer 
    Multiple(mul_num int)int
}
type Integer_Complement struct{
    name_str string
    sum      int
}
func (integer Integer_Complement) Add(add_num int)int{
    return integer.sum += add_num
}
func (integer Integer_Complement) Minus(min_num int)int{
    return integer.sum -= min_num
}
func main(){
    integer := &Integer_Complement{"Integer_Complement",0}
    var inter_interface = integer

    //断言判断
    o := inte_interface.(Integer)
    o.Add(5)

    if o,ok := inte_interface.(Integer_Max);ok{
        //没有实现Integer_Max，语句不会被执行
    }
}
```

### 类型查询
``` go
func main(){
    f,err := os.OpenFile("note.txt",os.O_RDWR|os.O_CREATE)

    if err != nil{
        return
    }

    defer f.Close()

    var i io.Reader = f

    switch v:=i.(type){
        case *os.File:
            //func_body
        default:
            //func_body
    }
}
```

接口优点:
解耦     将空接口传递作为泛型

###空接口  作为泛型
空接口有两个字段，一个是实例类型，另一个是指向绑定实例的指针
两个都为nil  空接口才为nil
``` go
type nil_interface interface{
    Normal_nil()
    Pointer_nil()
}
type nil_implement struct{
}
func (nil_inter nil_implement)Normal_nil(){
    println("normal_func")
}
func (nil_inter *nilimplement)Pointer_nil(){
    println("pointer_func")
}
func main(){
    var nil_imp *nil_implement = nil
    var nil_inter nil_interface = nil_imp

    println("%p",nil_inter) //地址为0x0

    if nil_inter == nil  //值为false  nil_inter != nil

    nil_inter.Normal_nil() //Error 实例为nil

    nil_inter.Pointer_nil() //输出pointer_nil
}
```

# goroutine 和chan
go不推荐使用内存进行线程通信，使用chan类似于管道

特性
1.go的执行是非阻塞的，不会等待
2.go的返回值会被忽略
3.调度器不能保证多个goroutine的执行次序
4.没有父子goroutine的概念，所以goroutine被平等地调度
5.执行时单独为main创建一个goroutine
6.不会暴露goroutine id，不能在其他的goroutine显式操作其他的

### func GOMAXPROC
``` go
func main(){
    println(runtime.GOMAXPROC(0))    //小于1  显示当前数量
    runtime.GOMAXPROC(2)    //大于1   设置最大数量

    go func(){
        println("runtime")

    }

    time.Sleep(5 * time.Second)   //防止提前退出
}
```

### chan通道
无缓存通道，可以用于通信和goroutine同步
``` go
func main(){
    c := make(chan struct{})
    
    defer close(c)

    gon func(i chan struct {}){

        println("chan on")

        //写通道c
        c <- struct{}{}
    }(c)

    //读通道c
    get_c := <-c
}
```
通道使用后要关闭

### WaitGroup
main goroutine调用Add 设置需要等待goroutine的数目，完成goroutine调用Done()
Wait()等待所以goroutine结束
``` go
var wg sync.WaitGroup
func main(){

    for i:=0;i<10;i++{
        wg.Add(1)
        go wait_group_test()
    }

    wg.Wait()
}

func wait_group_test(){
    println("wait group test")

    defer wg.Done()
}
```

### select
类似于unix多路复用   select一直循环
``` go
func main(){
    chanSelect := make(chan int)

	go selectTest(chanSelect)

	for i := 0; i < 10; i++ {
		println("select num = ", <-chanSelect)
	}
}

func selectTest(chanInt chan int) {

	for {
		select {
		case chanInt <- 0:

			println("chanInt <- 0")
		case chanInt <- 1:
			println("chanInt <- 1")

		}
	}
}
```

## 通知退出机制
关闭select监听的通道，使得select立马感知到关闭
``` go
func main(){
    	//退出机制
	chanOut := make(chan int)
	chanRand := selectClose(chanOut)

	println(<-chanRand)

	close(chanOut)

	println(<-chanRand)
}

func selectClose(chanInt chan int) chan int {

	ch := make(chan int)

	go func() {

	Label:
		for {

			select {
			case ch <- rand.Int():
				println("selectClose random")
			case <-chanInt:
				println("selectClose breakout")
				close(ch)
				break Label
			}

		}
	}()
	return ch
}
```