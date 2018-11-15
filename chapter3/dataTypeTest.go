package main

import (
	"bytes"
	"fmt"
	"math"
	"unicode/utf8"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	//intTest()
	//typeTransTest()
	//fmtTest()
	//floatTest()
	//strTest()

	//utf8Test()
	str2slice()
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"

	var message string
	var a, b, c int
	a = 1

	message = "Hello World!"

	fmt.Println(message, a, b, c)
}

/*
整数运算。
int8、int16、int32和int64四种截然不同大小的有符号整数类型，分别对应8、16、32、64bit大小的有符号整数，
与此对应的是uint8、uint16、uint32和uint64四种无符号整数类型。

Unicode字符rune类型是和int32等价的类型，通常用于表示一个Unicode码点。这两个名称可以互换使用。
同样byte也是uint8类型的等价类型，byte类型一般用于强调数值是一个原始的数据而不是一个小的整数。
无符号的整数类型uintptr，没有指定具体的bit大小但是足以容纳指针。uintptr类型只有在底层编程时才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。

int和int32也是不同的类型，即使int的大小也是32bit，在需要将int当作int32类型的地方需要一个显式的类型转换操作，反之亦然。

算术运算、逻辑运算和比较运算的二元运算符，它们按照优先级递减的顺序排列：
*      /      %      <<       >>     &       &^
+      -      |      ^
==     !=     <      <=       >      >=
&&
||

Go语言还提供了以下的bit位操作运算符，前面4个操作运算符并不区分是有符号还是无符号数：
&      位运算 AND
|      位运算 OR
^      位运算 XOR
&^     位清空 (AND NOT)
<<     左移
>>     右移

位操作运算符^作为二元运算符时是按位异或（XOR），当用作一元运算符时表示按位取反；
位操作运算符&^用于按位置零（AND NOT）：如果对应y中bit位为1的话, 表达式z = x &^ y结果z的对应的bit位为0，否则z对应的bit位等于x相应的bit位的值。


在x<<n和x>>n移位运算中，决定了移位操作bit数部分必须是无符号数；被操作的x数可以是有符号或无符号数。算术上，一个x<<n左移运算等价于乘以$2^n$，一个x>>n右移运算等价于除以$2^n$。

左移运算用零填充右边空缺的bit位，无符号数的右移运算也是用0填充左边空缺的bit位，但是有符号数的右移运算会用符号位的值填充左边空缺的bit位。因为这个原因，最好用无符号运算，这样你可以将整数完全当作一个bit位模式处理。

无符号数往往只有在位运算或其它特殊的运算场景才会使用，就像bit集合、分析二进制文件格式或者是哈希和加密操作等。它们通常并不用于仅仅是表达非负数量的场合。

类型不匹配的问题可以有几种不同的方法修复，最常见方法是将它们都显式转型为一个常见类型：

var compote = int(apples) + int(oranges)
*/
func intTest() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}  如果对应y中bit位为1的话, 表达式z = x &^ y结果z的对应的bit位为0，否则z对应的bit位等于x相应的bit位的值

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}

func typeTransTest() {
	//对于将一个大尺寸的整数类型转为一个小尺寸的整数类型，或者是将一个浮点数转为整数，可能会改变数值或丢失精度：

	f := 3.141 // a float64
	i := int(f)
	fmt.Println(f, i) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"
}

/*
注意fmt的两个使用技巧。通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，
但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。

字符面值通过一对单引号直接包含对应字符。最简单的例子是ASCII中类似'a'写法的字符面值，但是我们也可以通过转义的数值来表示任意的Unicode码点对应的字符，马上将会看到这样的例子。

字符使用%c参数打印，或者是用%q参数打印带单引号的字符
*/

func fmtTest() {
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}

