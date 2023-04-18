package service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errrs"
	"final-project/pkg/helpers"
	"final-project/repository/comment_repository"
	"net/http"
)

type CommentService interface {
	GetAllCommentByPhoto(photoId int) (*dto.GetCommentResponse, errrs.MessageErr)
	GetAllCommentByUser(userId int) (*dto.GetCommentResponse, errrs.MessageErr)
	GetCommentById(commentId int) (*dto.CommentResponse, errrs.MessageErr)
	CreateComment(commentPayload dto.NewCommentRequest) (*dto.NewCommentResponse, errrs.MessageErr)
	UpdateCommentById(commentId int, CommentPayload dto.NewCommentRequest) (*dto.NewCommentResponse, errrs.MessageErr)
	DeleteCommentById(commentId int) (*dto.NewCommentResponse, errrs.MessageErr)
}

type commentService struct {
	commentRepo comment_repository.CommentRepository
}

func NewCommentService(commentRepo comment_repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (c *commentService) GetAllCommentByPhoto(photoId int) (*dto.GetCommentResponse, errrs.MessageErr) {
	comments, err := c.commentRepo.GetAllCommentByPhoto(photoId)

	if err != nil {
		return nil, err
	}

	commentResponse := []dto.CommentResponse{}

	for _, eachComment := range comments {
		commentResponse = append(commentResponse, eachComment.EntityToProductResponseDto())
	}

	response := dto.GetCommentResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "comment data have been sent successfully",
		Data:       commentResponse,
	}

	return &response, nil
}

func (c *commentService) GetAllCommentByUser(userId int) (*dto.GetCommentResponse, errrs.MessageErr) {
	comments, err := c.commentRepo.GetAllCommentByUser(userId)

	if err != nil {
		return nil, err
	}

	commentResponse := []dto.CommentResponse{}

	for _, eachComment := range comments {
		commentResponse = append(commentResponse, eachComment.EntityToProductResponseDto())
	}

	response := dto.GetCommentResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "comment data have been sent successfully",
		Data:       commentResponse,
	}

	return &response, nil
}

func (c *commentService) GetCommentById(commenId int) (*dto.CommentResponse, errrs.MessageErr) {
	result, err := c.commentRepo.GetCommentById(commenId)

	if err != nil {
		return nil, err
	}

	response := result.EntityToProductResponseDto()

	return &response, nil
}

func (c *commentService) CreateComment(commentPayload dto.NewCommentRequest) (*dto.NewCommentResponse, errrs.MessageErr) {
	err := helpers.ValidateStruct(commentPayload)

	if err != nil {
		return nil, err
	}

	commentRequest := &entity.Comment{
		UserID:  commentPayload.UserID,
		PhotoID: commentPayload.PhotoID,
		Message: commentPayload.Message,
	}

	_, err = c.commentRepo.CreateComment(commentRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "comment has been successfully created",
	}

	return &response, nil
}

func (c *commentService) UpdateCommentById(commentId int, commentPayload dto.NewCommentRequest) (*dto.NewCommentResponse, errrs.MessageErr) {
	err := helpers.ValidateStruct(commentPayload)

	if err != nil {
		return nil, err
	}

	payload := &entity.Comment{
		UserID:  commentPayload.UserID,
		PhotoID: commentPayload.PhotoID,
		Message: commentPayload.Message,
	}

	err = c.commentRepo.UpdateCommentById(commentId, payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "comment data successfully updated",
	}

	return &response, nil
}

func (c *commentService) DeleteCommentById(photoId int) (*dto.NewCommentResponse, errrs.MessageErr) {
	err := c.commentRepo.DeleteCommentById(photoId)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		Result:     "success",
		Message:    "comment has been successfully deleted",
		StatusCode: http.StatusOK,
	}

	return &response, nil
}
