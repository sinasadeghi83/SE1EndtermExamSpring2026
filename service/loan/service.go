package loan

import (
	"redbank/data/event"
	"redbank/data/loan"
)

// Service implements the core business rules for the Loan domain, including
// underwriting via a pluggable EvaluationStrategy.
type Service struct {
	repo      loan.Repo
	publisher event.Publisher
	strategy  EvaluationStrategy
}

func NewService(repo loan.Repo, publisher event.Publisher, strategy EvaluationStrategy) *Service {
	return &Service{repo: repo, publisher: publisher, strategy: strategy}
}
