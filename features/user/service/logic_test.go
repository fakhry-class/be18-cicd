package service

import (
	"errors"
	"fakhry/cleanarch/features/user"
	"fakhry/cleanarch/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// create object data layer
	mockUserDataLayer := new(mocks.UserData)
	returnData := []user.Core{
		{ID: 1, Name: "Budi", Email: "budi@mail.com", Password: "qwerty"},
	}
	t.Run("test case success get all data", func(t *testing.T) {
		mockUserDataLayer.On("SelectAll").Return(returnData, nil).Once()
		//create object service
		srv := New(mockUserDataLayer)
		response, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		mockUserDataLayer.AssertExpectations(t)
	})
	t.Run("test case failed get all data", func(t *testing.T) {
		mockUserDataLayer.On("SelectAll").Return(nil, errors.New("error read data")).Once()
		srv := New(mockUserDataLayer)
		response, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, response)
		mockUserDataLayer.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockUserDataLayer := new(mocks.UserData)
	t.Run("test case success create user", func(t *testing.T) {
		inputData := user.Core{
			Name:        "budi",
			Email:       "budi@mail.com",
			Password:    "qwerty",
			Address:     "Jakarta",
			PhoneNumber: "0812345",
		}
		mockUserDataLayer.On("Insert", inputData).Return(nil).Once()
		srv := New(mockUserDataLayer)
		err := srv.Create(inputData)
		assert.Nil(t, err)
		mockUserDataLayer.AssertExpectations(t)
	})

	t.Run("test case failed create user error validate, empty name", func(t *testing.T) {
		inputData := user.Core{
			Email:       "budi@mail.com",
			Password:    "qwerty",
			Address:     "Jakarta",
			PhoneNumber: "0812345",
		}
		srv := New(mockUserDataLayer)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		mockUserDataLayer.AssertExpectations(t)
	})
}
