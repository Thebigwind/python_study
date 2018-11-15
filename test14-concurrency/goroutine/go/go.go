package main

import (
	"fmt"
	"time"
)

func main() {
	test4()
}
func test1() {
	name := "luffy"
	go func() {
		fmt.Printf("Hello,%s\n", name)
	}()
	name = "lxf"
	time.Sleep(time.Millisecond)
}

func test2() {
	names := []string{"Eric", "lxf", "luffy", "luxuefeng", "lu"}
	for _, name := range names {
		go func() {
			fmt.Printf("Hello,%s\n", name)
		}()
	}

	time.Sleep(time.Millisecond)
}

func test3() {
	names := []string{"Eric", "lxf", "luffy", "luxuefeng", "lu"}
	for _, name := range names {
		go func() {
			fmt.Printf("Hello,%s\n", name)
		}()
		time.Sleep(time.Millisecond)
	}
}

//可重载入函数，把迭代变量name作为参数传递给go函数
func test4() {
	names := []string{"Eric", "lxf", "luffy", "luxuefeng", "lu"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello,%s\n", name)
		}(name)
		time.Sleep(time.Millisecond)
	}
}
