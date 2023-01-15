package customContext

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// func Demo() {
// 	ctx := context.Background()
// 	// ctx := context.TODO()

// 	ctx, cancel := context.WithTimeout(ctx, time.Second)
// 	// ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	// timer := time.AfterFunc(time.Millisecond*1000, func() {
// 	// 	fmt.Println("Canceling context")
// 	// 	cancel()
// 	// })

// 	go func() {
// 		scanner := bufio.NewScanner(os.Stdin)
// 		scanner.Scan()
// 		// scanner.Text()
// 		// ctx.Done() closes channel without error "deadline exceeded"
// 		fmt.Println("Background context: ", ctx.Err(), " Time exceeded: ", errors.Is(ctx.Err(), context.DeadlineExceeded), <-ctx.Done())
// 		cancel()
// 	}()

// 	time.Sleep(time.Millisecond * 3000)
// 	fmt.Println("Background context: ", ctx.Err(), " Time exceeded: ", errors.Is(ctx.Err(), context.DeadlineExceeded), <-ctx.Done())
// }

type contextUserIdKey int

const (
	userIdKey contextUserIdKey = iota
)

func handler(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("user-id")

	ctx := context.WithValue(r.Context(), userIdKey, id)
	// ctx, _ = context.WithTimeout(ctx, time.Second)

	log.Println("handler started", ctx.Value(userIdKey)) // ctx.Value(userIdKey) == ""
	defer log.Println("handler ended")

	result, err := processRequest(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, result)
}

func processRequest(ctx context.Context) (string, error) {
	id := ctx.Value(userIdKey)

	if id == "" {
		return "", fmt.Errorf("user id not found")
	}

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err, "here is error")
		return "", err
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	case <-time.After(1 * time.Second):
		return fmt.Sprintf("hello user #%s", id), nil
		// fmt.Fprintf(w, "hello\n")
	}

}

func Serve() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Run() {
	// Demo()
	Serve()
}
