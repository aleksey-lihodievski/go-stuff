package mutexes

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Test() {
	n := 0
	// isEven := make(chan bool)

	// this goroutine executes first, so it will count n as 0 at the beginning
	go func() {
		ev := n%2 == 0

		time.Sleep(500 * time.Millisecond)

		fmt.Printf("n=%d and even=%t\n", n, ev)
		// isEven <- ev
	}()

	go func() {
		n++
	}()

	time.Sleep(1 * time.Second)
	// fmt.Printf("n=%d isEven=%t\n", n, <-isEven)
}

func TestMutexes() {
	n := 0
	m := sync.RWMutex{}
	// isEven := make(chan bool)

	// this goroutine executes first but in isolation, so it will has its own copy of n
	go func() {
		m.Lock()
		defer m.Unlock()

		ev := n%2 == 0

		time.Sleep(500 * time.Millisecond)

		fmt.Printf("n=%d isEven=%t\n", n, ev)
	}()

	go func() {
		m.Lock()
		defer m.Unlock()

		n++
	}()

	// fmt.Printf("n=%d isEven=%t\n", n, <-isEven)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup() {
	wg := sync.WaitGroup{}
	counter := uint64(0)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// atomic uses mutexes under the hood to prevent race conditions
			atomic.AddUint64(&counter, 1)
			fmt.Println("Counter: ", atomic.LoadUint64(&counter))
		}()
	}

	wg.Wait()
	fmt.Println(atomic.LoadUint64(&counter), " iterations completed")
}

func TryToBreakWaitGroup() {
	wg := sync.WaitGroup{}
	count := uint64(0)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println("Count: ", atomic.LoadUint64(&count))
			atomic.AddUint64(&count, 1)
		}()

		time.Sleep(500 * time.Millisecond)
	}

	wg.Wait()
	fmt.Println(count, " iterations completed")
}

func MakeDeadlock() {
	var m sync.Mutex
	fmt.Println("Locking")
	m.Lock()
	m.Unlock()
	m.Lock()
}

func Run() {
	// Test()
	// TestMutexes()
	// TestWaitGroup()
	// TryToBreakWaitGroup()
	// MakeDeadlock()
}
