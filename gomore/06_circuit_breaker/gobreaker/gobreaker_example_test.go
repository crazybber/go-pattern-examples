package gobreaker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

)

var cb *gobreaker.CircuitBreaker


func TestGoBreaker(t *testing.T) {
	body, err := Get("http://www.google.com/robots.txt")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body))
}


func initBreaker() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = gobreaker.NewCircuitBreaker(st)
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

