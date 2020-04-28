package gomore

import (
	"errors"
	"time"
)

//Message for msg in Message bus
type Message struct {
	Alarm    int
	priority int
}

//User for user
type User struct {
	ID   uint64
	Name string
}

//Session of user
type Session struct {
	User      User
	Timestamp time.Time
}

//Subscription for user
type Subscription struct {
	ch    chan Message
	Inbox chan Message
}

//Publish a message to subscription queue
func (s *Subscription) Publish(msg Message) error {
	if _, ok := <-s.ch; !ok {
		return errors.New("Topic has been closed")
	}
	//用go channel 作为队列,接收消息
	s.ch <- msg

	return nil
}

//Topic that user is interested in
type Topic struct {
	uid            uint64
	Name           string
	Subscribers    []Session //user list
	MessageHistory []Message //当前主题的消息历史,实际项目中需要限定大小并设置过期时间
}

//MesssageQueue of manager all topics
type MesssageQueue struct {
	Topics map[uint64]*Topic
}

//String remove Subscription
func (t *Topic) String() string {
	return t.Name
}

//Subscribe a topic
func (t *Topic) Subscribe(uid uint64) (Subscription, error) {
	// Get session or create one if it's the first

	// Add session to the Topic & MessageHistory

	// Create a subscription

	return Subscription{}, nil
}

//Unsubscribe remove Subscription
func (t *Topic) Unsubscribe(Subscription) error {

	return nil
}

//Delete topic
func (t *Topic) Delete() error {

	return nil
}
