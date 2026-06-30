package account

import (
	"redbank/data/account"
	"redbank/data/event"
)

// Service implements the core business rules for the Account domain.
// It depends only on the repository and publisher interfaces defined in data/.
type Service struct {
	repo      account.Repo
	publisher event.Publisher
}

func NewService(repo account.Repo, publisher event.Publisher) *Service {
	return &Service{repo: repo, publisher: publisher}
}
