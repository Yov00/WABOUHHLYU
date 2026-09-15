package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) ([]any, error)
}

type svc struct {
}

func NewService() Service {
	return &svc{}
}

func (s *svc) ListProducts(ctx context.Context) ([]any, error) {
	return nil, nil
}
