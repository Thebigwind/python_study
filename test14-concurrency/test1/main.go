package main

import (
	"fmt"
	"time"
)

func comFunc() {
	fmt.Println("This is a common function.")
}

func main() {
	go comFunc()
	time.Sleep(time.Second * 3)

}
