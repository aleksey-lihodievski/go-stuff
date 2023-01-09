package types

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

type CopyPerson struct {
	FirstName string
	LastName  string
	Email     string
}

type Person struct {
	FirstName string
	LastName  string
	Email     string
}

type Plane struct {
	Speed int
	High  int
}

func typeConvertions() {
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "John.Doe@gmail.com",
	}

	p2 := CopyPerson{
		FirstName: "Mary",
		LastName:  "Curie",
		Email:     "Mary.Curie@tut.by",
	}

	p1 = Person(p2)
	fmt.Println(p1, p2)

	plane := Plane{
		Speed: 5,
		High:  5,
	}
	fmt.Println(plane)

	// fmt.Println(Person(plane)) // does not allow
}

func constats() {
	const v1 int = 5
	const v2 = 5.0
	const large = 50000000000000000000000000000000000000000
	fmt.Println(v1*5, v1*v2*5, (large+large)/(large))
	fmt.Println(5/3, 5./3)
}

func typePrints() {
	var some int = 5
	fmt.Printf("var %d, type %T, size (bites) %d\n", some, some, unsafe.Sizeof(some))
}

func customErrors() {
	err := errors.New("Hello")

	var builder strings.Builder

	builder.WriteString(err.Error())
	builder.WriteString("Hello")

	fmt.Println(builder.String())
}

func pointers() {
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "John.Doe@gmail.com",
	}

	person2 := &Person{
		FirstName: "Mary",
		LastName:  "Curie",
		Email:     "Mary.Curie@tut.by",
	}

	fmt.Println(person1, person2)
}

func Run() {
	// typeConvertions()
	// constats()
	// typePrints()
	// customErrors()
	// pointers()
}
