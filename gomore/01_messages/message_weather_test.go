package messaging

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

/*不同城市的象天气预报可以应用这个模式。*/
//这个是一个第三方的例子
func TestMessageSubPub(t *testing.T) {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	//订阅全部
	all := p.SubscribeTopic(nil)

	//订阅包含天气的消息
	onlyweathers := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "weather")
		}
		return false
	})

	p.Publish("weather bad, SH")
	p.Publish("weather fine,SZ")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range onlyweathers {
			fmt.Println("Received:", msg)
		}
	}()
	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}
