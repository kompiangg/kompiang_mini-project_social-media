package websocketutils

type Message struct {
	MessageType int    `json:"message_type"`
	Content     string `json:"content"`
	SentTo      string `json:"sent_to"`
	SentBy      string `json:"sent_by"`
}
