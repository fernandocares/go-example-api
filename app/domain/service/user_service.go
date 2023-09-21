package service

import (
	log "github.com/sirupsen/logrus"
	"go-example-api/app/commons/constant"
	pkg "go-example-api/app/commons/pkg"
	"go-example-api/app/domain/model"
	"go-example-api/app/domain/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func ProvideUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		UserRepository: userRepository,
	}
}

func (u UserService) UpdateUserData(userId int, request model.User) model.User {
	data, err := u.UserRepository.FindUserById(userId)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.RoleID = request.RoleID
	data.Email = request.Email
	data.Name = request.Password
	data.Status = request.Status
	_, err = u.UserRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	return data
}

func (u UserService) GetUserById(userId int) model.User {
	data, err := u.UserRepository.FindUserById(userId)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	return data
}

func (u UserService) AddUserData(user model.User) model.User {
	data, err := u.UserRepository.Save(&user)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	return data
}

func (u UserService) GetAllUser() []model.User {
	log.Info("start to execute get all data user")

	data, err := u.UserRepository.FindAllUser()
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	return data
}

func (u UserService) DeleteUser(userId int) {
	err := u.UserRepository.DeleteUserById(userId)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}
}
