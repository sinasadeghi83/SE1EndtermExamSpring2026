package loan

type ApplyRequest struct {
	AccountID     string
	Principal     int64
	InterestRate  float64
	TermMonths    int
	CreditScore   int
	MonthlyIncome int64
	ExistingDebt  int64
}

type Response struct {
	ID           string
	AccountID    string
	Principal    int64
	InterestRate float64
	TermMonths   int
	Status       string
}
