package main

import "fmt"

func main() {
	var x [5]int
	//通过下标修改数组
	x[3] = 100
	fmt.Println(x) //[0 0 0 100 0

	//数组的长度：
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += float64(x[i]) //因为total是float64，x[i]是int，所以需要显示的进行转换：
	}
	fmt.Println(total / float64(len(x))) //因为total是float64，len(x)是int，所以需要显示的进行转换

	fmt.Println("-------------------")

	//range在数组中的使用
	var total2 float64 = 0
	for _, value := range x {
		total2 += float64(value)
	}
	fmt.Println(total2 / float64(len(x)))

	fmt.Println("------------------")

	x2 := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	var total3 float64 = 0
	for _, value := range x2 {
		total3 += value
	}
	fmt.Println(total3 / float64(len(x2))) //86.6

	fmt.Println("------------------")

	/*
		slice
	*/
	//通过make进行创建：
	//append函数
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2) //[1 2 3] [1 2 3 4 5]

	//copy函数
	fmt.Println("-----------------")
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4) //[1 2 3] [1 2]

	fmt.Println("-------------------")

	// Here we create an array `a` that will hold exactly
	// 5 `int`s. The type of elements and length are both
	// part of the array's type. By default an array is
	// zero-valued, which for `int`s means `0`s.
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The builtin `len` returns the length of an array.
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array
	// in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
