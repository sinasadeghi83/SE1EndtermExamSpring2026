package accounttest

import (
	"fmt"
	"sync"

	"redbank/data/account"
)

// Repo is an in-memory account.Repo implementation for unit tests.
type Repo struct {
	mu       sync.Mutex
	seq      int
	accounts map[string]account.Account
}

func New() *Repo {
	return &Repo{accounts: make(map[string]account.Account)}
}

func (r *Repo) Create(a account.Account) (account.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.seq++
	id := fmt.Sprintf("acc-%d", r.seq)
	a.ID = &id
	r.accounts[id] = a
	return a, nil
}

func (r *Repo) Find(id string) (account.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	a, ok := r.accounts[id]
	if !ok {
		return account.Account{}, account.ErrNotFound
	}
	return a, nil
}

func (r *Repo) FindByOwner(ownerID string) ([]account.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []account.Account
	for _, a := range r.accounts {
		if a.OwnerID == ownerID {
			result = append(result, a)
		}
	}
	return result, nil
}

func (r *Repo) Update(a account.Account) (account.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if a.ID == nil {
		return account.Account{}, account.ErrNotFound
	}
	if _, ok := r.accounts[*a.ID]; !ok {
		return account.Account{}, account.ErrNotFound
	}
	r.accounts[*a.ID] = a
	return a, nil
}

func (r *Repo) UpdateBalance(id string, delta int64) (account.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	a, ok := r.accounts[id]
	if !ok {
		return account.Account{}, account.ErrNotFound
	}
	a.Balance += delta
	r.accounts[id] = a
	return a, nil
}

func (r *Repo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.accounts[id]; !ok {
		return account.ErrNotFound
	}
	delete(r.accounts, id)
	return nil
}
