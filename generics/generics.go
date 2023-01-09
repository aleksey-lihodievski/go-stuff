package generics

import "fmt"

type Number interface {
	int | float32 | float64 | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64
}

func SumNumbers[T Number](one, two T) T {
	return one + two
}

func SumSlice[T Number](slice []T) (sum T) {
	for _, v := range slice {
		sum += v
	}
	return
}

func DemoGenerics() {
	someFloat := 3.14
	someInt := 3

	sliceOfInts := []int{1, 2, 3, 4}
	sliceOfFloats := []float64{1.1, 2.2, 3.3, 4.4}

	fmt.Println(SumNumbers(someFloat, someFloat))
	fmt.Println(SumNumbers(someInt, someInt))

	fmt.Println(SumSlice[int](sliceOfInts))
	fmt.Println(SumSlice(sliceOfFloats))
}

func Run() {
	// DemoGenerics()
}
