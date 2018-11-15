package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	/*
		//change var value
		i, j := 0, 1
		i, j = j, i
		fmt.Println("i:", i)
		fmt.Println("j:", j)
	*/

	/*
		//指针操作
		x := 1
		p := &x
		fmt.Println(p)  //0xc04200a2a0 内存地址
		fmt.Println(*p) //1	地址中的值
		*p = 2
		fmt.Println(p) //0xc04200a2a0 内存地址
		fmt.Println(x) //2 地址中的值
		*p++           //递增p所指向的值，p自身保持不变
		fmt.Println(x) //3,
		fmt.Println(p) //0xc04200a2a0
	*/

	/*
		//指针是可比较的，两个指针指向同一个变量或两个指针都是nil的情况下才相等
		var x, y int
		fmt.Println(&x == &x, &x == &y, &x == nil) //true,false,false
		fmt.Println(&x, &y)                        //0xc04200a2a0 0xc04200a2a8
		fmt.Println(x)
	*/

	/*
		//指针使用程序的命令行参数来设置整个程序内某些变量的值
		echo4()
	*/
	/*
		//new函数	另一种创建变量的方式,和取地址的普通变量一样，只是不需要引入一个虚拟的名字
		p := new(int)
		fmt.Println(*p) // 0
		*p = 2
		fmt.Println(*p) //2
		//每次调用new返回一个具有唯一地址的不同变量	例外：两个变量的类型不携带任何信息且是零值，如struct{}或[0]int
		p1 := new(int)
		q1 := new(int)
		fmt.Println(p1) //0xc042036228
		fmt.Println(q1) //0xc042036230

		p2 := new(struct{})
		q2 := new(struct{})
		fmt.Println(p2) //&{}
		fmt.Println(q2) //&{}

		p3 := new([0]int)
		q3 := new([0]int)
		fmt.Println(p3) //&[]
		fmt.Println(q3) //&[]
	*/

	/*
		//变量的生命周期：变量在程序执行过程中的存在时间。一直生存到它不可访问为止，这时它占用的存储空间被收回
		//函数的参数和返回值也是局部变量，他们在闭包函数被创建的时候创建
		//编译器可以选择使用堆或栈上的空间来分配
	*/

	//赋值
	/*
		x = 1
		*p = true
		person.name = "bob"
		connt[x] = count[x] + scale

		v++
		v--
	*/
	/*
		x := gcd(10, 7)
		fmt.Println(x)
	*/

	//可赋值性 左边的变量和右边的值类型相同，必须精确匹配，nil可以赋值给任何接口变量或引用类型

	//类型声明 type声明定义一个新的命名类型，它和已有的类型使用同样的底层类型 type name underlying-type
	//从浮点型转化为整型会丢失小数部分，从字符串转化为字节（[]byte）slice会分配一份字符串数据副本。
	//命名类型的底层类型决定了它的结构和表达方式，以及它支持的内部操作的集合

	//包 用于支持模块化、封装、编译隔离和重用
	//标识符是否对外可见：导出的标识符以大写字母开头

	//值传递
	val := 1000
	updateValue(val)
	fmt.Println("val:", val) //按值传递，传递的是变量的副本。 1000
	//指针传递
	updateValue2(&val)
	fmt.Println("val:", val) //按指针传递，不是副本。 1100
	//new创建指针
	val1 := 1000
	val2 := new(float64)
	updateValue3(&val1, val2)
	fmt.Println("val:", val1)
	fmt.Println("val2:", *val2)

	goog := Stock{454.43, 421.01, 435.29}
	fmt.Println("Original Stock Data:", goog)
	modifyStock(&goog)
	fmt.Println("Modified Stock Data:", goog)
}

func updateValue(val int) {
	val = val + 100
}

func updateValue2(someVal *int) {
	*someVal = *someVal + 100
}

func updateValue3(someVal *int, someVal2 *float64) {
	*someVal = *someVal + 100
	*someVal2 = *someVal2 + 1.75
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

//结构体指针
type Stock struct {
	high  float64
	low   float64
	close float64
}

func modifyStock(stock *Stock) {
	stock.high = 147.0
	stock.low = 141.0
	stock.close = 132.3
}

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func echo4() {
	flag.Parse() //更新标识变量的默认值
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("xxr")
	}
}

//下面两个newInt()效果一样
func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

var global *int

//使用堆分配空间，因为x从f()中逃逸
func f() {
	var x int
	x = 1
	global = &x
}

//使用栈分配空间，因为*y没有从g()中逃逸
func g() {
	y := new(int)
	*y = 1
}

//最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
