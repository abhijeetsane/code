package db

import (
	"context"
	"database/sql"

	"shopsvc/internal/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(conn *sql.DB) *Store {
	return &Store{db: conn}
}

func (s *Store) Save(ctx context.Context, o *models.Order) error {
	// INSERT ...
	return nil
}

func (s *Store) FindByID(ctx context.Context, id string) (*models.Order, error) {
	// SELECT ...
	return nil, nil
}

func (s *Store) ListByCustomer(ctx context.Context, customerID string) ([]*models.Order, error) {
	// SELECT ...
	return nil, nil
}

func (s *Store) Delete(ctx context.Context, id string) error {
	// DELETE ...
	return nil
}

func (s *Store) UpdateStatus(ctx context.Context, id, status string) error {
	// UPDATE ...
	return nil
}
