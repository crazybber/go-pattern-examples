package gomore

import (
	"fmt"
	"testing"
	"time"
)

////////////////////////////////
//通常意义上是，连接消息队列之后就可以发送消息
////当订阅著之后才会收到相关Topic消息的推送
//这里为了简化，直接订阅成功后发送消息，省去连接消息队列发送消息的步骤
////////////////////////////////
func TestMessageSubAndPub(t *testing.T) {

	exit := make(chan bool, 1)
	//创建一个队列
	msgQueue := Queue{Topics: map[string]Topic{}}

	//创建一个感兴趣的话题
	topic := msgQueue.AddTopic("i want apple", 10)

	//向队列订阅话题
	if subSCription123, ok := topic.Subscribe(123, "tom want apple"); ok {

		//订阅成功了

		go func() {
			EXIT
			for {
				select {
				case <-exit:
					break EXIT
				default:
					msg := Message{}
					subSCription123.Receive(&msg)
					fmt.Println(msg)
				}
				time.Sleep(200)
			}
		}()

		msg := Message{
			//Type 类型[code :1,2,3,4]
			Type: 1,
			Text: "here is a apple",
			From: Session{User{123, "lily"}, time.Now()},
		}
		subSCription123.Publish(msg)
		msg.Type++
		subSCription123.Publish(msg)
	}

	if subSCription456, ok := topic.Subscribe(456, "lily want peach"); ok {

		//订阅成功了
		//发送一个消息
		go func() {
			EXIT
			for {
				select {
				case <-exit:
					break EXIT
				default:
					msg := Message{}
					subSCription456.Receive(&msg)
					fmt.Println(msg)
				}
				time.Sleep(200)
			}
		}()

		msg := Message{
			//Type 类型[code :1,2,3,4]
			Type: 1,
			Text: "here is a peach",
			From: Session{User{123, "bob"}, time.Now()},
		}

		subSCription456.Publish(msg)

		msg.Type++

		subSCription456.Publish(msg)

	}

}
