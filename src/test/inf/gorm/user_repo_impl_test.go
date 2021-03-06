package gorm

import (
	gorm2 "example/hello/src/inf/gorm"
	"example/hello/src/test"
	tstmdl "example/hello/src/test/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

func TestFindById(t *testing.T) {
	db := mockDb(new(tstmdl.DbMock))
	service := gorm2.UserRepoImpl{Db: db}

	result, err := service.FindById(1)

	assert.Equal(t, result.ID, uint(1))
	assert.Equal(t, err, nil)
}

func TestCreate(t *testing.T) {
	db := mockDb(new(tstmdl.DbMock))
	service := gorm2.UserRepoImpl{Db: db}
	entity := test.MakeUserEntity(1)

	result, err := service.Create(entity)

	assert.Equal(t, result.ID, uint(1))
	assert.Equal(t, err, nil)
}

func TestFindAll(t *testing.T) {
	db := mockDb(new(tstmdl.DbMock))
	service := gorm2.UserRepoImpl{Db: db}

	result, err := service.FindAll()

	assert.Equal(t, len(result), 0)
	assert.Equal(t, err, nil)
}

func TestDelete(t *testing.T) {
	db := mockDb(new(tstmdl.DbMock))
	service := gorm2.UserRepoImpl{Db: db}

	err := service.Delete(1)

	assert.Equal(t, err, nil)
	db.AssertCalled(t, "Delete", mock.Anything, mock.Anything)
}

func TestUpdate(t *testing.T) {
	db := mockDb(new(tstmdl.DbMock))
	service := gorm2.UserRepoImpl{Db: db}
	entity := test.MakeUserEntity(1)

	result, err := service.Update(entity)

	assert.Equal(t, result.FirstName, "TestFirstName")
	assert.Equal(t, err, nil)
}

func mockDb(dbMock *tstmdl.DbMock) *tstmdl.DbMock {
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
