package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock sync.Mutex

func main() {

	testMap()
}

/*
互斥锁

其中Mutex为互斥锁，Lock()加锁，Unlock()解锁，使用Lock()加锁后，便不能再次对其进行加锁，
直到利用Unlock()解锁对其解锁后，才能再次加锁．适用于读写不确定场景，即读写次数没有明显的区别，
并且只允许只有一个读或者写的场景，所以该锁叶叫做全局锁。
func (m *Mutex) Unlock()用于解锁m，如果在使用Unlock()前未加锁，就会引起一个运行错误．
已经锁定的Mutex并不与特定的goroutine相关联，这样可以利用一个goroutine对其加锁，再利用其他goroutine对其解锁。

*/
func testMap() {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}

	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

	time.Sleep(time.Second)
	fmt.Println(a)
	/*
		我们利用了time.Sleep(time.Second)这个进行的延迟，goroute执行完毕，就进行输出，结果是进行了map的修改

		map[3:10 2:10 1:10 18:10 8:10]
		map[2:10 1:10 18:10 8:87 3:10]
	*/
}
