package account

type OpenRequest struct {
	OwnerID  string
	Type     string
	Currency string
}

type Response struct {
	ID       string
	OwnerID  string
	Type     string
	Currency string
	Balance  int64
	Status   string
}
