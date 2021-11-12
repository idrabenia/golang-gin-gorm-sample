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

func ToUserList(entities []*UserEntity) []*User {
	result := make([]*User, 0)

	for _, entity := range entities {
		result = append(result, ToUser(entity))
	}

	return result
}
