package user

import "time"

type Core struct {
	ID          uint
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	Address     string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserDataInterface interface {
	SelectAll() ([]Core, error)
	SelectById(id uint) (Core, error)
	Insert(input Core) error
	Delete(id uint) error
	Login(email string, password string) (dataLogin Core, err error)
}

type UserServiceInterface interface {
	GetAll() ([]Core, error)
	GetById(id uint) (Core, error)
	Create(input Core) error
	Login(email string, password string) (dataLogin Core, token string, err error)
}
