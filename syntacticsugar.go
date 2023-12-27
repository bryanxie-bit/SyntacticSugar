package SyntacticSugar

import "fmt"

/////////////////////////////////////////////////////////////// 短变量命名  :=

var x  = 10  //合法
x := 10      //在函数外使用 := 不合法

func shortDo(){
	x := 20     //可以在新的作用域内重复声明
	if true {   //可以在短语句块中声明相同的名称，例如：if、for、switch 中，但它们有各自作用域。
		x := 2
		fmt.Printf("x = %d\n", x) // x = 2
	}
	fmt.Printf("x = %d\n", x) // x = 1
}

/////////////////////////////////////////////////////////////// new函数 以下两个函数等价

/*
func newInt() *int {
	return new(int)
}

func newInt() *int {
	var x int
	return &x
}
 */


/////////////////////////////////////////////////////////////// ...与切片

//Go 函数定义中，我们可以使用...表示可变参数，用于表示可以接受任意个数但相同类型的参数

func Slice(a,b int, c...int){  //可变长参数
	fmt.Println()   //可变参函数
	//初始化切片 将第1位赋值99, 第10位 赋值 100，其他位赋值0
	arr := [...]int{0:99,9:100}
	fmt.Println(arr)
}

func SliceDo(){
	Slice(1,2,3,4)  //可传多个参数
}

/////////////////////////////////////////////////////////////// 接收者方法

type Instance struct{}

func (ins *Instance) Foo1() string {
	return ""
}
func (ins Instance) Foo2() string {
	return ""
}
func Receive(){
	/*
	Ins 值属于 Instance 类型，而非 *Instance，却能调用 Foo 方法，这是为什么呢?这其实就是 Go 编译器提供的语法糖!
	当一个变量可变时，我们对类型 T 的变量直接调用 *T 方法是合法的，因为 Go 编译器隐式地获取了它的地址。
	变量可变意味着变量可寻址，因此，下面的 Instance{}.Foo1() 会得到编译错误，就在于 Instance{} 值不能寻址。
	 */
	var _ = Instance{}.Foo1()   //报错

	var _ = Instance{}.Foo2()

	var a Instance
	a.Foo1()
	a.Foo2()
}

/////////////////////////////////////////////////////////////// for-range

// 对于切片、数组、字符串，其 for range 遍历方式有三种

func ForSliceArrayString() {
	a := []int{1, 2, 3}

	// 遍历一：不关心索引和数据的情况
	for range a {
	}

	// 遍历二：只关心索引的情况
	for index := range a {
		fmt.Println(index)
	}

	// 遍历三：关心索引和数据的情况
	for index, value := range a {
		fmt.Println(index, value)
	}
}

//map 也有三种 for range 遍历方式

func ForMap(){
	m := map[int]string{1: "Golang", 2: "Python", 3: "Java"}
	// 遍历一：不关心 key 和 value 的情况
	for range m {
	}

	// 遍历二：只关心 key 的情况
	for key := range m {
		fmt.Println(key)
	}

	// 遍历二：关心 key 和 value 的情况
	for key, value := range m {
		fmt.Println(key, value)
	}
}

//chanel 有两种遍历方式

func ForChanel(){
	ch := make(chan int, 10)

	// 遍历一：不关心 channel 数据
	for range ch {
	}

	// 遍历二：关心 channel 数据
	for data := range ch {
		fmt.Println(data)
	}
}

/////////////////////////////////////////////////////////////// init 函数

/*
	需要关注init函数的执行顺序
 */




