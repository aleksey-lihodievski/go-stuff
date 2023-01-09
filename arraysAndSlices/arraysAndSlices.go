package arraysAndSlices

import (
	"fmt"
	"sort"
)

func arrays() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr3 := [10]int{1: 1, 3: 3, 4: 4, 9: 9}

	arr4 := arr2

	arr4[2] = 6

	fmt.Println(len(arr1), len(arr2), len(arr3), len(arr4))
	fmt.Println(cap(arr1), cap(arr2), cap(arr3), cap(arr4))
	fmt.Println(arr1, arr2, arr3, arr4)

}

func slices() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1
	slice3 := append([]int{}, slice1...)
	// to copy slice and not refer to the same heap memory

	// slice1 also change, because both slice1 and slice2 refer to the same heap memory
	slice2[2] = 6
	fmt.Println(slice1, slice2, slice3)

	fmt.Println(cap(slice1), cap(slice2), cap(slice3))
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println(cap(slice1), cap(slice2), cap(slice3))

	changeSlice(slice1)
	fmt.Println(slice1, slice2, slice3)

	slice4 := [][]int{{1, 2, 3}, {4, 5, 6}}
	slice5 := append([][]int{}, slice4...)

	fmt.Println(slice4, slice5)
	slice4[0][0] = 10
	fmt.Println(slice4, slice5)

	arr := []int{3, 2, 1, 4, 5}
	sort.Ints(arr)
	fmt.Println(arr)
}

func changeSlice(slice []int) {
	for i := range slice {
		slice[i] = i * 5
	}
}

func Run() {
	// arrays()
	// slices()

	// p := make([]int, 5, 10)
	// p := new([]int)
	// new() // returns pointer
	// make() // returns value
}
