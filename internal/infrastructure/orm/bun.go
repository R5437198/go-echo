package orm

import (
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"go-echo/internal/infrastructure/db/postgres"
)

type BunOrm struct {
	DB *bun.DB
}

func New() *BunOrm {
	db, err := postgres.New()
	if err != nil {
		return nil
	}

	orm := bun.NewDB(db, pgdialect.New())

	return &BunOrm{DB: orm}
}
