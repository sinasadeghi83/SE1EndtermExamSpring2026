package transaction

// Repo defines the persistence contract for the Transaction entity.
// Implementations live in the infrastructure layer.
type Repo interface {
	Create(Transaction) (Transaction, error)
	Find(id string) (Transaction, error)
	FindByAccount(accountID string) ([]Transaction, error)
	UpdateStatus(id string, status Status) (Transaction, error)
}
