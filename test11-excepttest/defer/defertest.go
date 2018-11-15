package main

import (
	"fmt"
	"os"
)

/*
Defer is used to ensure that a function call is performed later in a program’s execution,
usually for purposes of cleanup. defer is often used where e.g.
ensure and finally would be used in other languages.

defer就是用来添加函数结束时执行的语句。注意这里强调的是添加，
而不是指定，Go中的defer是动态的
*/
/*
**defer在函数结束之前执行
多个defer的执行顺序： Multiple defers are stacked last-in first-out
 so that the most recently deferred function is run first.**
*/

func main() {
	/*
		//defer在函数结束之前执行
		defer goodbye()
		defer goodnight()

		fmt.Println("Hello world.")
	*/

	/*
		//defer是在return之前执行的
		fmt.Println("Hello world.")
		return
		defer goodbye()
		defer goodnight()
	*/

	//defer用于关闭文件
	f := createFile("E:\\defer.txt")
	defer closeFile(f)
	writeFile(f)

}
func goodnight() {
	fmt.Println("GoodNight")
}

func goodbye() {
	fmt.Println("Goodbye")
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

/*
//defer用于锁
func Divide(i int) error {
	mu.Lock()
	defer mu.Unlock()

	if i == 0 {
		return errors.New("Can't divide by zero!")
	}

	val /= i
	return nil
}
*/

/*
	fmt.Println(f())//1
	fmt.Println(f1())//5
	fmt.Println(f2())//1

函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
defer表达式可能会在设置函数返回值之后，在返回到调用函数之前，修改返回值，使最终的函数返回值与你想象的不一致。
*/
func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
