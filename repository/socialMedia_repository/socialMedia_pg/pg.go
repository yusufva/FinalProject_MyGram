package socialMedia_pg

import (
	"final-project/entity"
	"final-project/pkg/errrs"
	"final-project/repository/socialMedia_repository"

	"gorm.io/gorm"
)

type socialMediaPG struct {
	db *gorm.DB
}

func NewSocialMediaPg(db *gorm.DB) socialMedia_repository.SocialMediaRepository {
	return &socialMediaPG{
		db: db,
	}
}

func (s *socialMediaPG) GetAllSocialMedia() ([]*entity.SocialMedia, errrs.MessageErr) {
	var socmeds []*entity.SocialMedia
	err := s.db.Find(&socmeds).Error

	if err != nil {
		return nil, errrs.NewInternalServerError("error while getting social media data")
	}

	return socmeds, nil
}

func (s *socialMediaPG) GetSocialMediaById(socmedId int) (*entity.SocialMedia, errrs.MessageErr) {
	var socmed entity.SocialMedia
	err := s.db.First(&socmed, socmedId).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errrs.NewNotFoundError("social media not found")
		}
		return nil, errrs.NewInternalServerError("error while getting social media data")
	}

	return &socmed, nil
}

func (s *socialMediaPG) CreateSocialMedia(socmedPayload *entity.SocialMedia) (*entity.SocialMedia, errrs.MessageErr) {
	result := s.db.Create(&socmedPayload)
	if result.Error != nil {
		return nil, errrs.NewInternalServerError("error while creating social media data")
	}

	row := result.Row()
	var socmed entity.SocialMedia
	row.Scan(row, &socmed)

	return &socmed, nil
}

func (s *socialMediaPG) UpdateSocialMedia(socmedId int, socmedPayload *entity.SocialMedia) errrs.MessageErr {
	err := s.db.Model(socmedPayload).Updates(entity.SocialMedia{Name: socmedPayload.Name, Social_Media_URL: socmedPayload.Social_Media_URL}).Error

	if err != nil {
		return errrs.NewInternalServerError("error while updating social media data")
	}

	return nil
}

func (s *socialMediaPG) DeleteSocialMedia(socmedId int) errrs.MessageErr {
	var socmed entity.SocialMedia
	err := s.db.First(&socmed, socmedId).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errrs.NewNotFoundError("social media not found")
		}
		return errrs.NewInternalServerError("error while getting social media data")
	}

	result := s.db.Delete(&socmed)

	if result.Error != nil {
		return errrs.NewInternalServerError("error while deleting data")
	}
	return nil
}
