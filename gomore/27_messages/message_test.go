package messaging

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

////////////////////////////////
//通常意义上是，连接消息队列之后就可以发送消息
//当订阅著之后才会收到相关Topic消息的推送
////////////////////////////////

func TestMessageSubAndPubWithTopic(t *testing.T) {
	var wg sync.WaitGroup

	topicName := "seeking passengers"
	//假设评估
	topic := Topic{
		Name:          topicName,
		UserQueueSize: 5,
	}

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	//user 1

	//用户tom订阅拼车消息,订阅的是车主发布的拼车消息
	if subScriberTom, ok := topic.Subscribe(123, topicName); ok {

		go func() {
			defer wg.Done()
		EXIT:
			for {
				select {
				case <-ctx.Done():
					fmt.Println("tom receive cancel, exit")
					break EXIT
				default:
					msg := Message{}
					err := subScriberTom.Receive(&msg)
					if err == nil {
						fmt.Println("tom receive subscribed msg:", msg)
					}
				}
				time.Sleep(200)
			}
		}()
	}

	wg.Add(1)
	//订阅成功了
	//发送一个消息

	//用户Lily订阅拼车消息,订阅的是车主发布的拼车消息
	if subSCriptionLily, ok := topic.Subscribe(456, topicName); ok {
		go func() {
			defer wg.Done()
		EXIT:
			for {
				select {
				case <-ctx.Done():
					fmt.Println("lily receive cancel, exit")
					break EXIT
				default:
					msg := Message{}
					err := subSCriptionLily.Receive(&msg)
					if err == nil {
						fmt.Println("lily receive subscribed msg:", msg)
					}
				}
				time.Sleep(200)
			}
		}()
	}

	go func() {
		//模拟发送消息
		msg := Message{
			Text: "i am looking for 1 passenger",
			From: Session{User{123, "lily"}, time.Now()},
		}
		topic.Publish(msg)

		msg = Message{
			Text: "i am looking for 2 passenger",
			From: Session{User{123, "lucy"}, time.Now()},
		}

		topic.Publish(msg)

		msg = Message{
			Text: "i am looking for passenger as many as i can",
			From: Session{User{123, "rose"}, time.Now()},
		}

		topic.Publish(msg)
		time.Sleep(time.Second)
		cancel()

	}()

	wg.Wait()
	fmt.Println("all message done,exit it")

}
