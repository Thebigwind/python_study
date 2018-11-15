package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	//test1()
	test2()
}

//12345.6  ba9ebd859cc93d501d0b6cfb6c5728e7
//123456   e10adc3949ba59abbe56e057f20f883e
//12345 6  416c6a219ad9b56c0f8f7e36e99fad0f
func test1() {
	h := md5.New()
	b := []byte("adfjaldf;adsjfalsdjfaldfjalejffadfadfgsdfsfgsfgsfgr53564747#44")
	h.Write(b) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	/*
		Write函数会把MD5对象内部的字符串clear掉，然后把其参数作为新的内部字符串。
		而Sum函数则是先计算出内部字符串的MD5值，而后把输入参数附加到内部字符串后面。
		fmt.Println(md5.Sum(b)) //md5.Sum(b) = h.Write(b) + h.Sum(nil)。
	*/

	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
}

func test2() {
	hash := md5.New()
	b := []byte("test")
	hash.Write(b)
	fmt.Printf("%x %x\n", hash.Sum(nil), md5.Sum(b))
	hash.Write(nil)
	fmt.Printf("%x %x\n", hash.Sum(b), hash.Sum(nil))

}

func test3() {
	TestString := "Hi, pandaman!"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}
