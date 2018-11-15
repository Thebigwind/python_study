/*
goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。
goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，
Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，
当然会根据相应的数据伸缩

*/
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}

func main() {
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行

	/*
		hello
		world
		hello
		hello
		hello
		world
		world
		hello
	*/
}

/*
go关键字很方便的就实现了并发编程。 上面的多个goroutine运行在同一个进程里面，共享内存数据，
不过设计上我们要遵循：不要通过共享来通信，而要通过通信来共享。

runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
默认情况下，在Go 1.5将标识并发系统线程个数的runtime.GOMAXPROCS的初始值由1改为了运行环境的CPU核数。

GOMAXPROCS 设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置。
*/
