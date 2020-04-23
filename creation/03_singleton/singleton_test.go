package singleton

import (
	"sync"
	"testing"
)

const workerCount = 500

func TestWorkerSingleton(t *testing.T) {
	ins1 := GetWorkerInstance()
	ins2 := GetWorkerInstance()
	if ins1 != ins2 {
		t.Fatal("worker(instance) is not exactly the same")
	}
}

// 获取500次，Worker 是否总是同一个worker
func TestParallelWorkerSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(workerCount)
	instances := [workerCount]*Worker{}
	for i := 0; i < workerCount; i++ {
		go func(index int) {
			instances[index] = GetWorkerInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < workerCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("Worker instance is not equal")
		}
	}
}

func TestManagerSingleton(t *testing.T) {
	ins1 := GetManagerInstance()
	ins2 := GetManagerInstance()
	if ins1 != ins2 {
		t.Fatal("Manager(instance) is not exactly the same")
	}
}

// 获取500次，Manager 是否总是同一个Manager
func TestParallelManagerSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(workerCount)
	instances := [workerCount]*Manager{}
	for i := 0; i < workerCount; i++ {
		go func(index int) {
			instances[index] = GetManagerInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < workerCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("Manager instance is not exactly equal")
		}
	}
}
