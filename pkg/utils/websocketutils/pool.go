package websocketutils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Message    chan *Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Message:    make(chan *Message),
	}
}

func (pool *Pool) Start() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			pool.Close()
			return
		case client := <-pool.Register:
			pool.Clients[client] = true
			for poolClient := range pool.Clients {
				poolClient.Conn.WriteJSON(fmt.Sprintf("%s online", client.Username))
			}
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			for poolClient := range pool.Clients {
				poolClient.Conn.WriteJSON(fmt.Sprintf("%s going offline", client.Username))
			}
		case message := <-pool.Message:
			if message.SentTo == "" {
				for client := range pool.Clients {
					if err := client.Conn.WriteJSON(message.Content); err != nil {
						log.Println(err)
						return
					}
				}
			} else {
				for client := range pool.Clients {
					if client.Username == message.SentTo {
						if err := client.Conn.WriteJSON(message); err != nil {
							log.Println(err)
							return
						}
					}
				}
			}
		}
	}
}

func (pool *Pool) Close() {
	for client := range pool.Clients {
		client.Close()
	}
	close(pool.Message)
	close(pool.Unregister)
	close(pool.Register)
	log.Println("[INFO] WebSocket connection closed gracefully")
}
