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

type socialMediaHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) socialMediaHandler {
	return socialMediaHandler{
		socialMediaService: socialMediaService,
	}
}

func (sh *socialMediaHandler) GetAllSocialMedia(c *gin.Context) {
	allSocmed, err := sh.socialMediaService.GetAllSocialMedia()

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(allSocmed.StatusCode, allSocmed)
}

func (sh *socialMediaHandler) GetSocialMediaById(c *gin.Context) {
	socmedId, err := helpers.GetParamsId(c, "socmedId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := sh.socialMediaService.GetSocialMediaById(socmedId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, &response)
}

func (sh *socialMediaHandler) CreateSocialMedia(c *gin.Context) {
	var socmedRequest dto.NewSocialMediaRequest

	if err := c.ShouldBindJSON(&socmedRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newSocmed, err := sh.socialMediaService.CreateSocialMedia(user.ID, socmedRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newSocmed)
}

func (sh *socialMediaHandler) UpdateSocialMediaById(c *gin.Context) {

	var socmedRequest dto.NewSocialMediaRequest

	if err := c.ShouldBindJSON(&socmedRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	socmedId, err := helpers.GetParamsId(c, "socmedId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := sh.socialMediaService.UpdateSocialMediaById(socmedId, socmedRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (sh *socialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	socmedId, err := helpers.GetParamsId(c, "socmedId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := sh.socialMediaService.DeleteSocialMedia(socmedId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
