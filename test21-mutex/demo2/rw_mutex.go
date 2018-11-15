/*
读写锁

读写锁实际是一种特殊的自旋锁，它把对共享资源的访问者划分成读者和写者，读者只对共享资源进行读访问，
写者则需要对共享资源进行写操作。这种锁相对于自旋锁而言，能提高并发性，因为在多处理器系统中，
它允许同时有多个读者来访问共享资源，最大可能的读者数为实际的逻辑CPU数。写者是排他性的，
一个读写锁同时只能有一个写者或多个读者（与CPU数相关），但不能同时既有读者又有写者。


*/
package main

import (
	// "fmt"
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)

	// 多个同时读
	go read(1)
	go read(2)

	time.Sleep(2 * time.Second)
}

func read(i int) {
	println(i, "read start")

	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "read over")
}
