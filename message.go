package tinymq

type MessageID int64

type Message struct {
	ID      MessageID
	Content []byte
	Ack     bool
}
