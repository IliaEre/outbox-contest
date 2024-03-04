package domain

type Message struct {
	UserID  string `json:"user_id"`
	Payload string `json:"payload"`
}
