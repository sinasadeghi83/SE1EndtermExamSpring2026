package loan

import (
	"time"

	loanData "redbank/data/loan"
)

type Model struct {
	ID              string `gorm:"primaryKey;type:varchar"`
	AccountID       string `gorm:"type:varchar;index"`
	Principal       int64
	InterestRate    float64
	TermMonths      int
	Status          string `gorm:"type:varchar"`
	EvaluationScore *float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (Model) TableName() string { return "loans" }

func toModel(e loanData.Loan) Model {
	m := Model{
		AccountID:       e.AccountID,
		Principal:       e.Principal,
		InterestRate:    e.InterestRate,
		TermMonths:      e.TermMonths,
		Status:          string(e.Status),
		EvaluationScore: e.EvaluationScore,
	}
	if e.ID != nil {
		m.ID = *e.ID
	}
	return m
}

func toEntity(m Model) loanData.Loan {
	id := m.ID
	return loanData.Loan{
		ID:              &id,
		AccountID:       m.AccountID,
		Principal:       m.Principal,
		InterestRate:    m.InterestRate,
		TermMonths:      m.TermMonths,
		Status:          loanData.Status(m.Status),
		EvaluationScore: m.EvaluationScore,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}
