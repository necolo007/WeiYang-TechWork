package Model

type Message struct {
	Content string `json:"content,omitempty"`
	Receive string `json:"receive,omitempty"`
	Sender  string `json:"sender,omitempty"`
}
