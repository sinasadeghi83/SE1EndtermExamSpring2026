package transaction

import "time"

type Type string

const (
	TypeDeposit     Type = "deposit"
	TypeWithdrawal  Type = "withdrawal"
	TypeTransferIn  Type = "transfer_in"
	TypeTransferOut Type = "transfer_out"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
	StatusReversed  Status = "reversed"
)

// Transaction is the core domain entity for a ledger movement on an Account.
type Transaction struct {
	ID                    *string
	AccountID             string
	CounterpartyAccountID *string
	Type                  Type
	Amount                int64
	Currency              string
	Status                Status
	Reference             string
	CreatedAt             time.Time
}

func New(accountID string, txType Type, amount int64, currency string) *Transaction {
	return &Transaction{
		AccountID: accountID,
		Type:      txType,
		Amount:    amount,
		Currency:  currency,
		Status:    StatusPending,
	}
}
