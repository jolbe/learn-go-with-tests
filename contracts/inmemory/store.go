// ABOUTME: In-memory implementation of CustomerStore for testing.
// ABOUTME: Stores customers in a map, keyed by auto-incrementing ID.
package inmemory

import (
	"context"
	"strconv"

	"github.com/gregor-pifko/learn-go-with-tests/contracts"
)

func NewCustomerStore() *CustomerStore {
	return &CustomerStore{customers: make(map[string]contracts.Customer)}
}

type CustomerStore struct {
	i         int
	customers map[string]contracts.Customer
}

func (s *CustomerStore) CreateCustomer(ctx context.Context, name string) (contracts.Customer, error) {
	if name == "Dave" {
		return contracts.Customer{}, contracts.ErrDaveIsForbidden
	}

	newCustomer := contracts.Customer{
		Name: name,
		ID:   strconv.Itoa(s.i),
	}
	s.customers[newCustomer.ID] = newCustomer
	s.i++
	return newCustomer, nil
}

func (s *CustomerStore) GetCustomer(ctx context.Context, id string) (contracts.Customer, error) {
	return s.customers[id], nil
}

func (s *CustomerStore) UpdateCustomer(ctx context.Context, id string, name string) error {
	customer := s.customers[id]
	customer.Name = name
	s.customers[id] = customer
	return nil
}
