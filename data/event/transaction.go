package event

import "time"

const NameTransactionPosted = "transaction.posted"

type TransactionPosted struct {
	TransactionID string
	AccountID     string
	Amount        int64
	At            time.Time
}

func NewTransactionPosted(transactionID, accountID string, amount int64) TransactionPosted {
	return TransactionPosted{TransactionID: transactionID, AccountID: accountID, Amount: amount, At: time.Now()}
}

func (e TransactionPosted) Name() string          { return NameTransactionPosted }
func (e TransactionPosted) OccurredAt() time.Time { return e.At }
