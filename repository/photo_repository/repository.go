package photo_repository

import (
	"final-project/entity"
	"final-project/pkg/errrs"
)

type PhotoRepository interface {
	GetAllPhoto() ([]*entity.Photo, errrs.MessageErr)
	GetPhotoById(photoId int) (*entity.Photo, errrs.MessageErr)
	GetPhotoByUser(userId int) ([]*entity.Photo, errrs.MessageErr)
	CreateNewPhoto(photoPayload *entity.Photo) (*entity.Photo, errrs.MessageErr)
	UpdatePhotoById(photoPayload *entity.Photo) errrs.MessageErr
	DeletePhotoById(photoId int) errrs.MessageErr
}
