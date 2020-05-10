package circuit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var breaker *RequestBreaker

func TestBasicBreaker(t *testing.T) {

	var st Options
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts counters) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	breaker = NewRequestBreaker(st)

	body, err := Get("https://bing.com/robots.txt")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body))
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	body, err := breaker.Execute(func() (interface{}, error) {
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
