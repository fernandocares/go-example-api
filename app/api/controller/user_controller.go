package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-example-api/app/commons/constant"
	"go-example-api/app/commons/pkg"
	"go-example-api/app/domain/model"
	"go-example-api/app/domain/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService service.UserService
}

func ProvideUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (u *UserController) GetAllUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)
	data := u.UserService.GetAllUser()
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u *UserController) AddUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data user")
	var request model.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	data := u.UserService.AddUserData(request)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u *UserController) GetUserById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data := u.UserService.GetUserById(userID)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u *UserController) UpdateUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update user data by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	var request model.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request body. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	var updatedData = u.UserService.UpdateUserData(userID, request)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, updatedData))
}

func (u *UserController) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	u.UserService.DeleteUser(userID)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}
