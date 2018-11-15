package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Print("")

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println("-------------------------------------")

	fmt.Println(&x)         // "{1 9 42 144}"  这个类型的指针确实有自定义的String方法
	fmt.Println(x.String()) // "{1 9 42 144}" 直接调用了x变量的String()方法；这种情况下编译器会隐式地在x前插入&操作符，这样相当远我们还是调用的IntSet指针的String方法
	fmt.Println(x)          // "{[4398046511618 0 65536]}"  ，因为IntSet类型没有String方法，所以Println方法会直接以原始的方式理解并打印
}

/*
一个对象其实也就是一个简单的值或者一个变量，在这个对象中会包含一些方法，而一个方法则是一个和特殊类型关联的函数。
一个面向对象的程序会用方法来表达其属性和对应的操作，这样使用这个对象的用户就不需要直接去操作对象，而是借助方法来做这些事情。
OOP编程的两个关键点，封装和组合。
*/

/*
在函数声明时，在其名字之前放上一个变量，即是一个方法。这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。
*/

/*
在方法调用过程中，接收器参数一般会在方法名之前出现。这和方法声明是一样的，都是接收器参数在方法名字之前。下面是例子：

p := Point{1, 2}
q := Point{4, 6}
fmt.Println(Distance(p, q)) // "5", function call
fmt.Println(p.Distance(q))  // "5", method call

面的两个函数调用都是Distance，但是却没有发生冲突。
第一个Distance的调用实际上用的是包级别的函数geometry.Distance，
第二个则是使用刚刚声明的Point，调用的是Point类下声明的Point.Distance方法。

因为每种类型都有其方法的命名空间，我们在用Distance这个名字的时候，不同的Distance调用指向了不同类型里的Distance方法
*/

/*
基于指针对象的方法

当调用一个函数时，会对其每一个参数值进行拷贝，如果一个函数需要更新一个变量，或者函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，这种情况下我们就需要用到指针了。
对应到我们这里用来更新接收器的对象的方法，当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法，如下：

func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}
这个方法的名字是(*Point).ScaleBy。这里的括号是必须的；没有括号的话这个表达式可能会被理解为*(Point.ScaleBy)

在现实的程序里，一般会约定如果Point这个类有一个指针作为接收器的方法，那么所有Point的方法都必须有一个指针接收器，即使是那些并不需要这个指针接收器的函数。
只有类型(Point)和指向他们的指针(*Point)，才是可能会出现在接收器声明里的两种接收器。
在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的。




在每一个合法的方法调用表达式中，也就是下面三种情况里的任意一种情况都是可以的：

不论是接收器的实际参数和其接收器的形式参数相同，比如两者都是类型T或者都是类型*T：

Point{1, 2}.Distance(q) //  Point
pptr.ScaleBy(2)         // *Point
或者接收器实参是类型T，但接收器形参是类型*T，这种情况下编译器会隐式地为我们取变量的地址：

p.ScaleBy(2) // implicit (&p)
或者接收器实参是类型*T，形参是类型T。编译器会隐式地为我们解引用，取到指针指向的实际变量：

pptr.Distance(q) // implicit (*pptr)






1.不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。
2.在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的内部：
第一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用会产生一次拷贝；
第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向的始终是一块内存地址，就算你对其进行了拷贝。

*/

/*
通过嵌入结构体来扩展类型

import "image/color"

type Point struct{ X, Y float64 }

type ColoredPoint struct {
    Point
    Color color.RGBA
}

我们完全可以将ColoredPoint定义为一个有三个字段的struct，但是我们却将Point这个类型嵌入到ColoredPoint来提供X和Y这两个字段
我们可以直接认为通过嵌入的字段就是ColoredPoint自身的字段，而完全不需要在调用时指出Point，比如下面这样。

var cp ColoredPoint
cp.X = 1
fmt.Println(cp.Point.X) // "1"
cp.Point.Y = 2
fmt.Println(cp.Y) // "2"

对于Point中的方法我们也有类似的用法，我们可以把ColoredPoint类型当作接收器来调用Point里的方法，即使ColoredPoint里没有声明这些方法：


在类型中内嵌的匿名字段也可能是一个命名类型的指针，这种情况下字段和方法会被间接地引入到当前的类型中(译注：访问需要通过该指针指向的对象去取)。
添加这一层间接关系让我们可以共享通用的结构并动态地改变对象之间的关系。

下面这个ColoredPoint的声明内嵌了一个*Point的指针。

type ColoredPoint struct {
    *Point
    Color color.RGBA
}

p := ColoredPoint{&Point{1, 1}, red}
q := ColoredPoint{&Point{5, 4}, blue}
fmt.Println(p.Distance(*q.Point)) // "5"
q.Point = p.Point                 // p and q now share the same Point
p.ScaleBy(2)
fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
一个struct类型也可能会有多个匿名字段。我们将ColoredPoint定义为下面这样：

type ColoredPoint struct {
    Point
    color.RGBA
}


然后这种类型的值便会拥有Point和RGBA类型的所有方法，以及直接定义在ColoredPoint中的方法.
当编译器解析一个选择器到方法时，比如p.ScaleBy，
 (1)它会首先去找直接定义在这个类型里的ScaleBy方法，
 (2)然后找被ColoredPoint的内嵌字段们引入的方法，
 (3)然后去找Point和RGBA的内嵌字段引入的方法，
 (4)然后一直递归向下找。
如果选择器有二义性的话编译器会报错，比如你在同一级里有两个同名的方法。



方法只能在命名类型(像Point)或者指向类型的指针上定义，但是多亏了内嵌，有些时候我们给匿名struct类型来定义方法也有了手段。

下面是一个小trick。这个例子展示了简单的cache，其使用两个包级别的变量来实现，一个mutex互斥量(§9.2)和它所操作的cache：

var (
    mu sync.Mutex // guards mapping
    mapping = make(map[string]string)
)

func Lookup(key string) string {
    mu.Lock()
    v := mapping[key]
    mu.Unlock()
    return v
}
下面这个版本在功能上是一致的，但将两个包级别的变量放在了cache这个struct一组内：

var cache = struct {
    sync.Mutex
    mapping map[string]string
}{
    mapping: make(map[string]string),
}


func Lookup(key string) string {
    cache.Lock()
    v := cache.mapping[key]
    cache.Unlock()
    return v
}
*/

