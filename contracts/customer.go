// ABOUTME: Defines the Customer type, CustomerStore interface, and domain errors.
// ABOUTME: Core domain types for customer management contracts.
package contracts

import (
	"context"
	"errors"
)

type Customer struct {
	Name string
	ID   string
}

var ErrDaveIsForbidden = errors.New("dave is forbidden")

type CustomerStore interface {
	CreateCustomer(ctx context.Context, name string) (Customer, error)
	GetCustomer(ctx context.Context, id string) (Customer, error)
	UpdateCustomer(ctx context.Context, id string, name string) error
}
