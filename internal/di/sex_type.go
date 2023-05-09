package di

import (
	"github.com/uptrace/bun"
	sexTypeApplication "go-echo/internal/application/sex_type"
	sexTypeDomain "go-echo/internal/domain/sex_type"
	"go-echo/internal/infrastructure/orm"
	sexTypeRepository "go-echo/internal/infrastructure/sex_type"
	"go-echo/internal/presentation/handler/master_data/sex_type"
)

func dbDi() *bun.DB {
	db := orm.New()
	return db.Orm
}

func sexTypeRepositoryDi() sexTypeDomain.Repository {
	db := dbDi()
	return sexTypeRepository.New(db)
}

func sexTypeApplicationDi() sexTypeApplication.Application {
	rep := sexTypeRepositoryDi()
	return sexTypeApplication.New(rep)
}

func SexTypeHandler() *sex_type.Handler {
	app := sexTypeApplicationDi()
	return sex_type.New(app)
}
