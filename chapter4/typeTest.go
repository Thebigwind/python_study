package main

import (
	"crypto/sha256"
	"fmt"
	//"strings"
)

func main() {
	//utf8Test()
	//arrayTest()
	//sliceTest()
	/*
		a := [...]int{1, 2, 3, 45}
		reverse(a[:])
		fmt.Println(a) //[45 3 2 1]反转数组
	*/
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func utf8Test() {
	for i := 10000; i < 10001; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
	fmt.Println()
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}

//数组长度固定，元素数据类型相同
//slice 长度可变
func arrayTest() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	//输出索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	//仅输出元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	//数组长度由初始化个数决定
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // [3]int

	//数组长度是数组表达式的一部分，索引 [3]int 和 [4]int是两种不同的数组类型。数组的长度必须是常量表达式，这个值在程序编译时就可以决定
	//q1 := [3]int{1, 2, 3}
	//q1 = [4]int{1, 2, 3, 4}  //cannot use [4]int literal (type [4]int) as type [3]int in assignment

	//如果一个数组元素是可以比较的，那么这个数组也是可以比较的。使用 == 比较两个数组，比较的结果是数组元素的值是否完全相同
	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(a1 == b1, a1 == c1, b1 == c1)

	//
	t1 := sha256.Sum256([]byte("X"))
	t2 := sha256.Sum256([]byte("x"))

	fmt.Println(t1)
	fmt.Printf("%X\n", t1) //将数组或slice中值按16进制输出：2D711642B726B04401627CA9FBAC32F5C8530FB1903CC4DB02258717921A4881
	fmt.Println(t2)
	fmt.Printf("%X\n", t2)
	fmt.Printf("%T\n", t1) //输出一个类型：[32]uint8

}

//数组指针 :显示的传递一个数组的指针给函数，在函数内部对数组的修改都会反映到原始数组上
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

