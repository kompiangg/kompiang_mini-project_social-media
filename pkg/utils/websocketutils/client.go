package websocketutils

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Username string
	Conn     *websocket.Conn
	Pool     *Pool
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var pJson Message
		err = json.Unmarshal(p, &pJson)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		message := Message{
			MessageType: messageType,
			Content:     pJson.Content,
			SentTo:      pJson.SentTo,
			SentBy:      c.Username,
		}

		c.Pool.Message <- &message
	}
}

func (c *Client) Close() {
	c.Pool.Unregister <- c
	c.Conn.Close()
}
