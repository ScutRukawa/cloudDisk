package db

import (
	"fmt"
	mydb "cloudDisk/db/mysql"
	"log"
)
//注册用户
func UserSignUp(username string,passwd string) bool{
	stmt,err:=mydb.DBConn().Prepare(
		"insert ignore into tbl_user(`user_name`,`user_pwd`)values(?,?)")
	if err!=nil{
		log.Println(err.Error())
		return false
	}
	defer stmt.Close()

	ret,err:=stmt.Exec(username,passwd)
	if err!=nil{
		log.Println(err.Error())
		return false
	}
	rows,err:=ret.RowsAffected()
	if nil==err&&rows>0{
		fmt.Println("RowsAffected success")
		return true
	}
	return false
}
// UpdateToken : 刷新用户登录的token
func UpdateToken(username string, token string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"replace into tbl_user_token (`user_name`,`user_token`) values (?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func GetUserInfo(username string) (User, error) {
	user := User{}

	stmt, err := mydb.DBConn().Prepare(
		"select user_name,signup_at from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer stmt.Close()

	// 执行查询的操作
	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}