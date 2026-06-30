package loan

import (
	"redbank/infrastructure/eventbus"
	"redbank/infrastructure/postgres"
	postgresLoan "redbank/infrastructure/postgres/loan"
	"redbank/infrastructure/redis"
	loanSvc "redbank/service/loan"
)

type Chain struct {
	svc *loanSvc.Service
}

func NewChain() *Chain {
	repo := postgresLoan.NewRepo(postgres.Db)
	publisher := eventbus.NewPublisher(redis.GetClient(0))
	strategy := loanSvc.NewCompositeStrategy(
		loanSvc.NewCreditScoreStrategy(650),
		loanSvc.NewIncomeBasedStrategy(0.4),
	)
	svc := loanSvc.NewService(repo, publisher, strategy)
	return &Chain{svc: svc}
}
