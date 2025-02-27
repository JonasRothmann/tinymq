package tinymq

type connection struct {
	*queueManager
}

type Consumer struct {
	connection
}

func newConnection(qm *queueManager) connection {
	return connection{
		queueManager: qm,
	}
}
