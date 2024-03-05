package order

import (
	"order-service/domain"
)

type Service interface {
	Save(uuid string, order *domain.Order) (*int32, error)
}
