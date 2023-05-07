package postgres

import (
	"database/sql"
	"github.com/uptrace/bun/driver/pgdriver"
	"go-echo/internal/config"
)

type BunOrm struct {
	Orm *sql.DB
}

func New() (*sql.DB, error) {
	c, err := config.New()
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.DbUrl())))

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
