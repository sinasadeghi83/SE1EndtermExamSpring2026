package v1

import (
	"context"

	"connectrpc.com/connect"
	transactionChain "redbank/chain/transaction"
	transactionv1 "redbank/edge/connect/gen/transaction/v1"
	// @ahum: imports
)

type Edge struct {
	chain *transactionChain.Chain
}

func NewEdge() *Edge {
	return &Edge{
		chain: transactionChain.NewChain(),
	}
}

func toProto(r *transactionChain.Response) *transactionv1.Transaction {
	return &transactionv1.Transaction{
		Id:        r.ID,
		AccountId: r.AccountID,
		Type:      r.Type,
		Amount:    r.Amount,
		Currency:  r.Currency,
		Status:    r.Status,
	}
}

func (e *Edge) Health(c context.Context, req *connect.Request[transactionv1.HealthRequest]) (*connect.Response[transactionv1.HealthResponse], error) {
	res := connect.NewResponse(&transactionv1.HealthResponse{
		Message: "UP",
	})

	return res, nil
}

func (e *Edge) PostTransaction(c context.Context, req *connect.Request[transactionv1.PostTransactionRequest]) (*connect.Response[transactionv1.PostTransactionResponse], error) {
	result, err := e.chain.Post(c, transactionChain.PostRequest{
		AccountID: req.Msg.AccountId,
		Type:      req.Msg.Type,
		Amount:    req.Msg.Amount,
		Currency:  req.Msg.Currency,
		Reference: req.Msg.Reference,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&transactionv1.PostTransactionResponse{Transaction: toProto(result)}), nil
}

func (e *Edge) TransferFunds(c context.Context, req *connect.Request[transactionv1.TransferFundsRequest]) (*connect.Response[transactionv1.TransferFundsResponse], error) {
	from, to, err := e.chain.Transfer(c, transactionChain.TransferRequest{
		FromAccountID: req.Msg.FromAccountId,
		ToAccountID:   req.Msg.ToAccountId,
		Amount:        req.Msg.Amount,
		Currency:      req.Msg.Currency,
		Reference:     req.Msg.Reference,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&transactionv1.TransferFundsResponse{
		FromTransaction: toProto(from),
		ToTransaction:   toProto(to),
	}), nil
}

func (e *Edge) GetTransaction(c context.Context, req *connect.Request[transactionv1.GetTransactionRequest]) (*connect.Response[transactionv1.GetTransactionResponse], error) {
	result, err := e.chain.Get(c, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&transactionv1.GetTransactionResponse{Transaction: toProto(result)}), nil
}

func (e *Edge) ListTransactionsByAccount(c context.Context, req *connect.Request[transactionv1.ListTransactionsByAccountRequest]) (*connect.Response[transactionv1.ListTransactionsByAccountResponse], error) {
	results, err := e.chain.ListByAccount(c, req.Msg.AccountId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	transactions := make([]*transactionv1.Transaction, 0, len(results))
	for _, r := range results {
		transactions = append(transactions, toProto(r))
	}

	return connect.NewResponse(&transactionv1.ListTransactionsByAccountResponse{Transactions: transactions}), nil
}

// @ahum: methods
