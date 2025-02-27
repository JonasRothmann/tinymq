package receivers

import (
	"bufio"
	"context"
	"fmt"
	"net"
)

type TCPReciever struct {
	net.Listener

	// connections []net.Conn <- might not be needed
	msgChan chan []byte
	ctx     context.Context
}

func NewTCPReciever(addr string) (*TCPReciever, error) {
	ctx := context.Background()

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &TCPReciever{
		msgChan:  make(chan []byte),
		ctx:      ctx,
		Listener: listener,
	}, nil
}

func (r *TCPReciever) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Accept incoming connections
			conn, err := r.Listener.Accept()
			if err != nil {
				fmt.Printf("failed to accept incoming connection: %s", err)
				continue
			}
			go r.newConnection(ctx, conn)
		}

	}
}

func (r *TCPReciever) Close() {
	var cancel context.CancelFunc
	r.ctx, cancel = context.WithCancel(r.ctx)
	cancel()
	close(r.msgChan)
	r.Listener.Close()
}

func (r *TCPReciever) newConnection(ctx context.Context, conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("yo")
		}

		if reader.Buffered() > 0 {
			var b []byte
			if n, err := reader.Read(b); err == nil {
				fmt.Printf("Received %d bytes", n)
				r.msgChan <- b
			}
		} else {
			fmt.Println("Nothing to read")
		}
	}
}

func (r *TCPReciever) Listen() <-chan []byte {
	return r.msgChan
}
