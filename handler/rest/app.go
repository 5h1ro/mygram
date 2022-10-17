package rest

import (
	"fmt"
	"mygram/comment"
	"mygram/database"
	"mygram/docs"
	"mygram/middleware"
	"mygram/photo"
	"mygram/social_media"
	"mygram/user"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/thedevsaddam/govalidator"
)

const port = ":8080"

func StartApp() {
	database.InitializeDB()

	db, dbase := database.GetDB()

	uniqueRule := database.NewUniqueRule(db, "unique")
	govalidator.AddCustomRule("unique", uniqueRule.Rule)
	userRepository := user.NewRepository(dbase)
	userService := user.NewService(userRepository)
	userHandler := NewUser(userService)

	photoRepository := photo.NewRepository(dbase)
	photoService := photo.NewService(photoRepository)
	photoHandler := NewPhoto(photoService, userService)

	commentRepository := comment.NewRepository(dbase)
	commentService := comment.NewService(commentRepository)
	commentHandler := NewComment(commentService, photoService, userService)

	socialMediaRepository := social_media.NewRepository(dbase)
	socialMediaService := social_media.NewService(socialMediaRepository)
	socialMediaHandler := NewSocialMedia(socialMediaService, userService)

	route := gin.Default()

	docs.SwaggerInfo.Title = "MyGram Documentation"
	docs.SwaggerInfo.Description = "Ini adalah dokumentasi mygram"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := route.Group("/api/v1")
	auth := v1.Group("/users")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}

	mw := v1.Group("")
	{
		mw.Use(middleware.Authentication())
	}

	user := mw.Group("/users")
	{
		user.GET("/:userID", userHandler.FindUser)
		user.PUT("", userHandler.UpdateUser)
		user.DELETE("", userHandler.DeleteUser)
	}

	photo := mw.Group("/photos")
	{
		photo.Use(middleware.IsExistMiddleware())
		photo.GET("", photoHandler.GetPhoto)
		photo.POST("", photoHandler.CreatePhoto)
		photo.PUT("/:photoId", middleware.PhotoAuthorization(), photoHandler.UpdatePhoto)
		photo.DELETE("/:photoId", middleware.PhotoAuthorization(), photoHandler.DeletePhoto)
	}

	comment := mw.Group("/comments")
	{
		comment.Use(middleware.IsExistMiddleware())
		comment.GET("", commentHandler.GetComment)
		comment.POST("", commentHandler.CreateComment)
		comment.PUT("/:commentId", middleware.CommentAuthorization(), commentHandler.UpdateComment)
		comment.DELETE("/:commentId", middleware.CommentAuthorization(), commentHandler.DeleteComment)
	}

	socialMedia := mw.Group("/socialmedias")
	{
		socialMedia.Use(middleware.IsExistMiddleware())
		socialMedia.GET("", socialMediaHandler.GetSocialMedia)
		socialMedia.POST("", socialMediaHandler.CreateSocialMedia)
		socialMedia.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), socialMediaHandler.UpdateSocialMedia)
		socialMedia.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), socialMediaHandler.DeleteSocialMedia)
	}

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)
}
