package service

import (
	"errors"
	"fakhry/cleanarch/app/middlewares"
	"fakhry/cleanarch/features/user"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (dataLogin user.Core, token string, err error) {
	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return user.Core{}, "", err
	}
	token, err = middlewares.CreateToken(int(dataLogin.ID))
	if err != nil {
		return user.Core{}, "", err
	}
	return dataLogin, token, nil
}

// GetById implements user.UserServiceInterface.
func (service *userService) GetById(id uint) (user.Core, error) {
	return service.userData.SelectById(id)
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	// if input.Name == "" || input.Email == "" || input.Password == "" {
	// 	return errors.New("validation error. name/email/password required")
	// }
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validation error" + errValidate.Error())
	}
	err := service.userData.Insert(input)
	return err
}

// GetAll implements user.UserServiceInterface.
func (service *userService) GetAll() ([]user.Core, error) {
	result, err := service.userData.SelectAll()
	return result, err
}
