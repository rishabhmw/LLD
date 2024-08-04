package models

type Message struct {
	ID    string `json:"id"`
	Data  []byte `json:"data"`
	Topic string `json:"topic"`
}
