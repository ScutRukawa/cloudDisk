package handler

import (
	"fmt"
	"gopan/db"
	mydb "gopan/db/mysql"
	"gopan/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

func SignUpHandler(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		data,err:=ioutil.ReadFile("./static/view/signup.html")
		if err!=nil{
			io.WriteString(w,err.Error())
			return
		}
		io.WriteString(w,string(data))
	}else if r.Method=="POST"{
		fmt.Println("start parse")
		r.ParseForm()
		username:=r.Form.Get("username")
		passwd:=r.Form.Get("password")
		//passwdc:=r.Form.Get("passwdc")
		//if passwd!=passwdc{
		//	io.WriteString(w,"passwd is different from the first")
		//	return
		//}
		ret:=db.UserSignUp(username,passwd)
		if ret==false{
			log.Println("UserSignUp failed")
			//io.WriteString(w,"UserSignUp failed")
			http.NotFound(w,r)
		}
		//fmt.Fprintf(w,"UserSignUp success,your user name is:%s",username)
		if ret {
			w.Write([]byte("SUCCESS"))
		} else {
			w.Write([]byte("FAILED"))
		}
	}else{
		fmt.Println("method err")
	}
}
func SignInHandler(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		data,err:=ioutil.ReadFile("./static/view/signin.html")
		if err!=nil{
			http.NotFound(w,r)
		}
		io.WriteString(w,string(data))
		return
	} else if r.Method=="POST"{
		//1，密码校验
		r.ParseForm()
		username:=r.Form.Get("username")
		password:=r.Form.Get("password")
		ret:=checkUser(username,password)
		fmt.Println("ret:",ret)
		if !ret {
			log.Println("password is not correct")
			io.WriteString(w,"password is not correct")
			return
		}
		//2，生产访问token
		token := GenToken(username)
		upRes := db.UpdateToken(username, token)
		if !upRes {
			w.Write([]byte("FAILED"))
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
				Location: "http://" + r.Host + "/static/view/home.html",
				Username: username,
				Token:    token,
			},
		}
		w.Write(resp.JSONBytes())

		//w.Write([]byte("/static/view/home.html")) //三种有何却别

		//http.Redirect(w,r,"/static/view/home.html",http.StatusFound)
		//data,err:=ioutil.ReadFile("./static/view/home.html")
		//if err!=nil{
		//	fmt.Println("not fund")
		//	http.NotFound(w,r)
		//}
		//io.WriteString(w,string(data))
		fmt.Println("http://" + r.Host + "/static/view/home.html")

	}

}
func checkUser(username string,password string)bool{
	stmt,err:=mydb.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}
	row,err:=stmt.Query(username)
	if err!=nil{
		log.Println(err.Error())
		return false
	}else if row==nil{
		fmt.Println("username not fund")
		return false
	}
	pRows:=mydb.ParseRows(row)
	passwd,ok:=pRows[0]["user_pwd"]
	pass,ok2:=passwd.([]uint8)
	fmt.Println("type is :",reflect.TypeOf(passwd))
	fmt.Println("passwd is :",pass)
	fmt.Println("ok is :",ok2)
	fmt.Println("string(pass) :",string(pass))
	if ok&&len(pRows)>0&&string(pass)==password{
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