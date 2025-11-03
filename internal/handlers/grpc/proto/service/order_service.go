package service

import (
	"context"

	"github.com/fpessoa64/desafio_clean_arch/internal/entities"
	orderpb "github.com/fpessoa64/desafio_clean_arch/internal/handlers/grpc/proto"
	"github.com/fpessoa64/desafio_clean_arch/internal/usecase"
)

type OrderServiceServer struct {
	orderpb.UnimplementedOrderServiceServer
	UC *usecase.OrderUsecase
}

func NewOrderServiceServer(uc *usecase.OrderUsecase) *OrderServiceServer {
	return &OrderServiceServer{UC: uc}
}

func (s *OrderServiceServer) ListOrders(ctx context.Context, req *orderpb.ListOrdersRequest) (*orderpb.ListOrdersResponse, error) {
	orders, err := s.UC.List(ctx)
	if err != nil {
		return nil, err
	}
	resp := &orderpb.ListOrdersResponse{Orders: make([]*orderpb.Order, 0, len(orders))}
	for _, o := range orders {
		resp.Orders = append(resp.Orders, entityToProto(o))
	}
	return resp, nil
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	order := &entities.Order{
		Name:   req.GetName(),
		Amount: req.GetAmount(),
		Status: req.GetStatus(),
	}
	err := s.UC.Create(ctx, order)
	if err != nil {
		return nil, err
	}
	return &orderpb.CreateOrderResponse{Order: entityToProto(*order)}, nil
}

func entityToProto(o entities.Order) *orderpb.Order {
	return &orderpb.Order{
		Id:        o.ID,
		Name:      o.Name,
		Amount:    o.Amount,
		Status:    o.Status,
		CreatedAt: o.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
