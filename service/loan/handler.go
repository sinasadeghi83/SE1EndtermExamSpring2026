package loan

import "context"

func (s *Service) Apply(ctx context.Context, input ApplyInput) (DTO, error) {
	return DTO{}, nil
}

func (s *Service) Evaluate(ctx context.Context, id string) (EvaluationResult, error) {
	return EvaluationResult{}, nil
}

func (s *Service) Get(ctx context.Context, id string) (DTO, error) {
	return DTO{}, nil
}

func (s *Service) ListByAccount(ctx context.Context, accountID string) ([]DTO, error) {
	return nil, nil
}
