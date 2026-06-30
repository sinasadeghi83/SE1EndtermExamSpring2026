package account_test

import (
	"context"
	"testing"

	"redbank/infrastructure/accounttest"
	"redbank/infrastructure/eventtest"
	accountSvc "redbank/service/account"
)

func newTestService() *accountSvc.Service {
	repo := accounttest.New()
	publisher := eventtest.New()
	return accountSvc.NewService(repo, publisher)
}

func TestService_Open(t *testing.T) {
	// Arrange
	svc := newTestService()
	input := accountSvc.OpenInput{OwnerID: "owner-1", Type: "checking", Currency: "USD"}

	// Act
	_, err := svc.Open(context.Background(), input)

	// Assert
	if err != nil {
		t.Fatalf("Open() unexpected error: %v", err)
	}
}

func TestService_Get(t *testing.T) {
	// Arrange
	svc := newTestService()

	// Act
	_, err := svc.Get(context.Background(), "acc-1")

	// Assert
	if err != nil {
		t.Fatalf("Get() unexpected error: %v", err)
	}
}

func TestService_ListByOwner(t *testing.T) {
	// Arrange
	svc := newTestService()

	// Act
	_, err := svc.ListByOwner(context.Background(), "owner-1")

	// Assert
	if err != nil {
		t.Fatalf("ListByOwner() unexpected error: %v", err)
	}
}

func TestService_Close(t *testing.T) {
	// Arrange
	svc := newTestService()

	// Act
	err := svc.Close(context.Background(), "acc-1")

	// Assert
	if err != nil {
		t.Fatalf("Close() unexpected error: %v", err)
	}
}
