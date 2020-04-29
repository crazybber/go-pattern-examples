package contexts

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	go func() {
		cContext()
	}()

	time.Sleep(time.Second)

	reqest, err := http.NewRequest("GET", "http://localhost:8099/hello", nil) // http client get 请求
	assertEq(nil, err)
	client := &http.Client{}
	ctx, cancel := context.WithCancel(context.Background())
	reqest = reqest.WithContext(ctx)

	go func() {
		select {
		case <-time.After(2 * time.Second):
			cancel()
		}
	}()

	response, err := client.Do(reqest)

}

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func cContext() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8099", nil)
}
