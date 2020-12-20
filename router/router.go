package router

import (
	"cloudDisk/handler"

	"github.com/gin-gonic/gin"
)

// Router define route rule
func Router() *gin.Engine {

	router := gin.Default()

	router.Static("/static/", "./static")

	router.GET("/user/signup", handler.SignUpHandler)
	router.POST("/user/signup", handler.DoSignInHandler)

	router.GET("/user/signin", handler.SignInHandler)
	router.POST("/user/signin", handler.DoSignInHandler)

	router.GET("/file/upload", handler.UploadHandler)
	router.POST("/file/upload", handler.UploadHandler)

	router.GET("/file/upload/get", handler.GetFileMetaHandler)

	//router.User(handler.HTTPInterceptor)
	return router
}
