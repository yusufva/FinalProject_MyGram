package service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errrs"
	"final-project/pkg/helpers"
	"final-project/repository/photo_repository"
	"net/http"
)

type PhotoService interface {
	GetAllPhoto() (*dto.GetPhotoResponse, errrs.MessageErr)
	GetPhotoById(photoId int) (*dto.PhotoResponse, errrs.MessageErr)
	GetPhotoByUser(userId int) (*dto.GetPhotoResponse, errrs.MessageErr)
	CreateNewPhoto(userId int, photoPayload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errrs.MessageErr)
	UpdatePhotoById(photoId int, photoPayload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errrs.MessageErr)
	DeletePhotoById(photoId int) (*dto.NewPhotoResponse, errrs.MessageErr)
}

type photoService struct {
	photoRepo photo_repository.PhotoRepository
}

func NewPhotoService(photoRepo photo_repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepo: photoRepo,
	}
}

func (p *photoService) GetAllPhoto() (*dto.GetPhotoResponse, errrs.MessageErr) {
	photos, err := p.photoRepo.GetAllPhoto()

	if err != nil {
		return nil, err
	}

	photoResponse := []dto.PhotoResponse{}

	for _, eachPhoto := range photos {
		photoResponse = append(photoResponse, eachPhoto.EntityToProductResponseDto())
	}

	response := dto.GetPhotoResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "product data have been sent successfully",
		Data:       photoResponse,
	}

	return &response, nil
}

func (p *photoService) GetPhotoById(photoId int) (*dto.PhotoResponse, errrs.MessageErr) {
	result, err := p.photoRepo.GetPhotoById(photoId)

	if err != nil {
		return nil, err
	}

	response := result.EntityToProductResponseDto()

	return &response, nil
}

func (p *photoService) GetPhotoByUser(userId int) (*dto.GetPhotoResponse, errrs.MessageErr) {
	photos, err := p.photoRepo.GetPhotoByUser(userId)

	if err != nil {
		return nil, err
	}

	photoResponse := []dto.PhotoResponse{}

	for _, eachProduct := range photos {
		photoResponse = append(photoResponse, eachProduct.EntityToProductResponseDto())
	}

	response := dto.GetPhotoResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "photo data have been sent successfully",
		Data:       photoResponse,
	}

	return &response, nil
}

func (p *photoService) CreateNewPhoto(userId int, photoPayload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errrs.MessageErr) {
	err := helpers.ValidateStruct(photoPayload)

	if err != nil {
		return nil, err
	}

	photoRequest := &entity.Photo{
		Title:     photoPayload.Title,
		Caption:   photoPayload.Caption,
		Photo_Url: photoPayload.Photo_Url,
		UserID:    userId,
	}

	_, err = p.photoRepo.CreateNewPhoto(photoRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "photo has been successfully created",
	}

	return &response, nil
}

func (p *photoService) UpdatePhotoById(photoId int, photoPayload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errrs.MessageErr) {
	err := helpers.ValidateStruct(photoPayload)

	if err != nil {
		return nil, err
	}

	payload := &entity.Photo{
		ID:        photoId,
		Title:     photoPayload.Title,
		Caption:   photoPayload.Caption,
		Photo_Url: photoPayload.Photo_Url,
		UserID:    photoPayload.UserID,
	}

	err = p.photoRepo.UpdatePhotoById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "photo data successfully updated",
	}

	return &response, nil
}

func (p *photoService) DeletePhotoById(photoId int) (*dto.NewPhotoResponse, errrs.MessageErr) {
	err := p.photoRepo.DeletePhotoById(photoId)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		Result:     "success",
		Message:    "photo has been successfully deleted",
		StatusCode: http.StatusOK,
	}

	return &response, nil
}
