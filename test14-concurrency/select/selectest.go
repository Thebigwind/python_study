/*
如果存在多个channel的时候，Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。

select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，
当多个channel都准备好的时候，select是随机的选择一个执行的。
*/

package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //
		}
		quit <- 0
	}()
	fmt.Println("xx")
	fibonacci(c, quit)
	fmt.Println("oo")

}

/*
在select里面还有default语法，select其实就是类似switch的功能，
default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）。

select {
case i := <-c:
	// use i
default:
	// 当c阻塞的时候执行这里
}
*/

//有时候会出现goroutine阻塞的情况，如何避免整个程序进入阻塞的情况？可以利用select来设置超时，通过如下的方式实现：
func main_1() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}

/*
runtime goroutine
runtime包中有几个处理goroutine的函数：

Goexit
退出当前执行的goroutine，但是defer函数还会继续调用

Gosched
让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

NumCPU
返回 CPU 核数量

NumGoroutine
返回正在执行和排队的任务总数

GOMAXPROCS
用来设置可以并行计算的CPU核数的最大值，并返回之前的值。
*/
