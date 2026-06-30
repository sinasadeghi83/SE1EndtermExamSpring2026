package transaction

import (
	"context"
	"errors"

	transactionSvc "redbank/service/transaction"
)

var ErrInvalidRequest = errors.New("chain: invalid request")

func toResponse(d transactionSvc.DTO) *Response {
	return &Response{
		ID:        d.ID,
		AccountID: d.AccountID,
		Type:      d.Type,
		Amount:    d.Amount,
		Currency:  d.Currency,
		Status:    d.Status,
	}
}

func (c *Chain) Post(ctx context.Context, req PostRequest) (*Response, error) {
	if req.AccountID == "" || req.Amount <= 0 {
		return nil, ErrInvalidRequest
	}

	result, err := c.svc.Post(ctx, transactionSvc.PostInput{
		AccountID: req.AccountID,
		Type:      req.Type,
		Amount:    req.Amount,
		Currency:  req.Currency,
		Reference: req.Reference,
	})
	if err != nil {
		return nil, err
	}

	return toResponse(result), nil
}

func (c *Chain) Transfer(ctx context.Context, req TransferRequest) (*Response, *Response, error) {
	if req.FromAccountID == "" || req.ToAccountID == "" || req.Amount <= 0 {
		return nil, nil, ErrInvalidRequest
	}

	from, to, err := c.svc.Transfer(ctx, transactionSvc.TransferInput{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Currency:      req.Currency,
		Reference:     req.Reference,
	})
	if err != nil {
		return nil, nil, err
	}

	return toResponse(from), toResponse(to), nil
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
