package event

import "time"

const (
	NameAccountCreated = "account.created"
	NameAccountClosed  = "account.closed"
)

type AccountCreated struct {
	AccountID string
	OwnerID   string
	At        time.Time
}

func NewAccountCreated(accountID, ownerID string) AccountCreated {
	return AccountCreated{AccountID: accountID, OwnerID: ownerID, At: time.Now()}
}

func (e AccountCreated) Name() string          { return NameAccountCreated }
func (e AccountCreated) OccurredAt() time.Time { return e.At }

type AccountClosed struct {
	AccountID string
	At        time.Time
}

func NewAccountClosed(accountID string) AccountClosed {
	return AccountClosed{AccountID: accountID, At: time.Now()}
}

func (e AccountClosed) Name() string          { return NameAccountClosed }
func (e AccountClosed) OccurredAt() time.Time { return e.At }
