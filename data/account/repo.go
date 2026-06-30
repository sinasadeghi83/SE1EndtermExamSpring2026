package account

// Repo defines the persistence contract for the Account entity.
// Implementations live in the infrastructure layer.
type Repo interface {
	Create(Account) (Account, error)
	Find(id string) (Account, error)
	FindByOwner(ownerID string) ([]Account, error)
	Update(Account) (Account, error)
	UpdateBalance(id string, delta int64) (Account, error)
	Delete(id string) error
}
