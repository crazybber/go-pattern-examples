package fanout

import (
	"context"
	"sync"
	"sync/atomic"

	"go.uber.org/zap" //https://github.com/uber-go/zap
	//Blazing fast, structured, leveled logging in Go.
)

var (
	log, _ = zap.NewDevelopment()
)

//Settings of pipeline
const (
	MaxWorkers      = 16
	MaxQueueSize    = 512
	MasterQueueSize = MaxQueueSize * MaxWorkers
)

//IDispatcher Message
type IDispatcher interface {
	Before(context.Context) error
	After() error
	Process(interface{}) error
}

//worker each work will dispatch message to several channels
type worker struct {
	index      uint32
	mutex      *sync.Mutex
	running    bool
	chain      chan interface{}
	debug      bool
	idle       uint32
	dispatcher IDispatcher //hold a dispacher，需要自己实现一个dispatcher 工厂
}

//Pipeline of workers
type Pipeline struct {
	workers map[int]*worker
	chain   chan interface{}
}

//DispatcherBuilder create Dispatcher
type DispatcherBuilder func() IDispatcher

//Start run
func (p *Pipeline) Start(ctx context.Context) {
	go func(pipe *Pipeline) {
		for {
			expectationWorkers := len(pipe.chain) % MaxWorkers
			if expectationWorkers >= MaxWorkers {
				expectationWorkers = 0
			}
			select {
			case <-ctx.Done():
				return
			case val, ok := <-pipe.chain:
				if !ok {
					return
				}
				go pipe.workers[expectationWorkers].stream(val)
			}
		}
	}(p)
}

//Dispatch message to chains
func (p *Pipeline) Dispatch(msg interface{}) {
	p.chain <- msg
}

//NewPipeline create a Workflow  with a dispacher builder and some workers
func NewPipeline(d DispatcherBuilder, idle uint32, debug bool) *Pipeline {
	ch := make(chan interface{}, MasterQueueSize)
	wk := make(map[int]*worker)
	for i := 0; i < MaxWorkers; i++ {
		wk[i] = &worker{
			index:      uint32(i + 1),
			chain:      make(chan interface{}, MaxQueueSize),
			mutex:      new(sync.Mutex),
			debug:      debug,
			idle:       idle,
			dispatcher: d(), //build real dispatcher
		}
	}
	return &Pipeline{workers: wk, chain: ch}
}

func (c *worker) stream(val interface{}) {
	c.chain <- val
	if !c.running {
		c.mutex.Lock()
		c.running = true
		ctx, cancel := context.WithCancel(context.Background())
		defer func(w *worker, cancel context.CancelFunc) {
			if w.debug {
				log.Info("Worker leaving", zap.Any("index", w.index), zap.Any("idle", w.idle))
			}

			if c.dispatcher != nil {
				err := c.dispatcher.After()
				if err != nil {
					log.Error("can not finish track issue", zap.Error(err))
				}
			}

			cancel()
			w.mutex.Unlock()
			w.running = false
		}(c, cancel)

		if c.dispatcher != nil {
			err := c.dispatcher.Before(ctx)
			if err != nil {
				log.Error("can not start worker", zap.Error(err))
			}
		}

		var idle uint32 = 0
		for {
			select {
			case msg := <-c.chain:
				atomic.StoreUint32(&idle, 0)
				if msg != nil && c.dispatcher != nil {
					err := c.dispatcher.Process(msg)
					if err != nil {
						log.Error("can not process message", zap.Any("msg", &msg), zap.Error(err))
					}
				}
			default:
				atomic.AddUint32(&idle, 1)
				if i := atomic.LoadUint32(&idle); i > 0 {
					if i > c.idle {
						return
					}
				}
			}
		}
	}
}
