package loan

import "redbank/data/loan"

// Applicant carries the inputs an evaluation strategy needs that don't
// belong on the persisted Loan entity itself.
type Applicant struct {
	CreditScore        int
	MonthlyIncome      int64
	ExistingDebt       int64
	RequestedPrincipal int64
}

// EvaluationResult is the outcome produced by a LoanEvaluationStrategy.
type EvaluationResult struct {
	Approved bool
	Score    float64
	Reason   string
}

// EvaluationStrategy is the Strategy-pattern contract for pluggable loan
// underwriting rules. Concrete strategies are selected/composed by the
// service layer at evaluation time.
type EvaluationStrategy interface {
	Name() string
	Evaluate(applicant Applicant, l loan.Loan) (EvaluationResult, error)
}

// CreditScoreStrategy evaluates a loan primarily against the applicant's credit score.
type CreditScoreStrategy struct {
	MinimumScore int
}

func NewCreditScoreStrategy(minimumScore int) *CreditScoreStrategy {
	return &CreditScoreStrategy{MinimumScore: minimumScore}
}

func (s *CreditScoreStrategy) Name() string { return "credit_score" }

func (s *CreditScoreStrategy) Evaluate(applicant Applicant, l loan.Loan) (EvaluationResult, error) {
	return EvaluationResult{}, nil
}

// IncomeBasedStrategy evaluates a loan against the applicant's debt-to-income ratio.
type IncomeBasedStrategy struct {
	MaxDebtToIncomeRatio float64
}

func NewIncomeBasedStrategy(maxDebtToIncomeRatio float64) *IncomeBasedStrategy {
	return &IncomeBasedStrategy{MaxDebtToIncomeRatio: maxDebtToIncomeRatio}
}

func (s *IncomeBasedStrategy) Name() string { return "income_based" }

func (s *IncomeBasedStrategy) Evaluate(applicant Applicant, l loan.Loan) (EvaluationResult, error) {
	return EvaluationResult{}, nil
}

// CompositeStrategy runs multiple strategies and combines their results.
// Selected/ordered by the service layer; combination policy is decided there.
type CompositeStrategy struct {
	Strategies []EvaluationStrategy
}

func NewCompositeStrategy(strategies ...EvaluationStrategy) *CompositeStrategy {
	return &CompositeStrategy{Strategies: strategies}
}

func (s *CompositeStrategy) Name() string { return "composite" }

func (s *CompositeStrategy) Evaluate(applicant Applicant, l loan.Loan) (EvaluationResult, error) {
	return EvaluationResult{}, nil
}
