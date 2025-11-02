package usecase

import (
	"context"

	"github.com/fpessoa64/desafio_clean_arch/internal/entities"
	"github.com/fpessoa64/desafio_clean_arch/internal/repository"
)

type OrderUsecase struct {
	repo repository.OrderRepository
}

func NewOrderUsecase(repo repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}

func (u *OrderUsecase) Create(ctx context.Context, o *entities.Order) error {
	return u.repo.Create(ctx, o)
}

func (u *OrderUsecase) List(ctx context.Context) ([]entities.Order, error) {
	return u.repo.List(ctx)
}
