package comment_repository

import (
	"final-project/entity"
	"final-project/pkg/errrs"
)

type CommentRepository interface {
	GetAllCommentByPhoto(photoId int) ([]*entity.Comment, errrs.MessageErr)
	GetAllCommentByUser(userId int) ([]*entity.Comment, errrs.MessageErr)
	GetCommentById(commentId int) (*entity.Comment, errrs.MessageErr)
	CreateComment(commentPayload *entity.Comment) (*entity.Comment, errrs.MessageErr)
	UpdateCommentById(commentId int, commentPayload *entity.Comment) (*entity.Comment, errrs.MessageErr)
	DeleteCommentById(commentId int) errrs.MessageErr
}
