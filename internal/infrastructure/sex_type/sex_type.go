package sex_type

import (
	"context"
	"github.com/uptrace/bun"
	"go-echo/internal/domain/sex_type"
)

type repository struct {
	DB *bun.DB
}

func New(db *bun.DB) sex_type.Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FetchAll() ([]sex_type.Model, error) {
	var stms []sex_type.Model

	if err := r.DB.
		NewSelect().
		Model(&stms).
		Scan(context.Background()); err != nil {
		return nil, err
	}

	return stms, nil
}
