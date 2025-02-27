package tinymq

import (
	"context"
	"fmt"
)

type TinyMQ struct {
	ctx       context.Context
	receivers []Receiver

	messages []*Message

	*queueManager

	consumers []*Consumer
	producers []*Producer
}

type Receiver interface {
	Start(ctx context.Context)
	Close()
	Listen() <-chan []byte
}

func New(receivers ...Receiver) *TinyMQ {
	ctx := context.Background()

	return &TinyMQ{
		ctx:          ctx,
		messages:     []*Message{},
		consumers:    []*Consumer{},
		producers:    []*Producer{},
		receivers:    receivers,
		queueManager: newQueueManager(),
	}
}

func (r *TinyMQ) CreateProducer() *Producer {
	return &Producer{
		connection: connection{
			queueManager: r.queueManager,
		},
	}
}

func (r *TinyMQ) Close() {
	var close context.CancelFunc
	r.ctx, close = context.WithCancel(r.ctx)
	close()
}

func (r *TinyMQ) ReceiveMessage() <-chan Message {
	for {
		select {
		case msg := <-r.receiver.Listen():
			fmt.Printf("received msg: %s", msg)
		}
	}
}
