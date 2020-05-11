package gobreaker

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

)

var cb *CircuitBreaker


func TestGoBreaker(t *testing.T) {

	initBreaker()

	body, err := Get("https://bing.com/robots.txt")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body))
}


func initBreaker() {
	var st Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil

}
