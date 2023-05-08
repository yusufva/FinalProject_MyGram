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

// GetAllSocialMedia godoc
// @Tags Social Media
// @Description Get All Social Media Data
// @ID get-all-social-Media
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.GetSocialMediaResponse
// @Router /socmeds [get]
func (sh *socialMediaHandler) GetAllSocialMedia(c *gin.Context) {
	allSocmed, err := sh.socialMediaService.GetAllSocialMedia()

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(allSocmed.StatusCode, allSocmed)
}

// GetSocialMediaById godoc
// @Tags Social Media
// @Description Get Social Media By Id Data
// @ID get-social-media-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socmedId path int true "socmedId"
// @Success 200 {object} dto.SocialMediaResponse
// @Router /socmeds/{socmedId} [get]
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

// GetSocialMediaByUser godoc
// @Tags Social Media
// @Description Get Social Media By User Data
// @ID get-social-media-by-User
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param userId path int true "userId"
// @Success 200 {object} dto.GetSocialMediaResponse
// @Router /socmeds/users/{userId} [get]
func (sh *socialMediaHandler) GetSocialMediaByUser(c *gin.Context) {
	userId, err := helpers.GetParamsId(c, "userId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := sh.socialMediaService.GetSocialMediaByUser(userId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
	}

	c.JSON(http.StatusOK, &response)
}

// CreateSocialMedia godoc
// @Tags Social Media
// @Description Create New Social Media Data
// @ID create-new-social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewSocialMediaRequest true "request body json"
// @Success 201 {object} dto.NewSocialMediaResponse
// @Router /socmeds [post]
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

// UpdateSocialMediaById godoc
// @Tags Social Media
// @Description Update Social Media Data
// @ID update-social-media-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewSocialMediaRequest true "request body json"
// @Param socmedId path int true "socmedId"
// @Success 200 {object} dto.NewSocialMediaResponse
// @Router /socmeds/{socmedId} [put]
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
	user := c.MustGet("userData").(entity.User)

	response, err := sh.socialMediaService.UpdateSocialMediaById(socmedId, user.ID, socmedRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// DeleteSocialMediaById godoc
// @Tags Social Media
// @Description Delete Social Media Data
// @ID delete-social-media-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socmedId path int true "socmedId"
// @Success 200 {object} dto.NewSocialMediaResponse
// @Router /socmeds/{socmedId} [delete]
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