/*
Go语言提供了两种精度的浮点数，float32和float64
浮点数的范围极限值可以在math包找到。
常量math.MaxFloat32表示float32能表示的最大数值，大约是 3.4e38；
对应的math.MaxFloat64常量大约是1.8e308。它们分别能表示的最小值近似为1.4e-45和4.9e-324


一个float32类型的浮点数可以提供大约6个十进制数的精度，而float64则可以提供约15个十进制数的精度；
通常应该优先使用float64类型，因为float32类型的累计计算误差很容易扩散，并且float32能精确表示的正整数并不是很大
（译注：因为float32的有效bit位只有23个，其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差）

用Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，
但是对应表格的数据，使用%e（带指数）或%f的形式打印可能更合适。所有的这三个打印形式都可以指定打印的宽度和控制打印精度


*/

func floatTest() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

/*
 */
func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

/*
复数
Go语言提供了两种精度的复数类型：complex64和complex128，分别对应float32和float64两种浮点数精度。内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部：

var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y)                 // "(-5+10i)"
fmt.Println(real(x*y))           // "-5"
fmt.Println(imag(x*y))           // "10"
*/

/*
一个布尔类型的值只有两种：true和false。if和for语句的条件部分都是布尔类型的值，并且==和<等比较操作也会产生布尔型的值。
一元操作符!对应逻辑非操作，因此!true的值为false

布尔值可以和&&（AND）和||（OR）操作符结合，并且有短路行为：如果运算符左边值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值，因此下面的表达式总是安全的：

s != "" && s[0] == 'x'
其中s[0]操作如果应用于空字符串将会导致panic异常

因为&&的优先级比||高（助记：&&对应逻辑乘法，||对应逻辑加法，乘法比加法优先级要高），下面形式的布尔表达式是不需要加小括弧的：

if 'a' <= c && c <= 'z' ||
    'A' <= c && c <= 'Z' ||
    '0' <= c && c <= '9' {
    // ...ASCII letter or digit...
}
布尔值并不会隐式转换为数字值0或1，反之亦然。必须使用一个显式的if语句辅助转换：

i := 0
if b {
    i = 1
}
*/

// btoi returns 1 if b is true and 0 if false.
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// itob reports whether i is non-zero.
func itob(i int) bool { return i != 0 }

/*
Strings

内置的len函数可以返回一个字符串中的字节数目（不是rune字符数目），索引操作s[i]返回第i个字节的字节值，i必须满足0 ≤ i< len(s)条件约束。

s := "hello, world"
fmt.Println(len(s))     // "12"
fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
如果试图访问超出字符串索引范围的字节将会导致panic异常：

c := s[len(s)] // panic: index out of range

第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个字节。我们先简单说下字符的工作方式。
子字符串操作s[i:j]基于原始的s字符串的第i个字节开始到第j个字节（并不包含j本身）生成一个新字符串。生成的新字符串将包含j-i个字节。

fmt.Println(s[0:5]) // "hello"


字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变，当然我们也可以给一个字符串变量分配一个新字符串值。可以像下面这样将一个字符串追加到另一个字符串：

s := "left foot"
t := s
s += ", right foot"
这并不会导致原始的字符串值被改变，但是变量s将因为+=语句持有一个新的字符串值，但是t依然是包含原先的字符串值。

fmt.Println(s) // "left foot, right foot"
fmt.Println(t) // "left foot"
因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的：

s[0] = 'L' // compile error: cannot assign to s[0]

如果两个字符串共享相同的底层数据的话也是安全的，这使得复制任何长度的字符串代价是低廉的.
一个字符串s和对应的子字符串切片s[7:]的操作也可以安全地共享相同的内存，因此字符串切片操作代价也是低廉的

一个原生的字符串面值形式是`...`，使用反引号代替双引号。
在原生的字符串面值中，没有转义操作；全部的内容都是字面的意思，包含退格和换行，因此一个程序中的原生字符串面值可能跨越多行

*/

func strTest() {
	s := "abcd"
	//s[0] = "L"   cannot use "L" (type string) as type byte in assignment
	fmt.Println(s)
	s1 := `whtat's you name ?
	my name is "" :luffy.`
	fmt.Println(s1)
}

