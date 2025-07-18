package handler

import (
	"context"

	"github.com/FelipePn10/setspace/services/common/genproto/orders"
	"github.com/FelipePn10/setspace/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	// reister the OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}
	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}
