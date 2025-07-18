package types

import (
	"context"

	"github.com/FelipePn10/setspace/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *orders.Order) error
}
