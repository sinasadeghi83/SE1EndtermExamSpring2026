package transaction

import (
	"redbank/infrastructure/eventbus"
	"redbank/infrastructure/postgres"
	postgresAccount "redbank/infrastructure/postgres/account"
	postgresTransaction "redbank/infrastructure/postgres/transaction"
	"redbank/infrastructure/redis"
	transactionSvc "redbank/service/transaction"
)

type Chain struct {
	svc *transactionSvc.Service
}

func NewChain() *Chain {
	repo := postgresTransaction.NewRepo(postgres.Db)
	accountRepo := postgresAccount.NewRepo(postgres.Db)
	publisher := eventbus.NewPublisher(redis.GetClient(0))
	svc := transactionSvc.NewService(repo, accountRepo, publisher)
	return &Chain{svc: svc}
}
