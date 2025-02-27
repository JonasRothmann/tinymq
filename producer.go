package tinymq

type Producer struct {
	connection
}

func NewProducer(r *TinyMQ) *Producer {
	return &Producer{
		connection: newConnection(r.queueManager),
	}
}
