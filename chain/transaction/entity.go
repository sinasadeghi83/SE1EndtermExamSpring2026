package transaction

type PostRequest struct {
	AccountID string
	Type      string
	Amount    int64
	Currency  string
	Reference string
}

type TransferRequest struct {
	FromAccountID string
	ToAccountID   string
	Amount        int64
	Currency      string
	Reference     string
}

type Response struct {
	ID        string
	AccountID string
	Type      string
	Amount    int64
	Currency  string
	Status    string
}
