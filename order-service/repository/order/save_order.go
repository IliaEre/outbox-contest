package order

import (
	"github.com/fossoreslp/go-uuid-v4"
	"order-service/domain"
)

const StatusCreated = "CREATED"

var query = `
	INSERT INTO orders (name, surname, code, operation_code, transaction_code, status)
	VALUES ($1, $2, $3, $4, $5, $6)
	returning id;
	`

func (r *OrderRepository) CreateOrder(userUUID string, order domain.Order) (*int32, error) {
	det := order.Details
	var id int32

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(
		query,
		order.Name, order.Surname, det.Code, det.OperationCode, det.TransactionCode, StatusCreated).Scan(&id)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save to outbox
	err = r.createOutbox(userUUID, order)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &id, tx.Commit()
}

func (r *OrderRepository) createOutbox(userUUID string, order domain.Order) error {
	uuid, _ := uuid.NewString()
	json, _ := order.ToJSON()

	_, err := r.db.Exec(
		"INSERT INTO outbox (uuid, user_id, transaction_code, json) VALUES ($1, $2, $3, $4)",
		uuid, userUUID, order.Details.TransactionCode, json,
	)

	return err
}
