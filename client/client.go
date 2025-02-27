package tinymq

import "github.com/JonasRothmann/tinymq"

type ClientProducer interface {
	SendMessage(msg tinymq.Message) error
}

type ClientConsumer interface {
	ReadMessage() <-chan tinymq.Message
}

type ClientMessage struct {
	tinymq.Message
}

func (m *ClientMessage) Ack() {

}

type ClientType int

const (
	ClientTypeTCP ClientType = iota
	ClientTypeHTTP
	ClientTypeInProcess
)

type Client interface {
	CreateProducer() (*ClientProducer, error)
	CreateConsumer() (*ClientConsumer, error)
}
