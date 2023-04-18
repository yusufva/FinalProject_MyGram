package photo_pg

import (
	"final-project/entity"
	"final-project/pkg/errrs"
	"final-project/repository/photo_repository"

	"gorm.io/gorm"
)

type photoPG struct {
	db *gorm.DB
}

func NewPhotoPg(db *gorm.DB) photo_repository.PhotoRepository {
	return &photoPG{
		db: db,
	}
}

func (p *photoPG) GetAllPhoto() ([]*entity.Photo, errrs.MessageErr) {
	var photos []*entity.Photo
	err := p.db.Find(&photos).Error

	if err != nil {
		return nil, errrs.NewInternalServerError("error while getting photos data")
	}

	return photos, nil
}

func (p *photoPG) GetPhotoById(photoId int) (*entity.Photo, errrs.MessageErr) {
	var photo entity.Photo
	result := p.db.First(&photo, photoId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errrs.NewNotFoundError("photo not found")
		}
		return nil, errrs.NewInternalServerError("something went wrong")
	}

	return &photo, nil
}

func (p *photoPG) GetPhotoByUser(userId int) ([]*entity.Photo, errrs.MessageErr) {
	var photos []*entity.Photo
	err := p.db.Find(&photos, "user_id =?", userId)

	if err != nil {
		return nil, errrs.NewInternalServerError("error while getting photo data")
	}

	return photos, nil
}

func (p *photoPG) CreateNewPhoto(photoPayload *entity.Photo) (*entity.Photo, errrs.MessageErr) {
	result := p.db.Create(photoPayload)

	if result.Error != nil {
		return nil, errrs.NewInternalServerError("something went wrong")
	}

	row := result.Row()

	var photo entity.Photo
	row.Scan(row, &photo)

	return &photo, nil
}

func (p *photoPG) UpdatePhotoById(photoId int, photoPayload *entity.Photo) errrs.MessageErr {
	err := p.db.Model(photoPayload).Updates(entity.Photo{Title: photoPayload.Title, Caption: photoPayload.Caption, Photo_Url: photoPayload.Photo_Url}).Error

	if err != nil {
		return errrs.NewInternalServerError("error while updating photo")
	}

	return nil
}

func (p *photoPG) DeletePhotoById(photoId int) errrs.MessageErr {
	var photo entity.Photo
	result := p.db.First(&photo, photoId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errrs.NewNotFoundError("photo not found")
		}
		return errrs.NewInternalServerError("something went wrong")
	}

	err := p.db.Delete(&photo).Error

	if err != nil {
		return errrs.NewInternalServerError("error while deleting photo")
	}

	return nil
}
