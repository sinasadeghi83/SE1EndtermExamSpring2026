package loan

import "time"

type Status string

const (
	StatusPending     Status = "pending"
	StatusUnderReview Status = "under_review"
	StatusApproved    Status = "approved"
	StatusRejected    Status = "rejected"
	StatusDisbursed   Status = "disbursed"
	StatusClosed      Status = "closed"
)

// Loan is the core domain entity representing a loan request/contract tied to an Account.
type Loan struct {
	ID              *string
	AccountID       string
	Principal       int64
	InterestRate    float64
	TermMonths      int
	Status          Status
	EvaluationScore *float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func New(accountID string, principal int64, interestRate float64, termMonths int) *Loan {
	return &Loan{
		AccountID:    accountID,
		Principal:    principal,
		InterestRate: interestRate,
		TermMonths:   termMonths,
		Status:       StatusPending,
	}
}
