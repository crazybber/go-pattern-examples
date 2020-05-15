/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-11 10:55:28
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-11 21:35:39
 */

package circuit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var breaker *RequestBreaker

func TestBasicBreaker(t *testing.T) {

	readyToTrip := func(counts counters) bool {
		//失败率，可以由用户自己定义
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	breaker = NewRequestBreaker(Name("HTTP GET"), ReadyToTrip(readyToTrip))

	body, err := Get("https://bing.com/robots.txt")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body))
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	body, err := breaker.Do(func() (interface{}, error) {
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
