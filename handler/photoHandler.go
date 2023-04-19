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

type photoHandler struct {
	photoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) photoHandler {
	return photoHandler{
		photoService: photoService,
	}
}

func (ph *photoHandler) GetAllPhoto(c *gin.Context) {
	allPhotos, err := ph.photoService.GetAllPhoto()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, allPhotos)
}

func (ph *photoHandler) GetPhotoById(c *gin.Context) {
	photoId, err := helpers.GetParamsId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ph.photoService.GetPhotoById(photoId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, &response)
}

func (ph *photoHandler) GetPhotoByUser(c *gin.Context) {
	userId, err := helpers.GetParamsId(c, "userId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	allPhotos, err := ph.photoService.GetPhotoByUser(userId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, allPhotos)
}

// CreateNewPhoto godoc
// @Tags photos
// @Description Create New Photo Data
// @ID create-new-photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewPhotoRequest true "request body json"
// @Success 201 {object} dto.NewPhotoRequest
// @Router /photos [post]
func (ph *photoHandler) CreateNewPhoto(c *gin.Context) {
	var photoRequest dto.NewPhotoRequest

	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newPhoto, err := ph.photoService.CreateNewPhoto(user.ID, photoRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newPhoto)
}

func (ph *photoHandler) UpdatePhotoById(c *gin.Context) {
	var photoRequest dto.NewPhotoRequest

	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	photoId, err := helpers.GetParamsId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ph.photoService.UpdatePhotoById(photoId, photoRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (ph *photoHandler) DeletePhotoById(c *gin.Context) {
	photoId, err := helpers.GetParamsId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := ph.photoService.DeletePhotoById(photoId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}
