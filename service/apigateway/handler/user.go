package user

import (
	"cloudDisk/service/account/proto"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
)

var userCli proto.UserService

func main() {
	//创建user客户端  开启service 创建handler进行rpc调用

	service := micro.NewService()
	service.Init()

	userCli = proto.NewUserService("go.micro.service.user", service.Client())

}

//SignUpHandler is
func SignUpHandler(c *gin.Context) {
	dir, _ := os.Getwd()
	fmt.Println("pwd:", dir)
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

//DoSignupHandler is
func DoSignupHandler(c *gin.Context) { //post
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	res, err := userCli.SignUp(context.TODO(), &proto.ReqSignUp{
		Username: username,
		Password: passwd,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  res.Code,
		"code": res.Message,
	})
}
