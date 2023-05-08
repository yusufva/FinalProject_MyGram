package handler

import (
	"final-project/database"
	"final-project/docs"
	"final-project/repository/comment_repository/comment_pg"
	"final-project/repository/photo_repository/photo_pg"
	"final-project/repository/socialMedia_repository/socialMedia_pg"
	"final-project/repository/user_repository/user_pg"
	"final-project/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "MyGram API"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userHandler.Login)
		userRoute.POST("/register", userHandler.Register)
	}

	photoRoute := route.Group("/photos")
	{
		photoRoute.GET("/", photoHandler.GetAllPhoto)
		photoRoute.GET("/:photoId", photoHandler.GetPhotoById)
		photoRoute.GET("/users/:userId", photoHandler.GetPhotoByUser)
		photoRoute.POST("/", authService.Authentication(), photoHandler.CreateNewPhoto)
		photoRoute.PUT("/:photoId", authService.Authentication(), authService.AuthorizationPhoto(), photoHandler.UpdatePhotoById)
		photoRoute.DELETE("/:photoId", authService.Authentication(), authService.AuthorizationPhoto(), photoHandler.DeletePhotoById)

	}

	commentRoute := route.Group("/comments")
	{
		commentRoute.GET("/users", authService.Authentication(), commentHandler.GetAllCommentByUser)
		commentRoute.GET("/photos/:photoId", commentHandler.GetAllCommentByPhoto)
		commentRoute.GET("/:commentId", authService.Authentication(), commentHandler.GetCommentById)
		commentRoute.POST("/photos", authService.Authentication(), commentHandler.CreateComment)
		commentRoute.PUT("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.UpdateCommentById)
		commentRoute.DELETE("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.DeleteCommentById)
	}

	socmedRoute := route.Group("/socmeds")
	{
		socmedRoute.GET("/", socmedHandler.GetAllSocialMedia)
		socmedRoute.GET("/:socmedId", socmedHandler.GetSocialMediaById)
		socmedRoute.GET("/users/:userId", authService.Authentication(), socmedHandler.GetSocialMediaByUser)
		socmedRoute.POST("/", authService.Authentication(), socmedHandler.CreateSocialMedia)
		socmedRoute.PUT("/:socmedId", authService.Authentication(), authService.AuthorizationSocmed(), socmedHandler.UpdateSocialMediaById)
		socmedRoute.DELETE("/:socmedId", authService.Authentication(), authService.AuthorizationSocmed(), socmedHandler.DeleteSocialMedia)
	}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	route.Run(":" + port)
}
