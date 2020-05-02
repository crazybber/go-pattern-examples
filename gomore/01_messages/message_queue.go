package messaging

//Queue hold all topics
type Queue struct {
	Topics map[string]Topic //topic ID<-----> topic Object
}

//AddTopic to Queue
func (q *Queue) AddTopic(topicName string, topicUserQueueSize int) Topic {
	if q.Topics == nil {
		q.Topics = make(map[string]Topic)
	}
	if _, found := q.Topics[topicName]; !found {
		q.Topics[topicName] = Topic{UserQueueSize: topicUserQueueSize, Name: topicName}
	}
	return q.Topics[topicName]
}
