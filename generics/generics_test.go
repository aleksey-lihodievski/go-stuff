package generics

import (
	"testing"
)

func TestSumSlice(t *testing.T) {
	want := 10

	// someFloat := 3.14
	// someInt := 3

	sliceOfInts := []int{1, 2, 3, 4}
	// sliceOfFloats := []float64{1.1, 2.2, 3.3, 4.4}

	if got := SumSlice(sliceOfInts); got != want {
		t.Errorf("SumSlice(sliceOfInts) = %v, want %v", got, want)
	}

}
