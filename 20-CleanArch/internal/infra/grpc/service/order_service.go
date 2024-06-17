package service

import (
	"context"

	"github.com/rgoncalvesrr/desafio-clean-arch/internal/infra/grpc/pb"
	"github.com/rgoncalvesrr/desafio-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	listOrderUseCase usecase.ListOrderUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(context.Context, *pb.Blank) (*pb.OrderList, error) {

	orders, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var ordersResp []*pb.CreateOrderResponse

	for _, order := range orders {
		ordersResp = append(ordersResp, &pb.CreateOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.FinalPrice),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	result := &pb.OrderList{
		Orders: ordersResp,
	}

	return result, nil
}
