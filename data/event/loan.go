package event

import "time"

const (
	NameLoanApproved = "loan.approved"
	NameLoanRejected = "loan.rejected"
)

type LoanApproved struct {
	LoanID string
	At     time.Time
}

func NewLoanApproved(loanID string) LoanApproved {
	return LoanApproved{LoanID: loanID, At: time.Now()}
}

func (e LoanApproved) Name() string          { return NameLoanApproved }
func (e LoanApproved) OccurredAt() time.Time { return e.At }

type LoanRejected struct {
	LoanID string
	Reason string
	At     time.Time
}

func NewLoanRejected(loanID, reason string) LoanRejected {
	return LoanRejected{LoanID: loanID, Reason: reason, At: time.Now()}
}

func (e LoanRejected) Name() string          { return NameLoanRejected }
func (e LoanRejected) OccurredAt() time.Time { return e.At }