//slice:一个拥有相同类型元素的可变序列，拥有3个属性：指针、长度、容量
//go中内置函数len和cap返回slice的长度和容量
func sliceTest() {
	months := [...]string{1: "January", 2: "Feb", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "Oct", 11: "Nov"}
	fmt.Println(months)

	//slice操作符s[i:j]创建了一个隐形的slice,引用了序列s中从i到j-1索引位置的所有元素，s可以是数组，指向数组的指针，或slice，新的slice共j-i个元素
	Q2 := months[3:8]
	summer := months[6:9]
	fmt.Println(Q2)     //[March April May June July]
	fmt.Println(summer) //[June July August

	//输出两个slice中共同元素
	for _, q := range Q2 {
		for _, s := range summer {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	//如果slice的引用超过了被引用对象的容量，即cap(s)，会宕机；被引用的即months,从slice引用下标开始，到months的结束
	//如果slice的引用超过了被引用对象的长度，即len(s),最终slice会比原slice长。
	//fmt.Println(summer[:5]) //panic: runtime error: slice bounds out of rang
	endlessSummer := summer[:5] //slice容量范围内扩展了slice
	fmt.Println(endlessSummer)  //[June July August September Oct]

}

//slice包含了指向数组的指针，所以讲一个slice传递给函数，可以在函数内部修改底层数组的元素。
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
//slice没有指定长度，这种隐式区别的结果分别是创建有固定长度的数组和创建指向数组的slice
//slice无法比较，因此无法使用 == 比较两个slice是否元素相同。标准库提供了bytes.Equal来比较两个字节slice(slice []byte).但是对其它类型的slice,必须自己写函数来比较
func equal(x, y []stirng) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
*/

/*
测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断。
除了和nil相等比较外，一个nil值的slice的行为和其它任意0长度的slice一样；例如reverse(nil)也是安全的
*/

/*
内置的make函数创建一个指定元素类型、长度和容量的slice。
容量部分可以省略，在这种情况下，容量将等于长度。

make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]

在底层，make创建了一个匿名的数组变量，然后返回一个slice；只有通过返回的slice才能引用底层匿名的数组变量。
在第一种语句中，slice是整个数组的view。
在第二个语句中，slice只引用了底层数组的前len个元素，但是容量将包含整个的数组。额外的元素是留给未来的增长用的。
*/

//内置的append函数用于向slice追加元素：
func test3() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
}

/*
每次调用appendInt函数，必须先检测slice底层数组是否有足够的容量来保存新添加的元素。
如果有足够空间的话，直接扩展slice（依然在原有的底层数组之上），将新添加的y元素复制到新扩展的空间，并返回slice。
因此，输入的x和输出的z共享相同的底层数组。

如果没有足够的增长空间的话，appendInt函数则会先分配一个足够大的slice用于保存新的结果，先将输入的x复制到新的空间，然后添加y元素。
结果z和输入的x引用的将是不同的底层数组

虽然通过循环复制元素更直接，不过内置的copy函数可以方便地将一个slice复制另一个相同类型的slice。
copy函数的第一个参数是要复制的目标slice，第二个参数是源slice，目标和源的位置顺序和dst = src赋值语句是一致的。
两个slice可以共享同一个底层数组
*/
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

/*
0 cap=1	[0]
1 cap=2	[0 1]
2 cap=4	[0 1 2]
3 cap=4	[0 1 2 3]
4 cap=8	[0 1 2 3 4]
5 cap=8	[0 1 2 3 4 5]
6 cap=8	[0 1 2 3 4 5 6]
7 cap=8	[0 1 2 3 4 5 6 7]
8 cap=16	[0 1 2 3 4 5 6 7 8]
9 cap=16	[0 1 2 3 4 5 6 7 8 9]

让我们仔细查看i=3次的迭代。当时x包含了[0 1 2]三个元素，但是容量是4，因此可以简单将新的元素添加到末尾，不需要新的内存分配。
然后新的y的长度和容量都是4，并且和x引用着相同的底层数组。
在下一次迭代时i=4，现在没有新的空余的空间了，因此appendInt函数分配一个容量为8的底层数组，
将x的4个元素[0 1 2 3]复制到新空间的开头，然后添加新的元素i，新元素的值是4。
新的y的长度是5，容量是8；后面有3个空闲的位置，三次迭代都不需要分配新的空间。当前迭代中，y和x是对应不同底层数组的view

尽管底层数组的元素是间接访问的，但是slice对应结构体本身的指针、长度和容量部分是直接访问的。要更新这些信息需要像上面例子那样一个显式的赋值操作。从这个角度看，slice并不是一个纯粹的引用类型，它实际上是一个类似下面结构体的聚合类型：

type IntSlice struct {
    ptr      *int
    len, cap int
}
*/

//要删除slice中间的某个元素并保存原有的元素顺序，可以通过内置的copy函数将后面的子slice向前依次移动一位完成：
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

//如果删除元素后不用保持原来顺序的话，我们可以简单的用最后一个元素覆盖被删除的元素：
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

/*

==================================================================================================================



==================================================================================================================
*/

/*
内置的make函数可以创建一个map：

ages := make(map[string]int) // mapping from strings to ints
我们也可以用map字面值的语法创建map，同时还可以指定一些最初的key/value：

ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
这相当于

ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34

禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效
*/

/*
for name, age := range ages {
    fmt.Printf("%s\t%d\n", name, age)
}
Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序

如果要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序。下面是常见的处理方式：

import "sort"

var names []string
for name := range ages {
    names = append(names, name)
}
sort.Strings(names)
for _, name := range names {
    fmt.Printf("%s\t%d\n", name, ages[name])
}
*/

/*
map类型的零值是nil，也就是没有引用任何哈希表。

var ages map[string]int
fmt.Println(ages == nil)    // "true"
fmt.Println(len(ages) == 0) // "true"

map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。
但是向一个nil值的map存入元素将导致一个panic异常：

ages["carol"] = 21 // panic: assignment to entry in nil map

通过key作为索引下标来访问map将产生一个value。
如果key在map中是存在的，那么将得到与key对应的value；如果key不存在，那么将得到value对应类型的零值
*/

//要判断两个map是否包含相同的key和value，我们必须通过一个循环实现：

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

//读取多行输入，但是只打印第一次出现的行
/*
func dedup() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
*/

/*

******************************************************************************************************************
**********************************************struct**************************************************************
******************************************************************************************************************
 */

/*
结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员.
信息都需要绑定到一个实体中，可以作为一个整体单元被复制，作为函数的参数或返回值，或者是被存储到数组中

对成员取地址，然后通过指针访问：

position := &dilbert.Position
*position = "Senior " + *position // promoted, for outsourcing to Elbonia
点操作符也可以和指向结构体的指针一起工作：

var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"

*/

/*
一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身
但是S类型的结构体可以包含*S指针类型的成员，这可以让我们创建递归的数据结构，比如链表和树结构等
*/
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

/*
结构体类型的零值是每个成员都是零值。通常会将零值作为最合理的默认值
如果结构体没有任何成员的话就是空结构体，写作struct{}。它的大小为0，也不包含任何信息

考虑效率的话，较大的结构体通常会用指针的方式传入和返回，

func Bonus(e *Employee, percent int) int {
    return e.Salary * percent / 100
}

如果要在函数内部修改结构体成员的话，用指针传入是必须的；因为在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。

func AwardAnnualRaise(e *Employee) {
    e.Salary = e.Salary * 105 / 100
}
因为结构体通常通过指针处理，可以用下面的写法来创建并初始化一个结构体变量，并返回结构体的地址：

pp := &Point{1, 2}
它是下面的语句是等价的

pp := new(Point)
*pp = Point{1, 2}
*/

/*
结构体比较
如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用==或!=运算符进行比较
*/

/*
******************************************************************************************************************
**********************************************json**************************************************************
******************************************************************************************************************
 */
/*
JSON
基本的JSON类型有数字（十进制或科学记数法）、布尔值（true或false）、字符串，
其中字符串是以双引号包含的Unicode字符序列，支持和Go语言类似的反斜杠转义特性

这些基础类型可以通过JSON的数组和对象类型进行递归组合。
一个JSON数组是一个有序的值序列，写在一个方括号中并以逗号分隔;可以用于编码Go语言的数组和slice。
一个JSON对象是一个字符串到值的映射，写成以系列的name:value对形式，用花括号包含并以逗号分隔；可以用于编码Go语言的map类型（key类型是字符串）和结构体。

*/
/*
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}


//这样的数据结构特别适合JSON格式，并且在两种之间相互转换也很容易。
//将一个Go语言中类似movies的结构体slice转为JSON的过程叫编组（marshaling）。编组通过调用json.Marshal函数完成：
data, err := json.Marshal(movies)
if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
*/

/*
//对应将JSON数据解码为Go语言的数据结构，Go语言中一般叫unmarshaling，通过json.Unmarshal函数完成
var titles []struct{ Title string }
if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
*/
