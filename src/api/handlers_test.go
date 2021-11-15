package api

import (
	"bytes"
	"example/hello/src/model"
	"example/hello/src/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) FindById(id int) (*model.UserEntity, error) {
	m.Called(id)
	return test.MakeUserEntity(1), nil
}

func (m *UserServiceMock) Create(entity *model.UserEntity) (*model.UserEntity, error) {
	m.Called(entity)
	return test.MakeUserEntity(1), nil
}

func (m *UserServiceMock) FindAll() ([]*model.UserEntity, error) {
	m.Called()
	return []*model.UserEntity{test.MakeUserEntity(1)}, nil
}

func (m *UserServiceMock) Update(id int, command *model.UpdateUserCommand) (*model.UserEntity, error) {
	m.Called(id, command)
	return test.MakeUserEntity(1), nil
}

func (m *UserServiceMock) Delete(id int) error {
	m.Called(id)
	return nil
}

func TestFindAllUsers(t *testing.T) {
	context := makeContext()
	service := mockUserService(new(UserServiceMock))

	FindAllUsers(context, service)

	assert.Equal(t, context.Writer.Status(), 200)
	service.AssertCalled(t, "FindAll")
}

func TestProcessFindById(t *testing.T) {
	service := mockUserService(new(UserServiceMock))
	context := makeContext()
	context.Params = mockParams("id", "1")

	GetUser(context, service)

	assert.Equal(t, context.Writer.Status(), 200)
	service.AssertCalled(t, "FindById", 1)
}

func TestCreateUser(t *testing.T) {
	service := mockUserService(new(UserServiceMock))
	context := mockCreateAndUpdate(makeContext())

	CreateUser(context, service)

	assert.Equal(t, context.Writer.Status(), 200)
	service.AssertCalled(t, "Create", mock.Anything)
}

func TestUpdateUser(t *testing.T) {
	service := mockUserService(new(UserServiceMock))
	context := mockCreateAndUpdate(makeContext())
	context.Params = mockParams("id", "1")

	UpdateUser(context, service)

	assert.Equal(t, context.Writer.Status(), 200)
	service.AssertCalled(t, "Update", 1, mock.Anything)
}

func TestDeleteById(t *testing.T) {
	service := mockUserService(new(UserServiceMock))
	context := makeContext()
	context.Params = mockParams("id", "1")

	DeleteUser(context, service)

	service.AssertCalled(t, "Delete", 1)
}

func mockParams(key string, value string) gin.Params {
	return gin.Params{
		gin.Param{
			Key:   key,
			Value: value,
		},
	}
}

func makeContext() *gin.Context {
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	return context
}

func mockCreateAndUpdate(context *gin.Context) *gin.Context {
	context.Request = &http.Request{Header: http.Header{}}
	context.Request.Header.Set("Content-Type", "application/json")

	payload := []byte(`{"FirstName": "First", "LastName": "Last"}`)
	context.Request.Body = io.NopCloser(bytes.NewBuffer(payload))

	return context
}

func mockUserService(service *UserServiceMock) *UserServiceMock {
	service.
		On("FindAll").
		Return([]*model.UserEntity{test.MakeUserEntity(1)})

	service.
		On("FindById", 1).
		Return(test.MakeUserEntity(1), nil)

	service.
		On("Delete", 1).
		Return(test.MakeUserEntity(1), nil)

	service.
		On("Create", mock.Anything).
		Return(test.MakeUserEntity(1), nil)

	service.
		On("Update", 1, mock.Anything).
		Return(test.MakeUserEntity(1), nil)

	return service
}
