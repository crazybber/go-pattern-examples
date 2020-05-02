package messaging

import (
	"context"
	"errors"
	"fmt"
	"time"
)

//Message for msg in Message bus
type Message struct {
	Seq  int
	Text string
	From Session //消息来源
}

//User for user
type User struct {
	ID   uint64
	Name string
}

//Session inherit user
type Session struct {
	User
	Timestamp time.Time
}

//Subscription for user
//Subscription is a session
type Subscription struct {
	Session
	topicName string
	ctx       context.Context
	cancel    context.CancelFunc
	Ch        chan<- Message //发送队列
	Inbox     chan Message   //接收消息的队列
}

func newSubscription(uid uint64, topicName string, UserQueueSize int) Subscription {

	ctx, cancel := context.WithCancel(context.Background())

	return Subscription{
		Session: Session{User{ID: uid}, time.Now()},
		ctx:     ctx,
		cancel:  cancel,
		Ch:      make(chan<- Message, UserQueueSize), //用于跟单个用户通信的消息队列，用户发送消息
		Inbox:   make(chan Message, UserQueueSize),   //用于跟单个用户通信的消息队列,用户接收消息
	}

}

//Cancel Message
func (s *Subscription) Cancel() {
	s.cancel()
}

//Publish 这个表示用户订阅到感兴趣的主题的时候，同时也可以发送消息，
//Publish 但是，示例中不演示这个用途
//Publish 只有当channel无数据，且channel被close了，才会返回ok=false
//Publish a message to subscription queue
func (s *Subscription) Publish(msg Message) error {

	select {
	case <-s.ctx.Done():
		return errors.New("Topic has been closed")
	default:
		s.Ch <- msg
	}
	return nil
}

//Receive message
func (s *Subscription) Receive(out *Message) error {

	select {
	case <-s.ctx.Done():
		return errors.New("Topic has been closed")
	case <-time.After(time.Millisecond * 100):
		return errors.New("time out error")
	case *out = <-s.Inbox:
		return nil
	}
}

//Topic that user is interested in
//Topic should locate in MQ
type Topic struct {
	UserQueueSize  int
	Name           string
	Subscribers    map[uint64]Subscription //user list
	MessageHistory []Message               //当前主题的消息历史,实际项目中可能需要限定大小并设置过期时间
}

//Publish 只有当channel无数据，且channel被close了，才会返回ok=false
//Publish a message to subscription queue
func (t *Topic) Publish(msg Message) error {

	//将消息发布给当前Topic的所有人
	for usersID, subscription := range t.Subscribers {
		if subscription.ID == usersID {
			subscription.Inbox <- msg
		}
	}
	//save message history
	t.MessageHistory = append(t.MessageHistory, msg)

	fmt.Println("current histroy message count: ", len(t.MessageHistory))

	return nil
}

func (t *Topic) findUserSubscription(uid uint64, topicName string) (Subscription, bool) {
	// Get session or create one if it's the first

	if topicName != t.Name || t.Subscribers == nil || len(t.Subscribers) == 0 {
		return Subscription{}, false
	}
	if subscription, found := t.Subscribers[uid]; found {
		return subscription, true
	}
	return Subscription{}, false
}

//Subscribe a spec topic
func (t *Topic) Subscribe(uid uint64, topicName string) (Subscription, bool) {

	if t.Name != topicName {
		return Subscription{}, false
	}

	// Get session or create one if it's the first
	if _, found := t.findUserSubscription(uid, topicName); !found {
		if t.Subscribers == nil {
			t.Subscribers = make(map[uint64]Subscription)
		}
		t.Subscribers[uid] = newSubscription(uid, topicName, t.UserQueueSize)
	}

	return t.Subscribers[uid], true
}

//Unsubscribe remove Subscription
func (t *Topic) Unsubscribe(s Subscription) error {
	if _, found := t.findUserSubscription(s.ID, s.topicName); found {
		delete(t.Subscribers, s.ID)
	}
	return nil
}

//Delete topic
func (t *Topic) Delete() error {
	t.Subscribers = nil
	t.Name = ""
	t.MessageHistory = nil
	return nil
}
