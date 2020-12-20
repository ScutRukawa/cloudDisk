package handler

import (
	"cloudDisk/db"
	mydb "cloudDisk/db/mysql"
	"cloudDisk/util"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	dir, _ := os.Getwd()
	fmt.Println("pwd:", dir)
	c.Redirect(http.StatusFound, "./static/view/signup.html")
}
func DoSignupHandler(c *gin.Context) { //post
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	if len(username) < 3 || len(passwd) < 5 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "invalid parameter",
			"code": -1,
		})
		return
	}

	ret := db.UserSignUp(username, passwd)
	if ret {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Sign up Success",
			"code": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Sign up fail",
			"code": -2,
		})
	}
}
func SignInHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "./static/view/signin.html")
	return
}
func DoSignInHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	ret := checkUser(username, password)
	fmt.Println("ret:", ret)
	if !ret {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "password error",
			"code": -1,
		})
		return
	}
	//2，生产访问token
	token := GenToken(username)
	upRes := db.UpdateToken(username, token)
	if !upRes {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "token error",
			"code": -2,
		})
		return
	}
	//3，重定向到home
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
	return
}
func checkUser(username string, password string) bool {
	stmt, err := mydb.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	row, err := stmt.Query(username)
	if err != nil {
		log.Println(err.Error())
		return false
	} else if row == nil {
		fmt.Println("username not fund")
		return false
	}
	pRows := mydb.ParseRows(row)
	passwd, ok := pRows[0]["user_pwd"]
	pass, ok2 := passwd.([]uint8)
	fmt.Println("type is :", reflect.TypeOf(passwd))
	fmt.Println("passwd is :", pass)
	fmt.Println("ok is :", ok2)
	fmt.Println("string(pass) :", string(pass))
	if ok && len(pRows) > 0 && string(pass) == password {
		return true
	}
	log.Println("pass word is null")
	return false
}
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	//	token := r.Form.Get("token")

	// // 2. 验证token是否有效
	// isValidToken := IsTokenValid(token)
	// if !isValidToken {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }

	// 3. 查询用户信息
	user, err := db.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// 4. 组装并且响应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}
