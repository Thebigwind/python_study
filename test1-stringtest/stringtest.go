package main

import (
	"fmt"
	"strconv"
	"strings"
)

const World = "world"

func test1() {
	var hello = "hello"

	// 104 is the ascii code of char 'h'
	fmt.Println(hello[0]) // 104

	// strings are immutable
	// hello[0] = 'H' // error: cannot assign to hello[0]

	// all variables are addressable
	fmt.Println(&hello)

	// bytes in string are not addressable
	// fmt.Println(&hello[0]) // error: cannot take the address of hello[0]

	// string concatenation
	var helloWorld = hello + " " + World
	helloWorld += "!"
	fmt.Println(helloWorld) // hello world!

	// comparison
	var hello2 = helloWorld[:len(hello)]
	fmt.Println(hello == hello2)    // true
	fmt.Println(hello > helloWorld) // false

	// get string length
	fmt.Println(len(hello), len(helloWorld)) // 5 12

	// strings util functions
	fmt.Println(strings.HasPrefix(helloWorld, hello)) // true

	fmt.Println("--------------------")

	str := "保健\x61\142\u0042"
	fmt.Printf("%s\n", str)                    //保健abB
	fmt.Printf("%x, len: %d\n", str, len(str)) //e4bf9de581a5616242, len: 9

	fmt.Println("---------------------")
	s := "abcdef"
	for i, rn := range s {
		fmt.Printf("%d: 0x%xd %s \n", i, rn, string(rn))
	}
	/*
		0: 0x61d a
		1: 0x62d b
		2: 0x63d c
		3: 0x64d d
		4: 0x65d e
		5: 0x66d f
	*/
}
func main() {
	//test()
	TrimTest()
}

/*
按下标访问
不可变
字符串有地址，但是不能获取元素的地址，例如&str[1]
通过＋进行字符串连接
可对字符串进行比较
通过len获取字符串的长度
strings是go为我们提供的字符串包
*/

func test() {
	s := "@:/vol1/site-pacakges-test/tablib/packages/odf3/text.py"

	//res := strings.TrimLeft(s, "@:/vol1/site-pacakges-testb")
	//res1 := strings.Replace()

	fmt.Println(s[len(s)-5:])
	fmt.Println(s)

}

func TrimTest() {

	fmt.Println(strings.TrimRight("abba", "ba"))
	fmt.Println(strings.TrimRight("abcdaaaaa", "abcd"))
	fmt.Println(strings.TrimSuffix("abcddcba", "dcba"))
	fmt.Println(strings.TrimRight("ab/ba", "ba"))
	fmt.Println("--------------------------------")
	fmt.Println(strings.TrimLeft("/ab/air/test", "/ab"))
	fmt.Println(strings.TrimPrefix("/ab/air/test", "/ab"))

	//字符串s中是否包含substr，返回bool值 func Contains(s, substr string) bool
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	//Output:
	//true
	//false
	//true
	//true

	//字符串链接，把slice a通过sep链接起来 func Join(a []string, sep string) string
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	//Output:foo, bar, baz

	//在字符串s中查找sep所在的位置，返回位置值，找不到返回-1  func Index(s, sep string) int
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	//Output:4
	//-1

	//重复s字符串count次，最后返回重复的字符串  func Repeat(s string, count int) string
	fmt.Println("ba" + strings.Repeat("na", 2))
	//Output:banana

	//在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换  func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	//Output:oinky oinky oink
	//moo moo moo
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	//Output:oinky oinky oink
	//moo moo moo

	//把s字符串按照sep分割，返回slice func Split(s, sep string) []string
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	//Output:["a" "b" "c"]
	//["" "man " "plan " "canal panama"]
	//[" " "x" "y" "z" " "]
	//[""]

	//在s字符串的头部和尾部去除cutset指定的字符串 func Trim(s string, cutset string) string
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
	//Output:["Achtung"]

	//去除s字符串的空格符，并且按照空格分割返回slice func Fields(s string) []string
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//Output:Fields are: ["foo" "bar" "baz"]

	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	//Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	//Parse 系列函数把字符串转换为其他类型
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
