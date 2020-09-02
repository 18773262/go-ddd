package application

import (
	"orderContext/domain/order"
)

type OrderService interface {
	Pay(orderId string)

	Ship(orderId string) error

	Cancel(orderId string)
}

type service struct {
	repository order.OrderRepository
}

func NewOrderService() OrderService {
	return &service{repository: order.NewOrderRepository()}
}

func (s *service) Pay(orderId string) {
	order := s.repository.Get(orderId)
	order.Pay()
	s.repository.Update(order)

}

func (s *service) Cancel(orderId string) {
	order := s.repository.Get(orderId)
	order.Pay()
	s.repository.Update(order)
}

func (s *service) Ship(orderId string) error {
	order := s.repository.Get(orderId)
	result := order.Ship()
	if result != nil {
		return result
	}
	s.repository.Update(order)

	return nil
}
