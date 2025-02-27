package receivers

import (
	"testing"

	"github.com/JonasRothmann/tinymq"
	"github.com/stretchr/testify/assert"
)

func TestTCP(t *testing.T) {
	tcpReceiver, err := NewTCPReciever("localhost:3456")
	assert.NoError(t, err)

	mq := tinymq.New(tcpReceiver)

	client := client.NewTCPClient("localhost:3456")
	// or client := mq.NewClient() for in-process

	test := mq.GetQueue("test")
	assert.Len(t, mq.Queues(), 1)

	test2 := mq.GetQueue("test-2")
	assert.Len(t, mq.Queues(), 2)

	test2 = mq.GetQueue("test-2")
	assert.Len(t, mq.Queues(), 2)

}
