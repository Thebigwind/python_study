package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	//这个测试一个字符串是否符合一个表达式。
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) //----->true

	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要使用 Compile 一个优化的 Regexp 结构体。
	r, _ := regexp.Compile("p([a-z]+)ch")
	//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。
	fmt.Println(r.MatchString("peach")) //----->true

	//这是查找匹配字符串的。返回最左侧也就是第一次匹配的结果(单词分割符是任意非字母符号)
	fmt.Println(r.FindString("peach/punch paech")) //----->peach

	//这个也是查找第一次匹配的字符串的，但是返回的匹配开始和结束位置索引,如下是返回peach的p的索引0,和h的下一个索引5,左闭右开。
	fmt.Println(r.FindStringIndex("peach punch")) //----->[0 5]

	//FindStringSubmatch 返回第一次（也就是下面的peach）完全匹配和局部匹配的字符串。例如，这里会返回 p([a-z]+)ch 和 `([a-z]+) 的信息。
	fmt.Println(r.FindStringSubmatch("peach punch")) //----->[peach ea]

	//类似的，这个会返回完全匹配和局部匹配的索引位置。
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) //----->[0 5 1 3],int数组，左闭右开区间的拼接

	//带 All 的这个函数返回所有的匹配项，而不仅仅是首次匹配项。例如查找匹配表达式的所有项。
	//第二个参数：-1：返回索引匹配项；(n int)返回前n个匹配项
	fmt.Println(r.FindAllString("peach punch pinch", -1)) //----->[peach punch pinch]

	//这个函数提供一个正整数来限制匹配次数。
	fmt.Println(r.FindAllString("peach punch pinch", 2)) //----->[peach punch]

	//All 同样可以对应到上面的所有函数。
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) //----->[[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	//上面的例子中，我们使用了字符串作为参数，并使用了如 MatchString 这样的方法。我们也可以提供 []byte参数并将 String 从函数名中去掉。
	fmt.Println(r.Match([]byte("peach"))) //----->true

	//创建正则表达式常量时，可以使用 Compile 的变体MustCompile。因为 Compile 返回两个值，不能用语常量。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r) //----->p([a-z]+)ch

	//regexp 包也可以用来替换部分字符串为其他值。
	fmt.Println(r.ReplaceAllString("a peach", "<wangdy>")) //----->a <wangdy>

	//Func 变量允许传递匹配内容到一个给定的函数中，
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) //----->a PEACH

	/*#######################常见表达式###########################*/
	// 查找汉字
	testText := "Hello 世界, I am golang!"
	reg := regexp.MustCompile(`[\p{Han}]+`)
	fmt.Println(reg.FindAllString(testText, -1)) // ----->[世界]

	reg = regexp.MustCompile(`[\P{Han}]+`)
	fmt.Println(reg.FindAllString(testText, -1))        // ----->["Hello " ", I am golang!"]
	fmt.Printf("%q\n", reg.FindAllString(testText, -1)) // ----->["Hello " ", I am golang!"]

	//Email
	reg = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	fmt.Println(reg.MatchString("fjfha@qq.com"))

	//    用户名密码：
	reg = regexp.MustCompile(`[a-zA-Z]|\w{6,18}`)
	fmt.Println(reg.MatchString("w_dy_246"))
	//网上一堆……
	/*##################################################*/

	reg = regexp.MustCompile(`^[1-9][0-9]{0,}\w{1,2}$`)
	fmt.Println(reg.MatchString("10KB"))
	fmt.Println(reg.MatchString("10KB1h"))

	reg = regexp.MustCompile(`^[1-9][0-9]{0,}`)
	fmt.Println(reg.FindAllString("10KB", -1))
	fmt.Println(reg.FindString("10KB"))

	reg = regexp.MustCompile(`[a-zA-z]{1,2}`)

	//reg = regexp.MustCompile(`[k,m,g,t,K,M,G,T]{0,1}[b,B]{0,1}`)
	fmt.Println(reg.FindAllString("10KB", -1))
	fmt.Println(reg.FindAllString("10B", -1))
	fmt.Println(reg.FindAllString("10bk", -1))
	fmt.Println(reg.FindString("1KB"))

}
