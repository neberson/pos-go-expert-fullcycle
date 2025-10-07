package service

import (
	"context"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/infra/grpc/pb"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
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

func (s *OrderService) ListOrders(ctx context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	output, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order
	for _, orderDTO := range output {
		order := &pb.Order{
			Id:         orderDTO.ID,
			Price:      float32(orderDTO.Price),
			Tax:        float32(orderDTO.Tax),
			FinalPrice: float32(orderDTO.FinalPrice),
		}
		orders = append(orders, order)
	}

	return &pb.ListOrdersResponse{
		Orders: orders,
	}, nil
}
