package contracts_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/contracts"
	"github.com/gregor-pifko/learn-go-with-tests/contracts/expect"
	"github.com/gregor-pifko/learn-go-with-tests/contracts/inmemory"
)

func TestInMemoryCustomerStore(t *testing.T) {
	contracts.CustomerStoreContract{NewStore: func() contracts.CustomerStore {
		return inmemory.NewCustomerStore()
	}}.Test(t)
}

func TestFailingUpdateCustomer(t *testing.T) {
	var (
		ctx         = context.Background()
		updateErr   = errors.New("failed to update customer")
		stub        = contracts.NewCustomerStoreStub(inmemory.NewCustomerStore())
	)
	stub.UpdateCustomerFunc = func(ctx context.Context, id string, name string) error {
		return updateErr
	}

	customer, err := stub.CreateCustomer(ctx, "Bob")
	expect.NoErr(t, err)

	expect.Err(t, stub.UpdateCustomer(ctx, customer.ID, "Robert"), updateErr)

	got, err := stub.GetCustomer(ctx, customer.ID)
	expect.NoErr(t, err)
	expect.Equal(t, "Bob", got.Name)
}
