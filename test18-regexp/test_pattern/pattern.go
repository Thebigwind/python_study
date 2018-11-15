package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	pattern := "^[/0-9a-zA-Z_.]{1,}$" //""
	str := "/s.df"

	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("invalid pattern.", err.Error())
		return
	}
	/*
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("Panicing %s\n", e)
			}
		}()
	*/
	//reg := regexp.MustCompile(pattern)
	result := reg.FindString(str)
	fmt.Println("result:", result)
	str = "sdfa sd"
	if strings.ContainsAny(str, "&; |") {
		fmt.Println("err")
	}
}
