package gomore

import (
	"errors"
	"time"
)

//Message for msg in bus
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

	s.ch <- msg

	return nil
}

//Topic that user is interested in
type Topic struct {
	Subscribers    []Session
	MessageHistory []Message
}

//Subscribe a topic
func (t *Topic) Subscribe(uid uint64) (Subscription, error) {
	// Get session and create one if it's the first

	// Add session to the Topic & MessageHistory

	// Create a subscription

	return Subscription{}, nil
}

//Unsubscribe remove Subscription
func (t *Topic) Unsubscribe(Subscription) error {

	return nil
}

//Delete message
func (t *Topic) Delete() error {

	return nil
}
