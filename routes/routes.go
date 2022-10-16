package routes

import (
	"final_MyGram/controllers"
	"final_MyGram/middlewares"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type router struct {
	auth        *controllers.AuthController
	user        *controllers.UserController
	comment     *controllers.CommentController
	photo       *controllers.PhotoController
	socialMedia *controllers.SocialMediaController
}

func NewRouter(
	auth *controllers.AuthController,
	user *controllers.UserController,
	photo *controllers.PhotoController,
	comment *controllers.CommentController,
	socialMedia *controllers.SocialMediaController,
) *router {
	return &router{
		auth:        auth,
		user:        user,
		photo:       photo,
		comment:     comment,
		socialMedia: socialMedia,
	}
}

func (rt *router) Start(PORT string) {
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger())
	// v1 := r.Group("/api/v1")
	// swagger := v1.Group("/swagger")
	swagger := r.Group("/")
	{
		swagger.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	// user := v1.Group("/users")
	user := r.Group("/users")
	{
		user.POST("/register", rt.auth.RegisterController)
		user.POST("/login", rt.auth.LoginController)
	}
	// v1Auth := v1.Group("/auth")
	// secured := v1Auth.Group("/users").Use(middlewares.Auth())
	secured := r.Group("/users").Use(middlewares.Auth())
	{
		secured.PUT("/:userId", rt.user.UpdateUserController)
		secured.DELETE("", rt.user.DeleteUserController)
		secured.POST("/logout", rt.auth.LogoutController)
	}
	// photo := v1Auth.Group("/photos").Use(middlewares.Auth())
	photo := r.Group("/photos").Use(middlewares.Auth())
	{
		photo.POST("", rt.photo.CreatePhotoController)
		photo.GET("", rt.photo.GetAllPhotoController)
		photo.PUT("/:photoId", rt.photo.UpdatePhotoController)
		photo.DELETE("/:photoId", rt.photo.DeletePhotoController)
	}

	// comment := v1Auth.Group("/comments/").Use(middlewares.Auth())
	comment := r.Group("/comments").Use(middlewares.Auth())
	{
		comment.POST("", rt.comment.CreateCommentController)
		comment.GET("", rt.comment.GetAllCommentController)
		comment.PUT("/:commentId", rt.comment.UpdateCommentController)
		comment.DELETE("/:commentId", rt.comment.DeleteCommentController)
	}

	// socialMedia := v1Auth.Group("/socialmedias").Use(middlewares.Auth())
	socialMedia := r.Group("/socialmedias").Use(middlewares.Auth())
	{
		socialMedia.POST("", rt.socialMedia.CreateSocialMediaController)
		socialMedia.GET("", rt.socialMedia.GetAllSocialMediaController)
		socialMedia.PUT("/:socialMediaId", rt.socialMedia.UpdateSocialMediaController)
		socialMedia.DELETE("/:socialMediaId", rt.socialMedia.DeleteSocialMediaController)
	}
	r.Run(PORT)
}
