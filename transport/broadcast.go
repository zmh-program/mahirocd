package main

import "github.com/gofiber/websocket/v2"

type BroadcastManager struct {
	connections []*websocket.Conn
	channel     chan []byte
}

func NewBroadcastManager() *BroadcastManager {
	return &BroadcastManager{
		connections: make([]*websocket.Conn, 0),
		channel:     make(chan []byte),
	}
}

func (b *BroadcastManager) AddConnection(conn *websocket.Conn) {
	b.connections = append(b.connections, conn)
}

func (b *BroadcastManager) RemoveConnection(conn *websocket.Conn) {
	for i, c := range b.connections {
		if c == conn {
			b.connections = append(b.connections[:i], b.connections[i+1:]...)
			break
		}
	}
}

func (b *BroadcastManager) Broadcast() {
	for {
		select {
		case message := <-b.channel:
			for _, conn := range b.connections {
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					b.RemoveConnection(conn)
				}
			}
		}
	}
}

func (b *BroadcastManager) Send(message []byte) {
	b.channel <- message
}

func (b *BroadcastManager) Close() {
	close(b.channel)
}

func (b *BroadcastManager) Len() int {
	return len(b.connections)
}

func (b *BroadcastManager) IsEmpty() bool {
	return b.Len() == 0
}

func (b *BroadcastManager) Exec() {
	go b.Broadcast()
}

func (b *BroadcastManager) GetConnections() []*websocket.Conn {
	return b.connections
}

func (b *BroadcastManager) GetChannel() chan []byte {
	return b.channel
}

func (b *BroadcastManager) GetConnection(index int) *websocket.Conn {
	return b.connections[index]
}
