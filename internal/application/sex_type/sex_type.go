package sex_type

import (
	"go-echo/internal/domain/sex_type"
)

type application struct {
	sexTypeRep sex_type.Repository
}

type Application interface {
	FetchAll() ([]sex_type.Model, error)
}

func New(sexTypeRep sex_type.Repository) Application {
	return &application{sexTypeRep: sexTypeRep}
}

func (a *application) FetchAll() ([]sex_type.Model, error) {
	sts, err := a.sexTypeRep.FetchAll()
	if err != nil {
		return nil, err
	}

	return sts, nil
}
