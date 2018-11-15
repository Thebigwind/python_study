package main

import (
	"fmt"
	"math"
)

/*
golang中的接口就是一系列的未实现的方法组成。

**interface可以被任意的对象实现
一个对象可以实现任意多个interface**
*/
type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human  //an anonymous field of type Human
	school string
	loan   float32
}
type Employee struct {
	Human
	company string
	money   float32
}

//A human method to say hi
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//A human can sing a song
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee's method overrides Human's one
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

// Interface Men is implemented by Human, Student and Employee
// because it contains methods implemented by them.
type Men interface {
	SayHi()
	Sing(lyrics string)
	//Fuck()
}

/*



 */
// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	//a variable of the interface type Men
	var i Men

	//i can store a Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()               ////Hi, I am Mike you can call me on 222-222-XXX
	i.Sing("November rain") //La la la la... November rain

	//i can store an Employee too
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()                 //Hi, I am Sam, I work at Things Ltd.. Call me on 444-222-XXX
	i.Sing("Born to be wild") //La la la la... Born to be wild

	//a slice of Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//These elements are of different types that satisfy the Men interface
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

	fmt.Println("---------------------------------------")

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of
	// these structs as arguments to `measure`.
	measure(r)
	measure(c)
}
