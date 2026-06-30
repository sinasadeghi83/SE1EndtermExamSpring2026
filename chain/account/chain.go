package account

import (
	"redbank/infrastructure/eventbus"
	"redbank/infrastructure/postgres"
	postgresAccount "redbank/infrastructure/postgres/account"
	"redbank/infrastructure/redis"
	accountSvc "redbank/service/account"
)

type Chain struct {
	svc *accountSvc.Service
}

func NewChain() *Chain {
	repo := postgresAccount.NewRepo(postgres.Db)
	publisher := eventbus.NewPublisher(redis.GetClient(0))
	svc := accountSvc.NewService(repo, publisher)
	return &Chain{svc: svc}
}
