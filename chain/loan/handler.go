package loan

import (
	"context"
	"errors"

	loanSvc "redbank/service/loan"
)

var ErrInvalidRequest = errors.New("chain: invalid request")

func toResponse(d loanSvc.DTO) *Response {
	return &Response{
		ID:           d.ID,
		AccountID:    d.AccountID,
		Principal:    d.Principal,
		InterestRate: d.InterestRate,
		TermMonths:   d.TermMonths,
		Status:       d.Status,
	}
}

func (c *Chain) Apply(ctx context.Context, req ApplyRequest) (*Response, error) {
	if req.AccountID == "" || req.Principal <= 0 {
		return nil, ErrInvalidRequest
	}

	result, err := c.svc.Apply(ctx, loanSvc.ApplyInput{
		AccountID:    req.AccountID,
		Principal:    req.Principal,
		InterestRate: req.InterestRate,
		TermMonths:   req.TermMonths,
		Applicant: loanSvc.Applicant{
			CreditScore:        req.CreditScore,
			MonthlyIncome:      req.MonthlyIncome,
			ExistingDebt:       req.ExistingDebt,
			RequestedPrincipal: req.Principal,
		},
	})
	if err != nil {
		return nil, err
	}

	return toResponse(result), nil
}

func (c *Chain) Evaluate(ctx context.Context, id string) (loanSvc.EvaluationResult, error) {
	if id == "" {
		return loanSvc.EvaluationResult{}, ErrInvalidRequest
	}
	return c.svc.Evaluate(ctx, id)
}

func (c *Chain) Get(ctx context.Context, id string) (*Response, error) {
	if id == "" {
		return nil, ErrInvalidRequest
	}

	result, err := c.svc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return toResponse(result), nil
}

func (c *Chain) ListByAccount(ctx context.Context, accountID string) ([]*Response, error) {
	if accountID == "" {
		return nil, ErrInvalidRequest
	}

	results, err := c.svc.ListByAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}

	responses := make([]*Response, 0, len(results))
	for _, r := range results {
		responses = append(responses, toResponse(r))
	}
	return responses, nil
}
