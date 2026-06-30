package loantest

import (
	"fmt"
	"sync"

	"redbank/data/loan"
)

// Repo is an in-memory loan.Repo implementation for unit tests.
type Repo struct {
	mu    sync.Mutex
	seq   int
	loans map[string]loan.Loan
}

func New() *Repo {
	return &Repo{loans: make(map[string]loan.Loan)}
}

func (r *Repo) Create(l loan.Loan) (loan.Loan, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.seq++
	id := fmt.Sprintf("loan-%d", r.seq)
	l.ID = &id
	r.loans[id] = l
	return l, nil
}

func (r *Repo) Find(id string) (loan.Loan, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	l, ok := r.loans[id]
	if !ok {
		return loan.Loan{}, loan.ErrNotFound
	}
	return l, nil
}

func (r *Repo) FindByAccount(accountID string) ([]loan.Loan, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []loan.Loan
	for _, l := range r.loans {
		if l.AccountID == accountID {
			result = append(result, l)
		}
	}
	return result, nil
}

func (r *Repo) Update(l loan.Loan) (loan.Loan, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if l.ID == nil {
		return loan.Loan{}, loan.ErrNotFound
	}
	if _, ok := r.loans[*l.ID]; !ok {
		return loan.Loan{}, loan.ErrNotFound
	}
	r.loans[*l.ID] = l
	return l, nil
}

func (r *Repo) UpdateStatus(id string, status loan.Status) (loan.Loan, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	l, ok := r.loans[id]
	if !ok {
		return loan.Loan{}, loan.ErrNotFound
	}
	l.Status = status
	r.loans[id] = l
	return l, nil
}

func (r *Repo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.loans[id]; !ok {
		return loan.ErrNotFound
	}
	delete(r.loans, id)
	return nil
}