/*
当你根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了。你可以根据选择来调用接收器各不相同的方法。
下面的例子，变量op代表Point类型的addition或者subtraction方法，Path.TranslateBy方法会为其Path数组中的每一个Point来调用对应的方法：


*/
type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

/*
Go语言里的集合一般会用map[T]bool这种形式来表示，T代表元素类型。集合用map类型来表示虽然非常灵活，但我们可以以一种更好的形式来表示它.
例如在数据流分析领域，集合元素通常是一个非负整数，集合会包含很多元素，并且集合会经常进行并集、交集操作，这种情况下，bit数组会比map表现更加理想。
(译注：这里再补充一个例子，比如我们执行一个http下载任务，把文件按照16kb一块划分为很多块，需要有一个全局变量来标识哪些块下载完成了，这种时候也需要用到bit数组)

一个bit数组通常会用一个无符号数或者称之为“字”的slice来表示，每一个元素的每一位都表示集合里的一个值。当集合的第i位被设置时，我们才说这个集合包含元素i。
*/
// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

/*
因为每一个字都有64个二进制位，所以为了定位x的bit位，我们用了x/64的商作为字的下标，并且用x%64得到的值作为这个字内的bit的所在位置。
UnionWith这个方法里用到了bit位的“或”逻辑操作符号|来一次完成64个元素的或计算

*/

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

/*
这里留意一下String方法，是不是和3.5.4节中的intsToString方法很相似；bytes.Buffer在String方法里经常这么用。
当你为一个复杂的类型定义了一个String方法时，fmt包就会特殊对待这种类型的值，这样可以让这些类型在打印的时候看起来更加友好，而不是直接打印其原始的值。

这里要注意：我们声明的String和Has两个方法都是以指针类型*IntSet来作为接收器的，但实际上对于这两个类型来说，把接收器声明为指针类型也没什么必要。
不过另外两个函数就不是这样了，因为另外两个函数操作的是s.words对象，如果你不把接收器声明为指针对象，那么实际操作的是拷贝对象，而不是原来的那个对象。
*/

/*
封装

一个对象的变量或者方法如果对调用方是不可见的话，一般就被定义为“封装”。封装有时候也被叫做信息隐藏，同时也是面向对象编程最关键的一个方面。

Go语言只有一种控制可见性的手段：大写首字母的标识符会从定义它们的包中被导出，小写字母的则不会。
这种限制包内成员的方式同样适用于struct或者一个类型的方法。因而如果我们想要封装一个对象，我们必须将其定义为一个struct。



封装提供了三方面的优点。
1.因为调用方不能直接修改对象的变量值，其只需要关注少量的语句并且只要弄懂少量变量的可能的值即可。
2.隐藏实现的细节，可以防止调用方依赖那些可能变化的具体实现，这样使设计包的程序员在不破坏对外的api情况下能得到更大的自由。

把bytes.Buffer这个类型作为例子来考虑。这个类型在做短字符串叠加的时候很常用，所以在设计的时候可以做一些预先的优化，比如提前预留一部分空间，来避免反复的内存分配。又因为Buffer是一个struct类型，这些额外的空间可以用附加的字节数组来保存，且放在一个小写字母开头的字段中。这样在外部的调用方只能看到性能的提升，但并不会得到这个附加变量。Buffer和其增长算法我们列在这里，为了简洁性稍微做了一些精简：

3.是阻止了外部调用方对对象内部的值任意地进行修改。因为对象内部变量只可以被同一个包内的函数修改，所以包的作者可以让这些函数确保对象内部的一些值的不变性。

*/
type Buffer struct {
	buf     []byte
	initial [64]byte
	//
}

// Grow expands the buffer's capacity, if necessary,
// to guarantee space for another n bytes. [...]
func (b *Buffer) Grow(n int) {
	if b.buf == nil {
		b.buf = b.initial[:0] // use preallocated space initially
	}
	if len(b.buf)+n > cap(b.buf) {
		buf := make([]byte, b.Len(), 2*cap(b.buf)+n)
		copy(buf, b.buf)
		b.buf = buf
	}
}
