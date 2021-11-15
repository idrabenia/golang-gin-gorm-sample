package model

type UserRepo interface {
	FindById(id int) (*UserEntity, error)

	FindAll() ([]*UserEntity, error)

	Create(entity *UserEntity) (*UserEntity, error)

	Update(entity *UserEntity) (*UserEntity, error)

	Delete(id int) error
}
