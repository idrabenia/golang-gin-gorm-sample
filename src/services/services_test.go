package services

import (
	"example/hello/src/model"
	"example/hello/src/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

type DbMock struct {
	mock.Mock
}

func (db *DbMock) Create(value interface{}) (tx *gorm.DB) {
	db.Called(value)
	return &gorm.DB{Error: nil}
}

func (db *DbMock) Updates(values interface{}) (tx *gorm.DB) {
	db.Called(values)
	return &gorm.DB{Error: nil}
}

func (db *DbMock) Delete(value interface{}, conds ...interface{}) (tx *gorm.DB) {
	db.Called(value, conds)
	return &gorm.DB{Error: nil}
}

func (db *DbMock) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	db.Called(dest, conds)
	return &gorm.DB{Error: nil}
}

func (db *DbMock) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	db.Called(dest, conds)

	item := dest.(*model.UserEntity)
	item.ID = 1

	return &gorm.DB{Error: nil}
}

func TestFindById(t *testing.T) {
	db := mockDb(new(DbMock))
	service := UserService{Db: db}

	result, err := service.FindById(1)

	assert.Equal(t, result.ID, uint(1))
	assert.Equal(t, err, nil)
}

func TestCreate(t *testing.T) {
	db := mockDb(new(DbMock))
	service := UserService{Db: db}
	entity := test.MakeUserEntity(1)

	result, err := service.Create(entity)

	assert.Equal(t, result.ID, uint(1))
	assert.Equal(t, err, nil)
}

func TestFindAll(t *testing.T) {
	db := mockDb(new(DbMock))
	service := UserService{Db: db}

	result, err := service.FindAll()

	assert.Equal(t, len(result), 0)
	assert.Equal(t, err, nil)
}

func TestDelete(t *testing.T) {
	db := mockDb(new(DbMock))
	service := UserService{Db: db}

	err := service.Delete(1)

	assert.Equal(t, err, nil)
	db.AssertCalled(t, "Delete", mock.Anything, mock.Anything)
}

func TestUpdate(t *testing.T) {
	db := mockDb(new(DbMock))
	service := UserService{Db: db}
	command := MakeUpdateCommand()

	result, err := service.Update(1, command)

	assert.Equal(t, result.FirstName, "TestFirstName")
	assert.Equal(t, err, nil)
}

func mockDb(dbMock *DbMock) *DbMock {
	dbMock.
		On("First", mock.Anything, mock.Anything).
		Return(&gorm.DB{Error: nil})

	dbMock.
		On("Create", mock.Anything).
		Return(&gorm.DB{Error: nil})

	dbMock.
		On("Find", mock.Anything, mock.Anything).
		Return(&gorm.DB{Error: nil})

	dbMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(&gorm.DB{Error: nil})

	dbMock.
		On("Updates", mock.Anything).
		Return(&gorm.DB{Error: nil})

	return dbMock
}

func MakeUpdateCommand() *model.UpdateUserCommand {
	return &model.UpdateUserCommand{
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
	}
}
