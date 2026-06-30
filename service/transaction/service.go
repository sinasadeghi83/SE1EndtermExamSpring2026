package transaction

import (
	accountData "redbank/data/account"
	"redbank/data/event"
	"redbank/data/transaction"
)

// Service implements the core business rules for the Transaction domain.
// It needs the account repo too, since posting a transaction mutates account balances.
type Service struct {
	repo        transaction.Repo
	accountRepo accountData.Repo
	publisher   event.Publisher
}

func NewService(repo transaction.Repo, accountRepo accountData.Repo, publisher event.Publisher) *Service {
	return &Service{repo: repo, accountRepo: accountRepo, publisher: publisher}
}
