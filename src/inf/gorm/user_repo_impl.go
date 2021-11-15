package gorm

import (
	"example/hello/src/model"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RepoSet = wire.NewSet(
	NewUserRepo,
	wire.Bind(new(model.UserRepo), new(*UserRepoImpl)),
)

type UserRepoImpl struct {
	Db DbGorm
}

func NewUserRepo(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{Db: db}
}

func (u *UserRepoImpl) FindById(id int) (*model.UserEntity, error) {
	entity := model.UserEntity{}
	result := u.Db.First(&entity, id)

	if result.Error == nil {
		return &entity, nil
	} else {
		return nil, result.Error
	}
}

func (u *UserRepoImpl) FindAll() ([]*model.UserEntity, error) {
	var users []*model.UserEntity

	result := u.Db.Find(&users)

	if result.Error == nil {
		return users, nil
	} else {
		return nil, result.Error
	}
}

func (u *UserRepoImpl) Create(entity *model.UserEntity) (*model.UserEntity, error) {
	result := u.Db.Create(&entity)

	if result.Error == nil {
		return entity, nil
	} else {
		return nil, result.Error
	}
}

func (u *UserRepoImpl) Update(entity *model.UserEntity) (*model.UserEntity, error) {
	if result := u.Db.Updates(entity); result.Error == nil {
		return entity, nil
	} else {
		return nil, result.Error
	}
}

func (u *UserRepoImpl) Delete(id int) error {
	user, err := u.FindById(id)

	if err != nil {
		return err
	}

	if res := u.Db.Delete(user); res.Error != nil {
		return res.Error
	}

	return nil
}
