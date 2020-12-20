package handler

import (
	"cloudDisk/common"
	"cloudDisk/db"
	"cloudDisk/service/account/proto"
	"context"
)

type User struct{}

func (user *User) SignUp(ctx context.Context, req *proto.ReqSignUp, res *proto.ResSignUp) error {

	username := req.GetUsername()
	passwd := req.GetPassword()

	ret := db.UserSignUp(username, passwd)
	if ret == false {
		res.Code = common.StatusRegisterFaild
		res.Message = "注册失败"
	} else {
		res.Code = common.StatusOK
		res.Message = "注册成功"
	}
	//fmt.Fprintf(w,"UserSignUp success,your user name is:%s",username)
	return nil
}
