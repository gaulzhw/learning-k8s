package tcp

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/gaulzhw/learning-k8s/internal/sync/atomic"
	"github.com/gaulzhw/learning-k8s/internal/sync/wait"
)

type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

func (c *EchoClient) Close() error {
	c.Waiting.WaitWithTimeout(10 * time.Second)
	_ = c.Conn.Close()
	return nil
}

type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (h *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if h.closing.Get() {
		_ = conn.Close()
		return
	}

	c := &EchoClient{
		Conn: conn,
	}
	h.activeConn.Store(c, struct{}{})

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Print("Connection close")
				h.activeConn.Delete(c)
			} else {
				log.Print(err)
			}
			return
		}

		c.Waiting.Add(1)
		_, _ = conn.Write([]byte(msg))
		c.Waiting.Done()
	}
}

func (h *EchoHandler) Close() error {
	log.Print("handler shutting down")
	h.closing.Set(true)
	h.activeConn.Range(func(key, value any) bool {
		client := key.(*EchoClient)
		_ = client.Close()
		return true
	})
	return nil
}
