Go设置cookie
Go语言中通过net/http包中的SetCookie来设置：

http.SetCookie(w ResponseWriter, cookie *Cookie)
w表示需要写入的response，cookie是一个struct，让我们来看一下cookie对象是怎么样的

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

// MaxAge=0 means no 'Max-Age' attribute specified.
// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
我们来看一个例子，如何设置cookie

expiration := time.Now()
expiration = expiration.AddDate(1, 0, 0)
cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
http.SetCookie(w, &cookie)
　　

Go读取cookie
上面的例子演示了如何设置cookie数据，我们这里来演示一下如何读取cookie

cookie, _ := r.Cookie("username")
fmt.Fprint(w, cookie)
还有另外一种读取方式

for _, cookie := range r.Cookies() {
	fmt.Fprint(w, cookie.Name)
}
可以看到通过request获取cookie非常方便。