package tinymq

type LocalClient struct {
	mq *TinyMQ
}

func NewLocalClient(*mq) (*Client, error) {
	client := Client{
		clientType: ClientTypeInProcess,
	}

	return &client, nil
}
