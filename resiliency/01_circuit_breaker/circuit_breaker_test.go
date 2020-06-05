/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-11 10:55:28
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-22 17:21:06
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

var onStateChangeEvent = func(name string, from, to State) {
	fmt.Println("name:", name, "from:", from, "to", to)
}

var canOpenSwitch = func(current State, cnter counters) bool {

	if current == StateHalfOpen {
		return cnter.ConsecutiveFailures > 2
	}

	//失败率，可以由用户自己定义
	failureRatio := float64(cnter.TotalFailures) / float64(cnter.Requests)
	return cnter.Requests >= 3 && failureRatio >= 0.6
}

var canCloseSwitch = func(current State, cnter counters) bool {
	//失败率，可以由用户自己定义
	if cnter.ConsecutiveSuccesses > 2 {
		return true
	}
	//
	successRatio := float64(cnter.TotalFailures) / float64(cnter.Requests)
	return cnter.Requests >= 3 && successRatio >= 0.6
}

func TestObjectBreaker(t *testing.T) {

	jobToDo := func(ctx context.Context) (interface{}, error) {
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

	breaker = NewRequestBreaker(ActionName("HTTP GET"),
		WithBreakCondition(canOpenSwitch),
		WithCloseCondition(canCloseSwitch),
		WithShoulderHalfToOpen(2),
		WithStateChanged(onStateChangeEvent))

	body, err := breaker.Do(jobToDo)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body.([]byte)))

	log.Print("\nresult:", body.([]byte))
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
