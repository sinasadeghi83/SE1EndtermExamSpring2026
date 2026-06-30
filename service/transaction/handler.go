package transaction

import "context"

func (s *Service) Post(ctx context.Context, input PostInput) (DTO, error) {
	return DTO{}, nil
}

func (s *Service) Transfer(ctx context.Context, input TransferInput) (DTO, DTO, error) {
	return DTO{}, DTO{}, nil
}

func (s *Service) Get(ctx context.Context, id string) (DTO, error) {
	return DTO{}, nil
}

func (s *Service) ListByAccount(ctx context.Context, accountID string) ([]DTO, error) {
	return nil, nil
}
