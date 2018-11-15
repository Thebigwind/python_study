package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("cpus:", runtime.NumCPU())    //查看有多少个逻辑cpu, logical cpus cpus: 4
	fmt.Println("goroot:", runtime.GOROOT())  //T returns the root of the Go tree goroot: D:\go
	fmt.Println("os/platform:", runtime.GOOS) //查看目标操作系统 os/platform: windows
}
