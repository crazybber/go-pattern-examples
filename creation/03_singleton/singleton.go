package singleton

import "sync"

////////////////////////////////
//way 1
//使用 sync的	once.Do(){}确保执行一次
////////////////////////////////

//Worker Singleton 是单例模式类
type Worker struct{}

//better to be pointer
var onlyTheWorker *Worker

// init a control
var once sync.Once

//GetWorkerInstance 总是获取到同一个Worker对象(内存位置相同)
func GetWorkerInstance() *Worker {

	//be sure ,to do this,only once!
	once.Do(func() {
		onlyTheWorker = &Worker{}
	})

	return onlyTheWorker
}

//Manager Singleton 是单例模式类
type Manager struct{}

//better to be pointer
var instance *Manager

//better to be pointer
var onlyTheManager *Manager

////////////////////////////////
//way2
//使用func init(){}函数来初始化保证，只初始化一次,更简单.
////////////////////////////////

func init() {
	onlyTheManager = &Manager{}
}

//GetManagerInstance 总是获取到同一个Manager对象(内存位置相同)
func GetManagerInstance() *Manager {
	return onlyTheManager
}
