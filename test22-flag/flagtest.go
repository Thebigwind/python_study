package main

import (
	"errors"
	"flag"
	"fmt"
)

var name string

func init() {
	flag.CommandLine = flag.NewFlagSet("question 2-2", flag.ExitOnError)
	/*
		第一个参数：存储该命令参数的值的地址，如：&name;
		第二个参数：指定了参数名称，这里是name；
		第三个参数：指定在未追加命令时，餐的默认值
		第四个参数：命令参数的简短说明
	*/
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
