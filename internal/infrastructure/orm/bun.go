package orm

import (
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"go-echo/internal/infrastructure/db/postgres"
)

func New() (*bun.DB, error) {
	db, err := postgres.New()
	return bun.NewDB(db, pgdialect.New()), err
}
