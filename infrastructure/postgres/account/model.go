package account

import (
	"time"

	accountData "redbank/data/account"
)

type Model struct {
	ID        string `gorm:"primaryKey;type:varchar"`
	OwnerID   string `gorm:"type:varchar;index"`
	Type      string `gorm:"type:varchar"`
	Currency  string `gorm:"type:varchar(3)"`
	Balance   int64
	Status    string `gorm:"type:varchar"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Model) TableName() string { return "accounts" }

func toModel(e accountData.Account) Model {
	m := Model{
		OwnerID:  e.OwnerID,
		Type:     string(e.Type),
		Currency: e.Currency,
		Balance:  e.Balance,
		Status:   string(e.Status),
	}
	if e.ID != nil {
		m.ID = *e.ID
	}
	return m
}

func toEntity(m Model) accountData.Account {
	id := m.ID
	return accountData.Account{
		ID:        &id,
		OwnerID:   m.OwnerID,
		Type:      accountData.Type(m.Type),
		Currency:  m.Currency,
		Balance:   m.Balance,
		Status:    accountData.Status(m.Status),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
