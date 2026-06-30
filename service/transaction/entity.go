package transaction

type PostInput struct {
	AccountID string
	Type      string
	Amount    int64
	Currency  string
	Reference string
}

type TransferInput struct {
	FromAccountID string
	ToAccountID   string
	Amount        int64
	Currency      string
	Reference     string
}

type DTO struct {
	ID        string
	AccountID string
	Type      string
	Amount    int64
	Currency  string
	Status    string
}
