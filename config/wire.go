//go:generate go run github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"go-example-api/app/api/controller"
	"go-example-api/app/domain/repository"
	"go-example-api/app/domain/service"
	"gorm.io/gorm"
)

func InitAPI(db *gorm.DB) Initialization {
	wire.Build(
		repository.ProvideUserRepository,
		service.ProvideUserService,
		controller.ProvideUserController,
		ProvideInitialization,
	)
	return Initialization{}
}
