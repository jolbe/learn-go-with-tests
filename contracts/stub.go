// ABOUTME: Programmable test double for CustomerStore.
// ABOUTME: Allows selectively overriding methods while delegating to a real implementation.
package contracts

import "context"

type CustomerStoreStub struct {
	delegate           CustomerStore
	CreateCustomerFunc func(ctx context.Context, name string) (Customer, error)
	GetCustomerFunc    func(ctx context.Context, id string) (Customer, error)
	UpdateCustomerFunc func(ctx context.Context, id string, name string) error
}

// assert CustomerStoreStub implements CustomerStore
var _ CustomerStore = &CustomerStoreStub{}

func NewCustomerStoreStub(delegate CustomerStore) *CustomerStoreStub {
	return &CustomerStoreStub{delegate: delegate}
}

func (s *CustomerStoreStub) CreateCustomer(ctx context.Context, name string) (Customer, error) {
	if s.CreateCustomerFunc != nil {
		return s.CreateCustomerFunc(ctx, name)
	}
	return s.delegate.CreateCustomer(ctx, name)
}

func (s *CustomerStoreStub) GetCustomer(ctx context.Context, id string) (Customer, error) {
	if s.GetCustomerFunc != nil {
		return s.GetCustomerFunc(ctx, id)
	}
	return s.delegate.GetCustomer(ctx, id)
}

func (s *CustomerStoreStub) UpdateCustomer(ctx context.Context, id string, name string) error {
	if s.UpdateCustomerFunc != nil {
		return s.UpdateCustomerFunc(ctx, id, name)
	}
	return s.delegate.UpdateCustomer(ctx, id, name)
}
