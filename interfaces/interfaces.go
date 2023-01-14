package interfaces

import "fmt"

type Hello interface {
	SayHello()
}

type Goodbye interface {
	SayGoodbye()
}

type HelloGoodbye interface {
	Hello
	Goodbye
}

type HelloGoodbye2 interface {
	HelloStruct | GoodbyeStruct
}

func PrintHello(hello Hello) {
	fmt.Println(hello)
}

func Demo() {
	var h1 HelloGoodbye
	var h2 HelloGoodbye
	// is assignable
	h1 = &HelloStruct{Name: "Alex", Surname: "Lihodievski"}
	h2 = GoodbyeStruct{Name: "Alex"}

	h1.SayHello()
	h1.SayGoodbye()
	h2.SayHello()
	h2.SayGoodbye()
	PrintHello(h1)

	someSlice := []HelloGoodbye{h1, h2}
	fmt.Println(someSlice)

	// can not be assigned because HelloStruct2 recieves pointer
	// h = hp
}

func Run() {
	// Demo()
}

type HelloStruct struct {
	Name    string
	Surname string
}

type GoodbyeStruct struct {
	Name string
}

func (h *HelloStruct) SayHello() {
	fmt.Println("Hello", h.Name)
}

func (h *HelloStruct) SayGoodbye() {
	fmt.Println("Goodbye", h.Name)
}

func (h GoodbyeStruct) SayHello() {
	fmt.Println("Hello", h.Name)
}

func (h GoodbyeStruct) SayGoodbye() {
	fmt.Println("Goodbye", h.Name)
}
