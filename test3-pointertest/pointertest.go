package main

import (
	"fmt"
)

func main() {
	val := 1000
	updateValue(val)         //按值传递，传递的是变量的副本。
	fmt.Println("val:", val) //val: 1000

	res := updateValue2(val)
	fmt.Println("val:", res) //val: 1100

	fmt.Println("------------")
	val2 := 1000
	updateValue3(&val2)        //按指针传递，不是副本。
	fmt.Println("val2:", val2) //val2: 1100

	fmt.Println("------------")
	a1 := 1
	b1 := 2
	a1, b1 = swap1(a1, b1)
	fmt.Println("a1:", a1) //2
	fmt.Println("b1:", b1) //1

	fmt.Println("------------")
	a2 := 10
	b2 := 20
	swap2(&a2, &b2)
	fmt.Println("a2:", a2) //20
	fmt.Println("b2:", b2) //10

	fmt.Println("--------------")
	//通过new进行指针的创建

	p1 := 100
	p2 := new(float64) //通过new创建的指针
	updateValue4(&p1, p2)
	fmt.Println("p1:", p1)  //200
	fmt.Println("p2:", *p2) //1.75

	fmt.Println("--------------")
	//指针结构体
	goog := Stock{454.43, 421.01, 435.29}
	fmt.Println("Original Stock Data:", goog)
	modifyStock(&goog)
	fmt.Println("Modified Stock Data:", goog)

}

func updateValue(val int) {
	val = val + 100
}

func updateValue2(val int) int {
	val = val + 100
	return val
}

func updateValue3(someVal *int) {
	*someVal = *someVal + 100
}

func updateValue4(someVal *int, someVal2 *float64) {
	*someVal = *someVal + 100
	*someVal2 = *someVal2 + 1.75
}

func swap1(a int, b int) (int, int) {
	return b, a
}

func swap2(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

//指针结构体
type Stock struct {
	high  float64
	low   float64
	close float64
}

func modifyStock(stock *Stock) {
	stock.high = 475.10
	stock.low = 400.15
	stock.close = 450.75
}
