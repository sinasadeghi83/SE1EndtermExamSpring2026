package transactiontest

import (
	"fmt"
	"sync"

	"redbank/data/transaction"
)

// Repo is an in-memory transaction.Repo implementation for unit tests.
type Repo struct {
	mu           sync.Mutex
	seq          int
	transactions map[string]transaction.Transaction
}

func New() *Repo {
	return &Repo{transactions: make(map[string]transaction.Transaction)}
}

func (r *Repo) Create(t transaction.Transaction) (transaction.Transaction, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.seq++
	id := fmt.Sprintf("txn-%d", r.seq)
	t.ID = &id
	r.transactions[id] = t
	return t, nil
}

func (r *Repo) Find(id string) (transaction.Transaction, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.transactions[id]
	if !ok {
		return transaction.Transaction{}, transaction.ErrNotFound
	}
	return t, nil
}

func (r *Repo) FindByAccount(accountID string) ([]transaction.Transaction, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []transaction.Transaction
	for _, t := range r.transactions {
		if t.AccountID == accountID {
			result = append(result, t)
		}
	}
	return result, nil
}

func (r *Repo) UpdateStatus(id string, status transaction.Status) (transaction.Transaction, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.transactions[id]
	if !ok {
		return transaction.Transaction{}, transaction.ErrNotFound
	}
	t.Status = status
	r.transactions[id] = t
	return t, nil
}
