package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"time"

	"github.com/fpessoa64/desafio_clean_arch/internal/entities"
)

type OrderRepositorySqlite struct {
	db *sql.DB
}

func NewOrderRepositorySqlite(db *sql.DB) *OrderRepositorySqlite {
	return &OrderRepositorySqlite{db: db}
}

func (r *OrderRepositorySqlite) Create(ctx context.Context, o *entities.Order) error {
	query := `INSERT INTO orders (name, amount, status, created_at) VALUES (?, ?, ?, CURRENT_TIMESTAMP)`
	result, err := r.db.ExecContext(ctx, query, o.Name, o.Amount, o.Status)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	o.ID = id
	o.CreatedAt = time.Now()
	return nil
}

func (r *OrderRepositorySqlite) List(ctx context.Context) ([]entities.Order, error) {
	query := `SELECT id, name, amount, status, created_at FROM orders ORDER BY id DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := make([]entities.Order, 0)
	for rows.Next() {
		var o entities.Order
		if err := rows.Scan(&o.ID, &o.Name, &o.Amount, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	fmt.Println(orders)
	return orders, nil
}
