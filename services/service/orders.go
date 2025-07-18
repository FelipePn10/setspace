package service

import (
	"context"

	"github.com/FelipePn10/setspace/services/common/genproto/orders"
	"github.com/FelipePn10/setspace/services/orders/types"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	// store
}

func NewOrderService() types.OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDb = append(ordersDb, order)
	return nil
}
