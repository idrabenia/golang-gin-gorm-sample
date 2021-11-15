package api

import (
	"bytes"
	"example/hello/src/model"
	"example/hello/src/test/api"
	"example/hello/src/test/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindAllUsers(t *testing.T) {
	context := makeContext()
	service := mockUserService(new(services.UserServiceMock))

	FindAllUsers(context, service)

	assert.Equal(t, context.Writer.Status(), 200)
	service.AssertCalled(t, "FindAll")
}

func TestProcessFindById(t *testing.T) {
	service := mockUserService(new(services.UserServiceMock))
	context := makeContext()
	context.Params = mockParams("id", "1")

	GetUser(context, service)

	assert.Equal(t, context.Writer.Status(), 200)
	service.AssertCalled(t, "FindById", 1)
}

func TestCreateUser(t *testing.T) {
	service := mockUserService(new(services.UserServiceMock))
	context := mockCreateAndUpdate(makeContext())

	CreateUser(context, service)

	assert.Equal(t, context.Writer.Status(), 200)
	service.AssertCalled(t, "Create", mock.Anything)
}

func TestUpdateUser(t *testing.T) {
	service := mockUserService(new(services.UserServiceMock))
	context := mockCreateAndUpdate(makeContext())
	context.Params = mockParams("id", "1")

	UpdateUser(context, service)

	assert.Equal(t, context.Writer.Status(), 200)
	service.AssertCalled(t, "Update", 1, mock.Anything)
}

func TestDeleteById(t *testing.T) {
	service := mockUserService(new(services.UserServiceMock))
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

func mockUserService(service *services.UserServiceMock) *services.UserServiceMock {
	service.
		On("FindAll").
		Return([]*model.UserEntity{api.MakeUserEntity(1)})

	service.
		On("FindById", 1).
		Return(api.MakeUserEntity(1), nil)

	service.
		On("Delete", 1).
		Return(api.MakeUserEntity(1), nil)

	service.
		On("Create", mock.Anything).
		Return(api.MakeUserEntity(1), nil)

	service.
		On("Update", 1, mock.Anything).
		Return(api.MakeUserEntity(1), nil)

	return service
}
