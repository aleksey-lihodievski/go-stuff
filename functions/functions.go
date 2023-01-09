package functions

import "fmt"

func Defer() {
	num := 1
	defer fmt.Println("Defer, beginning of the function, num: ", num)

	num++
	fmt.Println("End of function num: ", num)
}

func recoveringFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered %v\n", r)
		}
	}()

	value := panicFunc()
	fmt.Println("In the end", value)
}

func panicFunc() int {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Panic func", i)

		if i == 3 {
			panic("Panic")
		}
	}

	return 1
}

func Run() {
	// Defer()
	// recoveringFunc()
}
