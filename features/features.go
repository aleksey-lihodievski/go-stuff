package features

import (
	"fmt"
)

func closure(num int) func() {
	res := num
	return func() {
		res += 1
		fmt.Println("Closure", res)
	}
}

func switches() {
	switch 'c' {
	case 'a', 'b', 'c':
		{
			fmt.Println("a, b, c")
		}

	case 'd', 'e', 'f':
		{
			fmt.Println("d, e, f")
		}

	default:
		{
			fmt.Println("default")
		}
	}

	c := 'c'

	switch {
	case c < 'd':
		{
			fmt.Println("c less than 'd'")
		}

	case c > 'a':
		{
			fmt.Println("c greater than 'a'")
		}
	}

	var t interface{}

	t = true

	switch h := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", h) // %T prints whatever type h has
	case bool:
		fmt.Printf("boolean %t is of type %T\n", h, h) // h has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *h) // h has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *h) // h has type *int

	}

}

func bitShifts() {
	var a uint8 = 1
	var b uint8 = 2

	fmt.Println(a<<3, b>>1)
}

func bitShifts2() {
	fmt.Println(1 << 10) // 1024
}

func iotaFunc() {
	const some = iota
	fmt.Println(some)
}

func Run() {
	// switches()
	// bitShifts()
	// bitShifts2()
	// iotaFunc()
}
