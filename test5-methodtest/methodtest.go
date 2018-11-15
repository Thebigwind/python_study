/*
在golang的世界中，一定要区分 方法和函数。

Go中没有类的概念，但是我们可以在一些类型上定义一些方法，也就是所谓的方法，跟函数不同。
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
方法与函数的区别
方法和函数定义语法区别在于：
方法前置实例接受参数，这个receiver可以是基础类型也可以是指针。

func (d *duck) quack() { // receiver
	// do something
}

func quack(d *duck) { // funciton argument
	// do something
}
*/
type Rectangle struct {
	length, width int
}

func (r Rectangle) AreaByValue() int {
	return r.length * r.width
}

func (r *Rectangle) AreaByReference() int {
	return r.length * r.width
}

//类型*T方法集包含所有receiver T + *T 方法
type IFace interface {
	SetSomeField(newValue string)
	GetSomeField() string
}

type Implementation struct {
	someField string
}

func (i *Implementation) GetSomeField() string {
	return i.someField
}

func (i *Implementation) SetSomeField(newValue string) {
	i.someField = newValue
}

func Create() *Implementation {
	return &Implementation{someField: "Hello"}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println("---------------")
	//对于下面的代码，没有什么区别：
	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is:", r1)
	fmt.Println("Rectangle area is :", r1.AreaByValue())        //12
	fmt.Println("Rectangle area is :", r1.AreaByReference())    //12
	fmt.Println("Rectangle area is :", (&r1).AreaByValue())     //12
	fmt.Println("Rectangle area is :", (&r1).AreaByReference()) //12

	fmt.Println("-----------------")
	//类型*T方法集包含所有receiver T + *T 方法
	var a IFace
	a = Create()
	a.SetSomeField("World")
	fmt.Println(a.GetSomeField())
}
