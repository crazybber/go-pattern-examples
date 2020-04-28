package gomore

import "testing"

func TestMessageSubAndPub(t *testing.T) {

	//创建一个队列
	msgQueue := MesssageQueue{Topics: map[uint64]*Topic{}}

	//创建一个话题
	topic := Topic{}

	//像队列订阅话题
	topic.Subscribe(123)
}
