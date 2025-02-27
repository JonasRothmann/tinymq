package tinymq

import (
	"fmt"
	"sync"
)

type Queue struct {
	Name     string
	qm       *queueManager
	Messages []*Message
}

type queueManager struct {
	mu     sync.Mutex
	queues map[string]*Queue
}

func newQueueManager() *queueManager {
	return &queueManager{
		queues: map[string]*Queue{},
	}
}

func (qm *queueManager) GetQueue(name string) *Queue {
	qm.mu.Lock()
	defer qm.mu.Unlock()

	if queue, ok := qm.queues[name]; ok == true {
		return queue
	}

	queue := Queue{
		qm:       qm,
		Name:     name,
		Messages: []*Message{},
	}
	qm.queues[name] = &queue

	return &queue
}

// TODO: Do we want to return a copy here?
func (qm *queueManager) Queues() map[string]*Queue {
	return qm.queues
}

func (qm *queueManager) DebugQueues() string {
	var content string

	for name, queue := range qm.queues {
		content += fmt.Sprintf("Name: %s, Message Count: %d\n", name, len(queue.Messages))
	}

	return content
}
