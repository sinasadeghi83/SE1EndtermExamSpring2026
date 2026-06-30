package v1

import (
	"context"

	"connectrpc.com/connect"
	loanChain "redbank/chain/loan"
	loanv1 "redbank/edge/connect/gen/loan/v1"
	// @ahum: imports
)

type Edge struct {
	chain *loanChain.Chain
}

func NewEdge() *Edge {
	return &Edge{
		chain: loanChain.NewChain(),
	}
}

func toProto(r *loanChain.Response) *loanv1.Loan {
	return &loanv1.Loan{
		Id:           r.ID,
		AccountId:    r.AccountID,
		Principal:    r.Principal,
		InterestRate: r.InterestRate,
		TermMonths:   int32(r.TermMonths),
		Status:       r.Status,
	}
}

func (e *Edge) Health(c context.Context, req *connect.Request[loanv1.HealthRequest]) (*connect.Response[loanv1.HealthResponse], error) {
	res := connect.NewResponse(&loanv1.HealthResponse{
		Message: "UP",
	})

	return res, nil
}

func (e *Edge) ApplyLoan(c context.Context, req *connect.Request[loanv1.ApplyLoanRequest]) (*connect.Response[loanv1.ApplyLoanResponse], error) {
	result, err := e.chain.Apply(c, loanChain.ApplyRequest{
		AccountID:     req.Msg.AccountId,
		Principal:     req.Msg.Principal,
		InterestRate:  req.Msg.InterestRate,
		TermMonths:    int(req.Msg.TermMonths),
		CreditScore:   int(req.Msg.CreditScore),
		MonthlyIncome: req.Msg.MonthlyIncome,
		ExistingDebt:  req.Msg.ExistingDebt,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&loanv1.ApplyLoanResponse{Loan: toProto(result)}), nil
}

func (e *Edge) EvaluateLoan(c context.Context, req *connect.Request[loanv1.EvaluateLoanRequest]) (*connect.Response[loanv1.EvaluateLoanResponse], error) {
	result, err := e.chain.Evaluate(c, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&loanv1.EvaluateLoanResponse{
		Approved: result.Approved,
		Score:    result.Score,
		Reason:   result.Reason,
	}), nil
}

func (e *Edge) GetLoan(c context.Context, req *connect.Request[loanv1.GetLoanRequest]) (*connect.Response[loanv1.GetLoanResponse], error) {
	result, err := e.chain.Get(c, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&loanv1.GetLoanResponse{Loan: toProto(result)}), nil
}

func (e *Edge) ListLoansByAccount(c context.Context, req *connect.Request[loanv1.ListLoansByAccountRequest]) (*connect.Response[loanv1.ListLoansByAccountResponse], error) {
	results, err := e.chain.ListByAccount(c, req.Msg.AccountId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	loans := make([]*loanv1.Loan, 0, len(results))
	for _, r := range results {
		loans = append(loans, toProto(r))
	}

	return connect.NewResponse(&loanv1.ListLoansByAccountResponse{Loans: loans}), nil
}

// @ahum: methods
