package handler

import (
	"order-service/app/usecase/order"
)

type OrderHandler struct {
	service *order.OrderService
}

func NewOrderHandler(s *order.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}
