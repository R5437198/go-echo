package sex_type

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Model struct {
	bun.BaseModel `bun:"table:sex_type"`

	Id        uuid.UUID
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository interface {
	FetchAll() ([]Model, error)
}
