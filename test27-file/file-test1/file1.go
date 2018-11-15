/*
文件操作
建立与打开文件
新建文件可以通过如下两个方法

func Create(name string) (file *File, err Error)
根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的。

func NewFile(fd uintptr, name string) *File
根据文件描述符创建相应的文件，返回一个文件对象

通过如下两个方法来打开文件：

func Open(name string) (file *File, err Error)
该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。

func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限


写文件
写文件函数：

func (file *File) Write(b []byte) (n int, err Error)
写入byte类型的信息到文件

func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
在指定位置开始写入byte类型的信息

func (file *File) WriteString(s string) (ret int, err Error)
写入string信息到文件
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "astaxie.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

	fmt.Println("------------------")

	userFile1 := "asatxie.txt"
	fl, err := os.Open(userFile1)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

/*
删除文件
Go语言里面删除文件和删除文件夹是同一个函数

func Remove(name string) Error

调用该函数就可以删除文件名为name的文件
*/
