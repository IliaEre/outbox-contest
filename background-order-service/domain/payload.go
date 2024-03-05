package domain

// Outbox - Define Go types to match the JSON schema
type Outbox struct {
	Schema  Schema  `json:"schema"`
	Payload Payload `json:"payload"`
}

type Schema struct {
	Type     string  `json:"type"`
	Fields   []Field `json:"fields"`
	Optional bool    `json:"optional"`
	Name     string  `json:"name"`
	Version  int     `json:"version"`
	Field    string  `json:"field,omitempty"`
}

type Field struct {
	Type     string `json:"type"`
	Optional bool   `json:"optional"`
	Field    string `json:"field"`
	Name     string `json:"name,omitempty"`
	Version  int    `json:"version,omitempty"`
}

type Payload struct {
	ID              int    `json:"id"`
	UUID            string `json:"uuid"`
	UserID          string `json:"user_id"`
	TransactionCode string `json:"transaction_code"`
	JSON            string `json:"json"`
	CreatedAt       int64  `json:"created_at,omitempty"`
	UpdatedAt       int64  `json:"updated_at,omitempty"`
}
