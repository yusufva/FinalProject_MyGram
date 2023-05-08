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

// GetAllCommentsByPhoto godoc
// @Tags comments
// @Description Get All Comment Data
// @ID get-all-comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "photoId"
// @Success 200 {object} dto.GetCommentResponse
// @Router /comments/photos/{photoId} [get]
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

// GetAllCommentByUser godoc
// @Tags comments
// @Description Get All Comment By User Data
// @ID get-all-comment-by-user
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.GetCommentResponse
// @Router /comments/users [get]
func (ch *commentHandler) GetAllCommentByUser(c *gin.Context) {
	user := c.MustGet("userData").(entity.User)

	allComments, err := ch.commentService.GetAllCommentByUser(user.ID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, allComments)
}

// GetAllCommentById godoc
// @Tags comments
// @Description Get All Comment By Id Data
// @ID get-all-comment-by-id
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param commentId path int true "commentId"
// @Success 200 {object} dto.CommentResponse
// @Router /comments/{commentId} [get]
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

// CreateComment godoc
// @Tags comments
// @Description Create New Comment Data
// @ID create-comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewCommentRequest true "request body json"
// @Success 201 {object} dto.NewCommentResponse
// @Router /comments/photos [post]
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

// UpdateCommentById godoc
// @Tags comments
// @Description Update Comment Data
// @ID update-comment-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewCommentRequest true "request body json"
// @Param commentId path int true "commentId"
// @Success 200 {object} dto.NewCommentResponse
// @Router /comments/{commentId} [put]
func (ch *commentHandler) UpdateCommentById(c *gin.Context) {
	var commentRequest dto.NewCommentRequest

	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	commentId, err := helpers.GetParamsId(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ch.commentService.UpdateCommentById(commentId, user.ID, commentRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

// DeleteCommentById godoc
// @Tags comments
// @Description Delete Comment Data
// @ID delete-comment-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param commentId path int true "commentId"
// @Success 200 {object} dto.NewCommentResponse
// @Router /comments/{commentId} [delete]
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
