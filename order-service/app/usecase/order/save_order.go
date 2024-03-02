package order

import (
	"order-service/domain"
)

// Ensure OrderService implements Service interface
var _ Service = &OrderService{}

func (s *OrderService) Save(userUUID string, order *domain.Order) (*int32, error) {
	id, err := s.repo.CreateOrder(userUUID, *order)
	if err != nil {
		return nil, err
	}

	return id, nil
}
