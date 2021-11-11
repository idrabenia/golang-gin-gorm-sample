package main

func ToUserEntity(user *User) *UserEntity {
	return &UserEntity{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func ToUser(entity *UserEntity) *User {
	return &User{
		Id:        int(entity.ID),
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	}
}
