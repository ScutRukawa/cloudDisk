package route

import (
	user "cloudDisk/service/apigateway/handler"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

//Router gw route
func Router() *gin.Engine {
	router := gin.Default()
	dir, _ := os.Getwd()
	fmt.Println("dir:", dir)
	router.Static("/static/", "../../static")

	router.GET("/user/signup", user.SignUpHandler)
	router.POST("/user/signup", user.DoSignupHandler)

	return router
}
