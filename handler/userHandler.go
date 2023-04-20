package handler

import (
	"final-project/dto"
	"final-project/pkg/errrs"
	"final-project/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

// Registers godoc
// @Tags Registers
// @Description User Registers
// @ID user-register
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewUserRequest true "request body json"
// @Success 201 {object} dto.NewUserRequest
// @Router /users/register [post]
func (uh *userHandler) Register(c *gin.Context) {
	var newUserRequest dto.NewUserRequest

	if err := c.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.CreateNewUser(newUserRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(result.StatusCode, result)
}

// Login godoc
// @Tags login
// @Description User Login
// @ID user-login
// @Accept json
// @Produce json
// @Param RequestBody body dto.LoginRequest true "request body json"
// @Success 201 {object} dto.LoginRequest
// @Router /users/login [post]
func (uh *userHandler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.Login(loginRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(result.StatusCode, result)
}
