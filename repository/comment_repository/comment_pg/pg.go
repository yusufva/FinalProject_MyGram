package comment_pg

import (
	"final-project/entity"
	"final-project/pkg/errrs"
	"final-project/repository/comment_repository"

	"gorm.io/gorm"
)

type commentPG struct {
	db *gorm.DB
}

func NewCommentPg(db *gorm.DB) comment_repository.CommentRepository {
	return &commentPG{
		db: db,
	}
}

func (c *commentPG) GetAllCommentByPhoto(photoId int) ([]*entity.Comment, errrs.MessageErr) {
	var comments []*entity.Comment
	err := c.db.Find(&comments, "photo_id =?", photoId)

	if err != nil {
		return nil, errrs.NewInternalServerError("error while getting comments data")
	}

	return comments, nil
}

func (c *commentPG) GetAllCommentByUser(userId int) ([]*entity.Comment, errrs.MessageErr) {
	var comments []*entity.Comment
	err := c.db.Find(&comments, "user_id = ?", userId).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errrs.NewNotFoundError("comment not found")
		}
		return nil, errrs.NewInternalServerError("error while getting comments data")
	}

	return comments, nil
}

func (c *commentPG) GetCommentById(commentId int) (*entity.Comment, errrs.MessageErr) {
	var comment entity.Comment
	result := c.db.First(&comment, commentId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errrs.NewNotFoundError("comment not found")
		}
		return nil, errrs.NewInternalServerError("something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) CreateComment(commentPayload *entity.Comment) (*entity.Comment, errrs.MessageErr) {
	result := c.db.Create(commentPayload)

	if result.Error != nil {
		return nil, errrs.NewInternalServerError("something went wrong")
	}

	row := result.Row()

	var comment entity.Comment
	row.Scan(row, &comment)

	return &comment, nil
}

func (c *commentPG) UpdateCommentById(commentId int, commentPayload *entity.Comment) errrs.MessageErr {
	err := c.db.Model(commentPayload).Updates(entity.Comment{Message: commentPayload.Message}).Error

	if err != nil {
		return errrs.NewInternalServerError("error while updating comment")
	}

	return nil
}

func (c *commentPG) DeleteCommentById(commentId int) errrs.MessageErr {
	var comment *entity.Comment
	err := c.db.First(&comment, commentId).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errrs.NewNotFoundError("comment not found")
		}
		return errrs.NewInternalServerError("error while getting comment data")
	}

	result := c.db.Delete(&comment)

	if result.Error != nil {
		return errrs.NewInternalServerError("error while deleting comment data")
	}

	return nil
}