/*
utf-8
UTF8是一个将Unicode码点编码为字节序列的变长编码。UTF8编码由Go语言之父Ken Thompson和Rob Pike共同发明的，现在已经是Unicode的标准。
UTF8编码使用1到4个字节来表示每个Unicode码点，ASCII部分字符只使用1个字节，常用字符部分使用2或3个字节表示。
每个符号编码后第一个字节的高端bit位用于表示总共有多少编码个字节。
如果第一个字节的高端bit为0，则表示对应7bit的ASCII字符，ASCII字符每个字符依然是一个字节，和传统的ASCII编码兼容。
如果第一个字节的高端bit是110，则说明需要2个字节；后续的每个高端bit都以10开头。更大的Unicode码点也是采用类似的策略处理。

0xxxxxxx                             runes 0-127    (ASCII)
110xxxxx 10xxxxxx                    128-2047       (values <128 unused)
1110xxxx 10xxxxxx 10xxxxxx           2048-65535     (values <2048 unused)
11110xxx 10xxxxxx 10xxxxxx 10xxxxxx  65536-0x10ffff (other values unused)
*/

//得益于UTF8编码优良的设计，诸多字符串操作都不需要解码操作。我们可以不用解码直接测试一个字符串是否是另一个字符串的前缀：

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

//或者是后缀测试：

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

//或者是包含子串测试：

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func utf8Test() {
	s := "Hello, 世界"
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	//一个简单的循环来统计字符串中字符的数目
	n := 0
	for _, _ = range s {
		n++
	}
	fmt.Println(n)
	fmt.Println(utf8.RuneCountInString(s))
	/*
			0	'H'	72
		1	'e'	101
		2	'l'	108
		3	'l'	108
		4	'o'	111
		5	','	44
		6	' '	32
		7	'世'	19990
		10	'界'	30028
		9
		9
	*/
	//将一个整数转型为字符串意思是生成以只包含对应Unicode码点字符的UTF8字符串：

	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"
	//如果对应码点的字符是无效的，则用\uFFFD无效字符作为替换：
	fmt.Println(string(1234567)) // "�"
}

/*
标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。
strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。

bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。
因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。在这种情况下，使用bytes.Buffer类型将会更有效

strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换.

unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。
每个函数有一个单一的rune类型的参数，然后返回一个布尔值。
*/

func basename(s string) string {
	//discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	//preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//一个字符串是包含的只读字节数组，一旦创建，是不可变的。相比之下，一个字节slice的元素则可以自由地修改。
//字符串和字节slice之间可以相互转换：
//一个[]byte(s)转换是分配了一个新的字节数组用于保存字符串数据的拷贝，然后引用这个底层的字节数组

func str2slice() {
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(b)
	fmt.Println(s2)
	fmt.Println()
	fmt.Println(string(b[0]))
}

/*
为了避免转换中不必要的内存分配，bytes包和strings同时提供了许多实用函数。下面是strings包中的六个函数：

func Contains(s, substr string) bool
func Count(s, sep string) int
func Fields(s string) []string
func HasPrefix(s, prefix string) bool
func Index(s, sep string) int
func Join(a []string, sep string) string
bytes包中也对应的六个函数：

func Contains(b, subslice []byte) bool
func Count(s, sep []byte) int
func Fields(s []byte) [][]byte
func HasPrefix(s, prefix []byte) bool
func Index(s, sep []byte) int
func Join(s [][]byte, sep []byte) []byte
*/

/*

bytes包还提供了Buffer类型用于字节slice的缓存。
一个Buffer开始是空的，但是随着string、byte或[]byte等类型数据的写入可以动态增长，一个bytes.Buffer变量并不需要初始化，因为零值也是有效的：
当向bytes.Buffer添加任意字符的UTF8编码时，最好使用bytes.Buffer的WriteRune方法，但是WriteByte方法对于写入类似'['和']'等ASCII字符则会更加有效。
*/
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

/*
字符串和数字的转换
strconv.Itoa(x))// "123 123"
strconv.FormatInt(int64(x), 2) // "1111011"		2表示二进制，FormatInt和FormatUint函数可以用不同的进制来格式化数字：

将一个字符串解析为整数，可以使用strconv包的Atoi或ParseInt函数
x, err := strconv.Atoi("123")             // x is an int
y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits  ParseInt函数的第三个参数是用于指定整型数的大小；例如16表示int16，0则表示int
*/
