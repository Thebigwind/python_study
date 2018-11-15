package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

/*
今天介绍一下io/ioutil 包。
Package ioutil implements some I/O utility functions.

就是跟操作文件、文件夹相关的函数，下面通过例子进行介绍。


*/

func main() {
	/*
		ReadAll
		读取 r 中的所有数据，返回读取的数据和遇到的错误。
		如果读取成功，则 err 返回 nil，而不是 EOF，因为 ReadAll 定义为读取
		所有数据，所以不会把 EOF 当做错误处理。

		func ReadAll(r io.Reader) ([]byte, error)
	*/
	link := "http://www.baidu.com"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get(link)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	//block forever at the next line
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

	fmt.Println("--------------")
	/*
		ReadDir
		ReadDir 读取指定目录中的所有目录和文件（不包括子目录）。
		返回读取到的文件信息列表和遇到的错误，列表是经过排序的。

		func ReadDir(dirname string) ([]os.FileInfo, error)
	*/

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	fmt.Println("------------------------------")
	/*
		ReadFile
		ReadFile 读取文件中的所有数据，返回读取的数据和遇到的错误。
		如果读取成功，则 err 返回 nil，而不是 EOF

		func ReadFile(filename string) ([]byte, error)
	*/

	// Perhaps the most basic file reading task is
	// slurping a file's entire contents into memory.
	dat, err := ioutil.ReadFile("test.txt")
	check(err)
	fmt.Print(string(dat))

	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.
	f, err := os.Open("test.txt")
	check(err)

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many
	// actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// You can also `Seek` to a known location in the file
	// and `Read` from there.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// The `io` package provides some functions that may
	// be helpful for file reading. For example, reads
	// like the ones above can be more robustly
	// implemented with `ReadAtLeast`.
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no built-in rewind, but `Seek(0, 0)`
	// accomplishes this.
	_, err = f.Seek(0, 0)
	check(err)

	// The `bufio` package implements a buffered
	// reader that may be useful both for its efficiency
	// with many small reads and because of the additional
	// reading methods it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close the file when you're done (usually this would
	// be scheduled immediately after `Open`ing with
	// `defer`).
	f.Close()

	fmt.Println("---------------------")
	/*
		TempDir
		操作系统中一般都会提供临时目录，比如linux下的/tmp目录（通过os.TempDir()可以获取到)。
		有时候，我们自己需要创建临时目录，比如Go工具链源码中（src/cmd/go/build.go），
		通过TempDir创建一个临时目录，用于存放编译过程的临时文件

		func TempDir(dir, prefix string) (name string, err error)
	*/

	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}

	fmt.Println("---------------------")

	/*
		TempFile
		TempFile 在 dir 目录中创建一个以 prefix 为前缀的临时文件，并将其以读
		写模式打开。返回创建的文件对象和遇到的错误。
		如果 dir 为空，则在默认的临时目录中创建文件（参见 os.TempDir），多次
		调用会创建不同的临时文件，调用者可以通过 f.Name() 获取文件的完整路径。
		调用本函数所创建的临时文件，应该由调用者自己删除。

		func TempFile(dir, prefix string) (f *os.File, err error)
	*/
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("---------------------")

	/*
		WriteFile
		WriteFile 向文件中写入数据，写入前会清空文件。
		如果文件不存在，则会以指定的权限创建该文件。
		返回遇到的错误。

		func WriteFile(filename string, data []byte, perm os.FileMode) error
	*/
	// To start, here's how to dump a string (or just
	// bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.
	f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// Use `Flush` to ensure all buffered operations have
	// been applied to the underlying writer.
	w.Flush()
}

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}
