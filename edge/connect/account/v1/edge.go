package v1

import (
	"context"

	"connectrpc.com/connect"
	accountChain "redbank/chain/account"
	accountv1 "redbank/edge/connect/gen/account/v1"
	// @ahum: imports
)

type Edge struct {
	chain *accountChain.Chain
}

func NewEdge() *Edge {
	return &Edge{
		chain: accountChain.NewChain(),
	}
}

func toProto(r *accountChain.Response) *accountv1.Account {
	return &accountv1.Account{
		Id:       r.ID,
		OwnerId:  r.OwnerID,
		Type:     r.Type,
		Currency: r.Currency,
		Balance:  r.Balance,
		Status:   r.Status,
	}
}

func (e *Edge) Health(c context.Context, req *connect.Request[accountv1.HealthRequest]) (*connect.Response[accountv1.HealthResponse], error) {
	res := connect.NewResponse(&accountv1.HealthResponse{
		Message: "UP",
	})

	return res, nil
}

func (e *Edge) OpenAccount(c context.Context, req *connect.Request[accountv1.OpenAccountRequest]) (*connect.Response[accountv1.OpenAccountResponse], error) {
	result, err := e.chain.Open(c, accountChain.OpenRequest{
		OwnerID:  req.Msg.OwnerId,
		Type:     req.Msg.Type,
		Currency: req.Msg.Currency,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&accountv1.OpenAccountResponse{Account: toProto(result)}), nil
}

func (e *Edge) GetAccount(c context.Context, req *connect.Request[accountv1.GetAccountRequest]) (*connect.Response[accountv1.GetAccountResponse], error) {
	result, err := e.chain.Get(c, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&accountv1.GetAccountResponse{Account: toProto(result)}), nil
}

func (e *Edge) ListAccountsByOwner(c context.Context, req *connect.Request[accountv1.ListAccountsByOwnerRequest]) (*connect.Response[accountv1.ListAccountsByOwnerResponse], error) {
	results, err := e.chain.ListByOwner(c, req.Msg.OwnerId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	accounts := make([]*accountv1.Account, 0, len(results))
	for _, r := range results {
		accounts = append(accounts, toProto(r))
	}

	return connect.NewResponse(&accountv1.ListAccountsByOwnerResponse{Accounts: accounts}), nil
}

func (e *Edge) CloseAccount(c context.Context, req *connect.Request[accountv1.CloseAccountRequest]) (*connect.Response[accountv1.CloseAccountResponse], error) {
	if err := e.chain.Close(c, req.Msg.Id); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&accountv1.CloseAccountResponse{Success: true}), nil
}

// @ahum: methods
