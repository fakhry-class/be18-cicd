package data

import (
	"errors"
	"fakhry/cleanarch/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (dataLogin user.Core, err error) {
	var data User
	// select * from users where email = ? and password = ?
	tx := repo.db.Where("email = ? and password = ?", email, password).Find(&data)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, errors.New("data not found")
	}
	dataLogin = ModelToCore(data)
	return dataLogin, nil
}

// SelectById implements user.UserDataInterface.
func (repo *userQuery) SelectById(id uint) (user.Core, error) {
	var result User
	tx := repo.db.First(&result, id)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, errors.New("data not found")
	}

	resultCore := ModelToCore(result)
	return resultCore, nil
}

// Delete implements user.UserDataInterface.
func (*userQuery) Delete(id uint) error {
	panic("unimplemented")
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	// mapping dari struct core to struct gorm model
	userGorm := CoreToModel(input)

	// simpan ke DB
	tx := repo.db.Create(&userGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAll implements user.UserDataInterface.
func (repo *userQuery) SelectAll() ([]user.Core, error) {
	var usersData []User
	tx := repo.db.Find(&usersData) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}
	// fmt.Println("users:", usersData)
	//mapping dari struct gorm model ke struct core (entity)
	var usersCore = ListModelToCore(usersData)

	return usersCore, nil
}
