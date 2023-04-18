package handler

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errrs"
	"final-project/pkg/helpers"
	"final-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) commentHandler {
	return commentHandler{
		commentService: commentService,
	}
}

func (ch *commentHandler) GetAllCommentByPhoto(c *gin.Context) {
	photoId, err := helpers.GetParamsId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	allComments, err := ch.commentService.GetAllCommentByPhoto(photoId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, allComments)
}

func (ch *commentHandler) GetAllCommentByUser(c *gin.Context) {
	userId, err := helpers.GetParamsId(c, "userId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	allComments, err := ch.commentService.GetAllCommentByUser(userId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, allComments)
}

func (ch *commentHandler) GetCommentById(c *gin.Context) {
	commentId, err := helpers.GetParamsId(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ch.commentService.GetCommentById(commentId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, &response)
}

func (ch *commentHandler) CreateComment(c *gin.Context) {
	var commentRequest dto.NewCommentRequest

	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newComment, err := ch.commentService.CreateComment(user.ID, commentRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newComment)
}

func (ch *commentHandler) UpdateCommentById(c *gin.Context) {
	var commentRequest dto.NewCommentRequest

	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	commentId, err := helpers.GetParamsId(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ch.commentService.UpdateCommentById(commentId, commentRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (ch *commentHandler) DeleteCommentById(c *gin.Context) {
	commentId, err := helpers.GetParamsId(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ch.commentService.DeleteCommentById(commentId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
	}

	c.JSON(response.StatusCode, response)
}
