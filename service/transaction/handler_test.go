package transaction_test

import (
	"context"
	"testing"

	"redbank/infrastructure/accounttest"
	"redbank/infrastructure/eventtest"
	"redbank/infrastructure/transactiontest"
	transactionSvc "redbank/service/transaction"
)

func newTestService() *transactionSvc.Service {
	repo := transactiontest.New()
	accountRepo := accounttest.New()
	publisher := eventtest.New()
	return transactionSvc.NewService(repo, accountRepo, publisher)
}

func TestService_Post(t *testing.T) {
	// Arrange
	svc := newTestService()
	input := transactionSvc.PostInput{AccountID: "acc-1", Type: "deposit", Amount: 1000, Currency: "USD"}

	// Act
	_, err := svc.Post(context.Background(), input)

	// Assert
	if err != nil {
		t.Fatalf("Post() unexpected error: %v", err)
	}
}

func TestService_Transfer(t *testing.T) {
	// Arrange
	svc := newTestService()
	input := transactionSvc.TransferInput{FromAccountID: "acc-1", ToAccountID: "acc-2", Amount: 500, Currency: "USD"}

	// Act
	_, _, err := svc.Transfer(context.Background(), input)

	// Assert
	if err != nil {
		t.Fatalf("Transfer() unexpected error: %v", err)
	}
}

func TestService_Get(t *testing.T) {
	// Arrange
	svc := newTestService()

	// Act
	_, err := svc.Get(context.Background(), "txn-1")

	// Assert
	if err != nil {
		t.Fatalf("Get() unexpected error: %v", err)
	}
}

func TestService_ListByAccount(t *testing.T) {
	// Arrange
	svc := newTestService()

	// Act
	_, err := svc.ListByAccount(context.Background(), "acc-1")

	// Assert
	if err != nil {
		t.Fatalf("ListByAccount() unexpected error: %v", err)
	}
}
