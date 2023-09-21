package repository

import (
	log "github.com/sirupsen/logrus"
	"go-example-api/app/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func ProvideUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (u *UserRepository) FindAllUser() ([]model.User, error) {
	var users []model.User

	var err = u.db.Joins("Role").Find(&users).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) FindUserById(id int) (model.User, error) {
	user := model.User{
		ID: id,
	}
	err := u.db.Joins("Role").First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepository) Save(user *model.User) (model.User, error) {
	var err = u.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return model.User{}, err
	}
	return *user, nil
}

func (u *UserRepository) DeleteUserById(id int) error {
	err := u.db.Delete(&model.User{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}
