/*
Unmarshal接受一个byte数组和空接口指针的参数。和sql中读取数据类似，先定义一个数据实例，然后传其指针地址。

与编码类似，golang会将json的数据结构和go的数据结构进行匹配。匹配的原则就是寻找tag的相同的字段，然后查找字段。
查询的时候是大小写不敏感的：

Unmarshal

func Unmarshal(data [] byte, v interface{}) error
1
把 JSON 转换回对象的方法:
这个函数会把传入的 data 作为一个JSON来进行解析，解析后的数据存储在参数 v 中。
这个参数 v 也是任意类型的参数（但一定是一个类型的指针），原因是我们在是以此
函数进行JSON 解析的时候，这个函数不知道这个传入参数的具体类型，所以它需要接收所有的类型。


Encoders and Decoders

NewDecoder returns a new decoder that reads from r.

func NewDecoder(r io.Reader) *Decoder

A Decoder reads and decodes JSON values from an input stream.

type Decoder struct {
        // contains filtered or unexported fields
}

An Encoder writes JSON values to an output stream.

type Encoder struct {
        // contains filtered or unexported fields
}
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var jsonString string = `{
    "username":"rsj217@gmail.com",
    "password":"123"
}`

func Decode(r io.Reader) (u *User, err error) {
	u = new(User)
	err = json.NewDecoder(r).Decode(u)
	if err != nil {
		return
	}
	return
}

func main() {
	user, err := Decode(strings.NewReader(jsonString))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v\n", user)

	fmt.Println("------------------------")

	var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
