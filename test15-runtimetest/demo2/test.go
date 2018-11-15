package main

import (
	"fmt"
	"runtime"
)

/*
GOMAXPROCS
这个需要着重介绍一下了。

func GOMAXPROCS(n int) int

GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and
returns the previous setting. If n < 1, it does not change the current setting. The
number of logical CPUs on the local machine can be queried with NumCPU. This call
will go away when the scheduler improves.

如果要在 goroutine 中使用多核，可以使用 runtime.GOMAXPROCS 函数修改，当参数小于 1 时使用默认值。

Gosched

func Gosched()

Gosched yields the processor, allowing other goroutines to run. It does not suspend
the current goroutine, so execution resumes automatically.
这个函数的作用是让当前 goroutine 让出 CPU，当一个 goroutine 发生阻塞，Go 会自动地把与该 goroutine
 处于同一系统线程的其他 goroutine 转移到另一个系统线程上去，以使这些 goroutine 不阻塞
*/

func main() {
	/*
		//单核
		runtime.GOMAXPROCS(1)
		exit := make(chan int)
		go func() {
			defer close(exit)
			go func() {
				fmt.Println("b")
			}()
		}()

		for i := 0; i < 10; i++ {
			fmt.Println("a:", i)

			if i == 4 {
				runtime.Gosched() //切换任务
			}
		}
		<-exit
	*/

	//多核  每次运行结果不一样
	runtime.GOMAXPROCS(2)
	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)

		if i == 4 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit

}

/*
a: 0
a: 1
a: 2
a: 3
a: 4
b
a: 5
a: 6
a: 7
a: 8
a: 9
*/
