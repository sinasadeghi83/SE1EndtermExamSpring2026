package loan_test

import (
	"testing"

	"redbank/data/loan"
	loanSvc "redbank/service/loan"
)

func TestCreditScoreStrategy_Evaluate(t *testing.T) {
	// Arrange
	strategy := loanSvc.NewCreditScoreStrategy(650)
	applicant := loanSvc.Applicant{CreditScore: 700, RequestedPrincipal: 50000}

	// Act
	_, err := strategy.Evaluate(applicant, loan.Loan{})

	// Assert
	if err != nil {
		t.Fatalf("Evaluate() unexpected error: %v", err)
	}
}

func TestIncomeBasedStrategy_Evaluate(t *testing.T) {
	// Arrange
	strategy := loanSvc.NewIncomeBasedStrategy(0.4)
	applicant := loanSvc.Applicant{MonthlyIncome: 5000, ExistingDebt: 1000, RequestedPrincipal: 50000}

	// Act
	_, err := strategy.Evaluate(applicant, loan.Loan{})

	// Assert
	if err != nil {
		t.Fatalf("Evaluate() unexpected error: %v", err)
	}
}

func TestCompositeStrategy_Evaluate(t *testing.T) {
	// Arrange
	strategy := loanSvc.NewCompositeStrategy(
		loanSvc.NewCreditScoreStrategy(650),
		loanSvc.NewIncomeBasedStrategy(0.4),
	)
	applicant := loanSvc.Applicant{CreditScore: 700, MonthlyIncome: 5000, ExistingDebt: 1000, RequestedPrincipal: 50000}

	// Act
	_, err := strategy.Evaluate(applicant, loan.Loan{})

	// Assert
	if err != nil {
		t.Fatalf("Evaluate() unexpected error: %v", err)
	}
}
