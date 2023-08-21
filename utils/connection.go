package utils

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"os/signal"
	"time"
)

type Connection struct {
	url      string
	conn     *websocket.Conn
	closed   bool
	callback func(message []byte)
	tick     int
}

func NewConnection(url string, callback func(message []byte)) *Connection {
	conn := &Connection{
		url:      url,
		closed:   true,
		callback: callback,
		conn:     nil,
		tick:     0,
	}
	return conn
}

func (c *Connection) Connect() error {
	if !c.closed {
		return c.Close()
	}

	conn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		fmt.Println("Websocket connected failed, retrying... [", c.tick, "]")
		c.TryReconnect()
		return err
	}

	c.closed = false
	c.tick = 0
	fmt.Println("Websocket connected successfully to", c.url)
	c.conn = conn
	c.Run()

	return nil
}

func (c *Connection) ExecWithBlock() {
	_ = c.Connect()
}

func (c *Connection) ExecWithoutBlock() {
	go c.ExecWithBlock()
}

func (c *Connection) Run() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for !c.closed {
			_, message := c.ReadMessage()
			if len(message) == 0 {
				if c.closed {
					return
				}
				continue
			}
			c.callback(message)
		}
	}()

	for !c.closed {
		select {
		case <-ticker.C:
			c.Ping()
		case <-interrupt:
			c.CloseBySelf()
			fmt.Println("Websocket connection closed by user.")
			os.Exit(0)
		}
	}

	go func() {
		// receive interrupt signal
		<-interrupt
		c.CloseBySelf()
		fmt.Println("Websocket connection closed by user.")
		os.Exit(0)
	}()
}

func (c *Connection) ReadMessage() (messageType int, p []byte) {
	if c.closed {
		return 0, nil
	}

	defer func() {
		if err := recover(); err != nil {
			c.TryReconnect()
		}
	}()
	t, p, err := c.conn.ReadMessage()
	if err != nil {
		c.TryReconnect()
	}
	return t, p
}

func (c *Connection) WriteMessage(messageType int, data []byte) bool {
	if c.closed {
		return false
	}
	err := c.conn.WriteMessage(messageType, data)
	if err != nil {
		c.TryReconnect()
		return false
	}
	return true
}

func (c *Connection) Ping() bool {
	return c.WriteMessage(websocket.TextMessage, []byte("ping"))
}

func (c *Connection) CloseBySelf() {
	if c.closed {
		return
	}
	c.closed = true
	_ = c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

func (c *Connection) Close() error {
	if c.closed {
		return nil
	}
	c.closed = true
	fmt.Println("Websocket connection closed.")
	return c.conn.Close()
}

func (c *Connection) IsClosed() bool {
	return c.closed
}

func (c *Connection) IsConnected() bool {
	return c.conn != nil
}

func (c *Connection) Reconnect() error {
	if err := c.Close(); err != nil {
		return err
	}
	return c.Connect()
}

func (c *Connection) GetConnection() *websocket.Conn {
	return c.conn
}

func (c *Connection) TryReconnect() {
	c.tick++
	SetTimeoutSync(func() {
		_ = c.Connect()
	}, 3000)
}
