package transaction

import (
	"time"

	transactionData "redbank/data/transaction"
)

type Model struct {
	ID                    string  `gorm:"primaryKey;type:varchar"`
	AccountID             string  `gorm:"type:varchar;index"`
	CounterpartyAccountID *string `gorm:"type:varchar"`
	Type                  string  `gorm:"type:varchar"`
	Amount                int64
	Currency              string `gorm:"type:varchar(3)"`
	Status                string `gorm:"type:varchar"`
	Reference             string `gorm:"type:varchar"`
	CreatedAt             time.Time
}

func (Model) TableName() string { return "transactions" }

func toModel(e transactionData.Transaction) Model {
	m := Model{
		AccountID:             e.AccountID,
		CounterpartyAccountID: e.CounterpartyAccountID,
		Type:                  string(e.Type),
		Amount:                e.Amount,
		Currency:              e.Currency,
		Status:                string(e.Status),
		Reference:             e.Reference,
	}
	if e.ID != nil {
		m.ID = *e.ID
	}
	return m
}

func toEntity(m Model) transactionData.Transaction {
	id := m.ID
	return transactionData.Transaction{
		ID:                    &id,
		AccountID:             m.AccountID,
		CounterpartyAccountID: m.CounterpartyAccountID,
		Type:                  transactionData.Type(m.Type),
		Amount:                m.Amount,
		Currency:              m.Currency,
		Status:                transactionData.Status(m.Status),
		Reference:             m.Reference,
		CreatedAt:             m.CreatedAt,
	}
}
