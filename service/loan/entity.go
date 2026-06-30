package loan

type ApplyInput struct {
	AccountID    string
	Principal    int64
	InterestRate float64
	TermMonths   int
	Applicant    Applicant
}

type DTO struct {
	ID           string
	AccountID    string
	Principal    int64
	InterestRate float64
	TermMonths   int
	Status       string
}
