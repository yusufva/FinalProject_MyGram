package handler

import (
	"final-project/database"
	"final-project/repository/comment_repository/comment_pg"
	"final-project/repository/photo_repository/photo_pg"
	"final-project/repository/socialMedia_repository/socialMedia_pg"
	"final-project/repository/user_repository/user_pg"
	"final-project/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	var port = "8080"
	database.InitializeDatabase()

	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewUserPg(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	photoRepo := photo_pg.NewPhotoPg(db)
	photoService := service.NewPhotoService(photoRepo)
	photoHandler := NewPhotoHandler(photoService)

	commentRepo := comment_pg.NewCommentPg(db)
	commentService := service.NewCommentService(commentRepo)
	commentHandler := NewCommentHandler(commentService)

	socmedRepo := socialMedia_pg.NewSocialMediaPg(db)
	socmedService := service.NewSocialMediaService(socmedRepo)
	socmedHandler := NewSocialMediaHandler(socmedService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, socmedRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userHandler.Login)
		userRoute.POST("/register", userHandler.Register)
	}

	photoRoute := route.Group("/photos")
	{
		photoRoute.GET("/", photoHandler.GetAllPhoto)
		photoRoute.POST("/", authService.Authentication(), photoHandler.CreateNewPhoto)
	}

	commentRoute := route.Group("/comments")
	{
		commentRoute.GET("/photos/:photoId", commentHandler.GetAllCommentByPhoto)
	}

	socmedRoute := route.Group("/socmeds")
	{
		socmedRoute.GET("/", socmedHandler.GetAllSocialMedia)
	}

	route.Run(":" + port)
}
