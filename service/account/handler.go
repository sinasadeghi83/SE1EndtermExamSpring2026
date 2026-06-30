package account

import "context"

func (s *Service) Open(ctx context.Context, input OpenInput) (DTO, error) {
	return DTO{}, nil
}

func (s *Service) Get(ctx context.Context, id string) (DTO, error) {
	return DTO{}, nil
}

func (s *Service) ListByOwner(ctx context.Context, ownerID string) ([]DTO, error) {
	return nil, nil
}

func (s *Service) Close(ctx context.Context, id string) error {
	return nil
}

func (s *Service) AdjustBalance(ctx context.Context, id string, delta int64) (DTO, error) {
	return DTO{}, nil
}
