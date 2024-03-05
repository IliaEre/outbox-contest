package domain

import "encoding/json"

type Order struct {
	Name    string       `json:"name"`
	Surname string       `json:"surname"`
	Details OrderDetails `json:"details"`
}

type OrderDetails struct {
	Code            int    `json:"code"`
	OperationCode   int    `json:"operation_code"`
	TransactionCode string `json:"transaction_code"`
}

func (o *Order) ToJSON() (string, error) {
	jsonData, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
