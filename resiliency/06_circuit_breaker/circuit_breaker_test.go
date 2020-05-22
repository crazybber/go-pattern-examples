/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-11 10:55:28
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-22 16:37:21
 */

package circuit

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var breaker *RequestBreaker

func TestObjectBreaker(t *testing.T) {

	jobToDo := func() (interface{}, error) {
		resp, err := http.Get("https://bing.com/robots.txt")
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}

	whenCondition := func(counts counters) bool {
		//失败率，可以由用户自己定义
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	breaker = NewRequestBreaker(Name("HTTP GET"), BreakIf(whenCondition))

	body, err := breaker.Do(jobToDo)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body.([]byte)))
}

func TestFunctionalBreaker(t *testing.T) {

	//something need to do
	jobToDo := func(ctx context.Context) error {
		resp, err := http.Get("https://bing.com/robots.txt")
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
		return nil
	}

	//wrapper and control job with a breaker
	circuitWork := Breaker(jobToDo, 2 /* failureThreshold */)

	params := context.TODO()

	// do the job as usually
	res := circuitWork(params)

	log.Print("\nresult:", res)

}
