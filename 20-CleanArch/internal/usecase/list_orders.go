package usecase

import "github.com/rgoncalvesrr/desafio-clean-arch/internal/entity"

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	ordersOutPut := []OrderOutputDTO{}
	orders, err := c.OrderRepository.FindAll()

	if err != nil {
		return ordersOutPut, err
	}

	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		ordersOutPut = append(ordersOutPut, dto)
	}

	return ordersOutPut, nil
}
