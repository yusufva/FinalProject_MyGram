package service

import (
	"final-project/entity"
	"final-project/pkg/errrs"
	"final-project/pkg/helpers"
	"final-project/repository/comment_repository"
	"final-project/repository/photo_repository"
	"final-project/repository/socialMedia_repository"
	"final-project/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AuthorizationPhoto() gin.HandlerFunc
	AuthorizationComment() gin.HandlerFunc
	AuthorizationSocmed() gin.HandlerFunc
}

type authService struct {
	userRepo        user_repository.UserRepository
	photoRepo       photo_repository.PhotoRepository
	commentRepo     comment_repository.CommentRepository
	socialMediaRepo socialMedia_repository.SocialMediaRepository
}

func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository, commentRepo comment_repository.CommentRepository, socialMediaRepo socialMedia_repository.SocialMediaRepository) AuthService {
	return &authService{
		userRepo:        userRepo,
		photoRepo:       photoRepo,
		commentRepo:     commentRepo,
		socialMediaRepo: socialMediaRepo,
	}
}

func (a *authService) AuthorizationPhoto() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("userData").(entity.User)

		photoId, err := helpers.GetParamsId(c, "photoId")

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		photo, err := a.photoRepo.GetPhotoById(photoId)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if photo.UserID != user.ID {
			unauthorizedErr := errrs.NewUnauthorizedError("you are not authorized to modify the product data")
			c.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		c.Next()
	}
}

func (a *authService) AuthorizationComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("userData").(entity.User)

		commentId, err := helpers.GetParamsId(c, "commentId")

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		comment, err := a.commentRepo.GetCommentById(commentId)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if comment.UserID != user.ID {
			unauthorizedErr := errrs.NewUnauthorizedError("you are not authorized to modify the product data")
			c.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		c.Next()
	}
}

func (a *authService) AuthorizationSocmed() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("userData").(entity.User)

		socmedId, err := helpers.GetParamsId(c, "socmedId")

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		socmed, err := a.socialMediaRepo.GetSocialMediaById(socmedId)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if socmed.UserID != user.ID {
			unauthorizedErr := errrs.NewUnauthorizedError("you are not authorized to modify the social media data")
			c.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		c.Next()
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invalidTokenErr = errrs.NewUnauthenticatedError("invalid token")
		bearerToken := c.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			c.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		c.Set("userData", user)

		c.Next()
	}
}
