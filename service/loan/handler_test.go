package loan_test

import (
	"context"
	"testing"

	"redbank/infrastructure/eventtest"
	"redbank/infrastructure/loantest"
	loanSvc "redbank/service/loan"
)

func newTestService(strategy loanSvc.EvaluationStrategy) *loanSvc.Service {
	repo := loantest.New()
	publisher := eventtest.New()
	return loanSvc.NewService(repo, publisher, strategy)
}

func TestService_Apply(t *testing.T) {
	// Arrange
	svc := newTestService(loanSvc.NewCreditScoreStrategy(650))
	input := loanSvc.ApplyInput{AccountID: "acc-1", Principal: 100000, InterestRate: 5.5, TermMonths: 12}

	// Act
	_, err := svc.Apply(context.Background(), input)

	// Assert
	if err != nil {
		t.Fatalf("Apply() unexpected error: %v", err)
	}
}

func TestService_Evaluate(t *testing.T) {
	// Arrange
	svc := newTestService(loanSvc.NewIncomeBasedStrategy(0.4))

	// Act
	_, err := svc.Evaluate(context.Background(), "loan-1")

	// Assert
	if err != nil {
		t.Fatalf("Evaluate() unexpected error: %v", err)
	}
}

func TestService_Get(t *testing.T) {
	// Arrange
	svc := newTestService(loanSvc.NewCreditScoreStrategy(650))

	// Act
	_, err := svc.Get(context.Background(), "loan-1")

	// Assert
	if err != nil {
		t.Fatalf("Get() unexpected error: %v", err)
	}
}

func TestService_ListByAccount(t *testing.T) {
	// Arrange
	svc := newTestService(loanSvc.NewCreditScoreStrategy(650))

	// Act
	_, err := svc.ListByAccount(context.Background(), "acc-1")

	// Assert
	if err != nil {
		t.Fatalf("ListByAccount() unexpected error: %v", err)
	}
}
