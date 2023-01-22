package goroutines

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	fmt.Println("Walk", t.Value)
	// fmt.Println("Walk", t.Value, t.Left)
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	// fmt.Println("Before send", t.Value)
	ch <- t.Value
	// fmt.Println("After send", t.Value)

	if t.Right != nil {
		Walk(t.Right, ch)
	}

	// fmt.Println("After walk", t.Value)

	// select {
	// case <-ch:
	// default:
	// 	if t.Left != nil {
	// 		go Walk(t.Left, ch)
	// 	}
	// }

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go Walk(t1, chan1)
	go Walk(t2, chan2)

	// fmt.Println("Before range")
	for v1 := range chan1 {
		v2, open := <-chan2
		if !open {
			break
		}

		fmt.Printf("v1 = %d, v2 = %d\n", v1, v2)

		if v1 != v2 {
			fmt.Println("Different")
			return false
		}
	}

	fmt.Println("Same")
	return true
}

func RunTree() {
	t1 := tree.New(4)
	t2 := tree.New(4)

	Same(t1, t2)
}

func GetExpensiveRandInt(ch chan int) {
	delay := rand.Intn(5)
	time.Sleep(time.Duration(delay) * time.Second)
	ch <- rand.Intn(100)
}

func RunExpensive() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go GetExpensiveRandInt(ch1)
	go GetExpensiveRandInt(ch2)

	var res1 int
	var res2 int
	got := 0

	for got < 2 {
		select {
		case res1 = <-ch1:
			got++
			fmt.Println("Got res1: ", res1)
		case res2 = <-ch2:
			got++
			fmt.Println("Got res2: ", res2)
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("No result, waiting...")
		}
	}

	fmt.Println("Result of expensive calculus = ", res1+res2)
}

func Run() {
	// RunTree()
	// RunExpensive()
}
