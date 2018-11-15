/*
上面我们介绍了默认的非缓存类型的channel，不过Go也允许指定channel的缓冲大小，
很简单，就是channel可以存储多少元素。ch:= make(chan bool, 4)，创建了可以存储
4个元素的bool型channel。在这个channel 中，前4个元素可以无阻塞的写入。
当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。

ch := make(chan type, value)

当 value = 0 时，channel 是无缓冲阻塞读写的，
当value > 0 时，channel 有缓冲、是非阻塞的，
直到写满 value 个元素才阻塞写入。
*/

package main

import "fmt"

func main() {
	c := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("-------------------------")
	/*
		上面这个例子中，我们需要读取两次c，这样不是很方便，Go考虑到了这一点，
		所以也可以通过range，像操作slice或者map一样操作缓存类型的channel，
	*/
	c = make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		fmt.Println("x:", x)
	}
	fmt.Println("xx")
	close(c)
	fmt.Println("oo")

}

/*
for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。
上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。
关闭channel之后就无法再发送任何数据了，在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。
如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。

1.记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
2.另外记住一点的就是channel不像文件之类的，不需要经常去关闭，
只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的
*/
