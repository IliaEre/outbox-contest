package order

import "order-service/repository/order"

type OrderService struct {
	repo *order.OrderRepository
}

func NewOrderService(r *order.OrderRepository) *OrderService {
	return &OrderService{repo: r}
}
