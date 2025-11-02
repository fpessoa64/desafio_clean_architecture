package repository

import (
	"context"

	"github.com/fpessoa64/desafio_clean_arch/internal/entities"
)

type OrderRepository interface {
	Create(ctx context.Context, o *entities.Order) error
	List(ctx context.Context) ([]entities.Order, error)
}
