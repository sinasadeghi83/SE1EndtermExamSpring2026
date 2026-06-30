package account

type OpenInput struct {
	OwnerID  string
	Type     string
	Currency string
}

type DTO struct {
	ID       string
	OwnerID  string
	Type     string
	Currency string
	Balance  int64
	Status   string
}
