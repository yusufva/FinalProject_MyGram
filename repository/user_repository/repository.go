package user_repository

import (
	"final-project/entity"
	"final-project/pkg/errrs"
)

type UserRepository interface {
	CreateNewUser(user entity.User) errrs.MessageErr
	GetUserById(userId int) (*entity.User, errrs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errrs.MessageErr)
}
