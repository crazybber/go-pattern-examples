package gomore

import (
	"errors"
	"time"
)

//Message for msg in Message bus
type Message struct {
	//Type 类型[code :1,2,3,4]
	Type int
	Text string
	From Session //消息来源
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
//Subscription is a session
type Subscription struct {
	Session
	ch    chan<- Message //发送队列
	Inbox <-chan Message //接收消息的队列
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
//Topic locate in MQ
type Topic struct {
	Name           string
	Subscribers    map[uint64]Session //user list
	MessageHistory []Message          //当前主题的消息历史,实际项目中可能需要限定大小并设置过期时间
	subscription   Subscription
}

//Queue hold all topics
type Queue struct {
	Topics map[string]*Topic //topic ID<-----> topic Object
}

//String remove Subscription
func (t *Topic) String() string {
	return t.Name
}

func (t *Topic) findSession(uid uint64) (Session, bool) {
	// Get session or create one if it's the first
	var Session session
	if t.Subscribers == nil || len(t.Subscribers) == 0 {
		return Session{}, false
	}
	if session, found := t.Subscribers[uid]; found {
		return session, true
	}
	return Session{}, false
}

func (t *Topic) addSession(uid uint64) Session {

	var Session session
	// Get session or create one if it's the first
	if session, found := t.findSession(uid); !found {
		if t.Subscribers == ni {
			t.Subscribers = make(map[uint64]Session)
		}
		session = Session{User{uid, "no name"}, time.Now()}
		t.Subscribers[uid] = session
	}
	return session
}

//Subscribe a spec topic
func (t *Topic) Subscribe(uid uint64) (Subscription, error) {

	session := t.addSession(uid)

	// Create a subscription from copy
	subscription := Subscription{session, t.subscription.ch, t.subscription.Inbox}

	return subscription, nil
}

//Unsubscribe remove Subscription
func (t *Topic) Unsubscribe(s Subscription) error {
	if t.findSession(s.User.ID) {
		delete(t.Subscribers, s.User.ID)
	}
	return nil
}

//Delete topic
func (t *Topic) Delete() error {
	t.Subscribers = nil
	t.Name = ""
	t.MessageHistory = nil
	t.subscription = Subscription{}
	return nil
}
