package service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errrs"
	"final-project/pkg/helpers"
	"final-project/repository/socialMedia_repository"
	"net/http"
)

type SocialMediaService interface {
	GetAllSocialMedia() (*dto.GetSocialMediaResponse, errrs.MessageErr)
	GetSocialMediaById(socmedId int) (*dto.SocialMediaResponse, errrs.MessageErr)
	CreateSocialMedia(socmedPayload dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errrs.MessageErr)
	UpdateSocialMediaById(socmedId int, socmedPayload dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errrs.MessageErr)
	DeleteSocialMedia(socmedId int) (*dto.NewSocialMediaResponse, errrs.MessageErr)
}

type socialMediaService struct {
	socialMediaRepo socialMedia_repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepo socialMedia_repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{
		socialMediaRepo: socialMediaRepo,
	}
}

func (s *socialMediaService) GetAllSocialMedia() (*dto.GetSocialMediaResponse, errrs.MessageErr) {
	socmeds, err := s.socialMediaRepo.GetAllSocialMedia()

	if err != nil {
		return nil, err
	}

	socmedResponse := []dto.SocialMediaResponse{}

	for _, eachComment := range socmeds {
		socmedResponse = append(socmedResponse, eachComment.EntityToProductResponseDto())
	}

	response := dto.GetSocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "comment data have been sent successfully",
		Data:       socmedResponse,
	}

	return &response, nil
}

func (s *socialMediaService) GetSocialMediaById(socmedId int) (*dto.SocialMediaResponse, errrs.MessageErr) {
	result, err := s.socialMediaRepo.GetSocialMediaById(socmedId)

	if err != nil {
		return nil, err
	}

	response := result.EntityToProductResponseDto()

	return &response, nil
}

func (s *socialMediaService) CreateSocialMedia(socmedPayload dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errrs.MessageErr) {
	err := helpers.ValidateStruct(socmedPayload)

	if err != nil {
		return nil, err
	}

	socmedRequest := &entity.SocialMedia{
		Name:             socmedPayload.Name,
		Social_Media_URL: socmedPayload.Social_Media_URL,
		UserID:           socmedPayload.UserID,
	}

	_, err = s.socialMediaRepo.CreateSocialMedia(socmedRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediaResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "social media has been successfully created",
	}

	return &response, nil
}

func (s *socialMediaService) UpdateSocialMediaById(socmedId int, socmedPayload dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errrs.MessageErr) {
	err := helpers.ValidateStruct(socmedPayload)

	if err != nil {
		return nil, err
	}

	payload := &entity.SocialMedia{
		Name:             socmedPayload.Name,
		Social_Media_URL: socmedPayload.Social_Media_URL,
		UserID:           socmedPayload.UserID,
	}

	err = s.socialMediaRepo.UpdateSocialMedia(socmedId, payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediaResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "photo data successfully updated",
	}

	return &response, nil
}

func (s *socialMediaService) DeleteSocialMedia(socmedId int) (*dto.NewSocialMediaResponse, errrs.MessageErr) {
	err := s.socialMediaRepo.DeleteSocialMedia(socmedId)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediaResponse{
		Result:     "success",
		Message:    "photo has been successfully deleted",
		StatusCode: http.StatusOK,
	}

	return &response, nil
}
