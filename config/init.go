package config

import (
	"go-example-api/app/api/controller"
	"go-example-api/app/domain/repository"
	"go-example-api/app/domain/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController
}

func ProvideInitialization(userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController) Initialization {
	return Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
	}
}
