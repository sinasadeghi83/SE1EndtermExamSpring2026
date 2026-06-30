package loan

// Repo defines the persistence contract for the Loan entity.
// Implementations live in the infrastructure layer.
type Repo interface {
	Create(Loan) (Loan, error)
	Find(id string) (Loan, error)
	FindByAccount(accountID string) ([]Loan, error)
	Update(Loan) (Loan, error)
	UpdateStatus(id string, status Status) (Loan, error)
	Delete(id string) error
}
