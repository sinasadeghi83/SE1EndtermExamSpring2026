package account

import (
	"context"
	"errors"

	accountSvc "redbank/service/account"
)

var ErrInvalidRequest = errors.New("chain: invalid request")

func toResponse(d accountSvc.DTO) *Response {
	return &Response{
		ID:       d.ID,
		OwnerID:  d.OwnerID,
		Type:     d.Type,
		Currency: d.Currency,
		Balance:  d.Balance,
		Status:   d.Status,
	}
}

func (c *Chain) Open(ctx context.Context, req OpenRequest) (*Response, error) {
	if req.OwnerID == "" || req.Currency == "" {
		return nil, ErrInvalidRequest
	}

	result, err := c.svc.Open(ctx, accountSvc.OpenInput{
		OwnerID:  req.OwnerID,
		Type:     req.Type,
		Currency: req.Currency,
	})
	if err != nil {
		return nil, err
	}

	return toResponse(result), nil
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

func (c *Chain) ListByOwner(ctx context.Context, ownerID string) ([]*Response, error) {
	if ownerID == "" {
		return nil, ErrInvalidRequest
	}

	results, err := c.svc.ListByOwner(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	responses := make([]*Response, 0, len(results))
	for _, r := range results {
		responses = append(responses, toResponse(r))
	}
	return responses, nil
}

func (c *Chain) Close(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidRequest
	}
	return c.svc.Close(ctx, id)
}
