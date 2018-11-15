/*
首先定义需要编码的结构，然后调用encoding/json标准库的Marshal方法生成json byte数组，转换成string类型即可。

golang和json的大部分数据结构匹配，对于复合结构，go可以借助结构体和空接口实现json的数组和对象结构。
通过struct tag可以灵活的修改json编码的字段名和输出控制

Marshal

func Marshal(v interface{}) ([]byte, error)
1
把对象转换为JSON:
　　• 布尔型转换为 JSON 后仍是布尔型　， 如true -> true

　　• 浮点型和整数型转换后为JSON里面的常规数字，如 1.23 -> 1.23

　　• 字符串将以UTF-8编码转化输出为Unicode字符集的字符串，特殊字符比如<将会被转义为\u003c

　　• 数组和切片被转换为JSON 里面的数组，[]byte类会被转换为base64编码后的字符串，slice的零值被转换为null

　　• 结构体会转化为JSON对象，并且只有结构体里边以大写字母开头的可被导出的字段才会被转化输出，而这些可导出的字段会作为JSON对象的字符串索引

　　• 转化一个map 类型的数据结构时，该数据的类型必须是 map[string]T（T 可以是encoding/json 包支持的任意数据类型）

*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Account struct {
	Email  string
	Passwd string
	Money  float64
}

type User struct {
	Name  string
	Age   int
	Role  []string
	Skill map[string]float64
}

type Label struct {
	group []string
}

func test() {
	
	label := Label{
		group: "" 
	}
}
func main() {
	account := Account{
		Email:  "1181124306@qq.com",
		Passwd: "xxoo",
		Money:  41314.3,
	}
	acc, err := json.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(acc)
	fmt.Println(string(acc))

	fmt.Println("----------------------------")
	skill := make(map[string]float64)
	skill["python"] = 69.0
	skill["golang"] = 87.3
	skill["sql"] = 88.4

	user := User{
		Name:  "luffy",
		Age:   28,
		Role:  []string{"Owner", "Master"},
		Skill: skill,
	}

	us, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(us)
	fmt.Println(string(us))
}

/*
func main1() {
	account := Account{
		Email:  "1181124306@qq.com",
		Passwd: "xxoo",
		Money:  41314.3,
	}
	acc, err := json.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(acc)
	fmt.Println(string(acc))
}
*/
