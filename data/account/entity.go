package account

import "time"

type Type string

const (
	TypeChecking Type = "checking"
	TypeSavings  Type = "savings"
)

type Status string

const (
	StatusActive Status = "active"
	StatusFrozen Status = "frozen"
	StatusClosed Status = "closed"
)

// Account is the core domain entity for a bank account.
// Balance is stored in the smallest currency unit (e.g. cents) to avoid float rounding.
type Account struct {
	ID        *string
	OwnerID   string
	Type      Type
	Currency  string
	Balance   int64
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(ownerID string, accType Type, currency string) *Account {
	return &Account{
		OwnerID:  ownerID,
		Type:     accType,
		Currency: currency,
		Balance:  0,
		Status:   StatusActive,
	}
}
