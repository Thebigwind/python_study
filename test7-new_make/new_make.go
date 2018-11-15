/*
区别
make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。

内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，
并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：new返回指针。
下面两种形式效果一样：
func newInt1() *int { return new(int) }

func newInt2() *int {
    var i int
    return &i
}

内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，
并且返回一个有初始值(非零)的T类型，而不是*T。

These examples illustrate the difference between new and make.

var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
	var p *[]int = new([]int)
	*p = make([]int, 100, 100)

// Idiomatic:
	v := make([]int, 100)
*/