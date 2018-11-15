/*
在http包中，有个Get方法如下：

func (c *Client) Get(url string) (resp *Response, err error)

我们看到了返回值，有error。我们可以对这个error进行处理：

resp, err := c.Get("http://blog.csdn.net/wangshubo1989?viewmode=contents")
if err != nil {
    log.Println(err)
    return
}

Go 语言使用 error 类型来返回函数执行过程中遇到的错误，如果返回的 error 值为 nil，
则表示未遇到错误，否则 error 会返回一个字符串，用于说明遇到了什么错误。通俗的说，error就是一个接口而已，定义如下：
声明

type error interface {
    Error() string

}

New方法
将字符串 text 包装成一个 error 对象返回
New returns an error that formats as the given text.

func New(text string) error

例子：
看看io.go中的定义：

var ErrShortWrite    = errors.New("short write")
var ErrShortBuffer   = errors.New("short buffer")
var EOF              = errors.New("EOF")
var ErrUnexpectedEOF = errors.New("unexpected EOF")
var ErrNoProgress    = errors.New("multiple Read calls return no data or error")

*/
package main

import (
	"fmt"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
