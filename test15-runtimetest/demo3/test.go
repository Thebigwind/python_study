package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*

GOMAXPROCS和sync配合使用
sync这个包我们还没有介绍，这里就说明一点：
sync.WaitGroup只有3个方法，Add()，Done()，Wait()。其中Done()是Add(-1)的别名。
简单的来说，使用Add()添加计数，Done()减掉一个计数，计数不为0, 阻塞Wait()的运行
*/
func main() {
	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	go func() {
		defer wg.Done()

		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting To Finish")

	wg.Wait()

	fmt.Println("\nTerminating Program")
}
