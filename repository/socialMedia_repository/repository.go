package socialMedia_repository

import (
	"final-project/entity"
	"final-project/pkg/errrs"
)

type SocialMediaRepository interface {
	GetAllSocialMedia() ([]*entity.SocialMedia, errrs.MessageErr)
	GetSocialMediaById(socmedId int) (*entity.SocialMedia, errrs.MessageErr)
	CreateSocialMedia(socmedPayload *entity.SocialMedia) (*entity.SocialMedia, errrs.MessageErr)
	UpdateSocialMedia(socmedId int, socmedPayload *entity.SocialMedia) errrs.MessageErr
	DeleteSocialMedia(socmedId int) errrs.MessageErr
}
