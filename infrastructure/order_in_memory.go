package infrastructure

import (
	"context"
	"orderContext/domain/order"
	"sync"
)

var fakeOrders = make(map[string]*order.Order)

var lockMutex = &sync.RWMutex{}

type repository struct{}

func NewOrderRepository() order.Repository {
	return &repository{}
}

func (r *repository) GetOrders(_ context.Context) []*order.Order {
	lockMutex.RLock()
	defer lockMutex.RUnlock()

	var result []*order.Order

	for _, v := range fakeOrders {
		result = append(result, v)
	}

	return result
}

func (r *repository) Get(_ context.Context, id string) *order.Order {
	lockMutex.RLock()
	defer lockMutex.RUnlock()

	return fakeOrders[id]
}

func (r *repository) Update(_ context.Context, o *order.Order) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	fakeOrders[string(o.Id())] = o
}

func (r *repository) Create(_ context.Context, o *order.Order) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	fakeOrders[string(o.Id())] = o
}