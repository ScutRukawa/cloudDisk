package db

import (
	"fmt"
	mydb "cloudDisk/db/mysql"
	"time"
)

type Userfile struct{
	UserName string
	FileHash  string
	FileName string
	FileSize int64
	UploadAt string
	LastUpdated string
}
//update user file table
func OnUserFileUploadFinished(UserName,FileHash,FileName string,FileSize int64) bool{
	stmt,err:=mydb.DBConn().Prepare(
		"insert ignore into tbl_user_file(`user_name`,`file_sha1`,`file_name`," +
			"`file_size`,`upload_at`)values(?,?,?,?,?)")
	if err!=nil{
		fmt.Println("prepare err :",err)
		return false
	}
	defer stmt.Close()
	_,err=stmt.Exec(UserName,FileHash,FileName,FileSize,time.Now())
	if err!=nil{
		fmt.Println("stmt exec err:",err)
	}
	return true
}

func QueryUseFileMeta(username string) ([]Userfile,error){
	stmt,err:=mydb.DBConn().Prepare("select file_size,file_sha1,file_name,last_update,upload_at " +
		"from tbl_user_file where user_name=?")
	if err!=nil{
		fmt.Println("QueryUseFileMeta Prepare err:", err)
		return nil,err
	}
	var userfiles []Userfile
	var userfile = Userfile{}
	rows,err2:=stmt.Query(username)
	if err2!=nil{
		fmt.Println("QueryUseFileMeta Query err:", err)
		return nil,err2
	}
	for rows.Next(){
		err=rows.Scan(&userfile.FileSize,&userfile.FileHash,
			&userfile.FileName,&userfile.LastUpdated,&userfile.UploadAt)
		if err!=nil{
			fmt.Println("QueryUseFileMeta scan err:", err)
			break
		}
		userfiles=append(userfiles,userfile)
	}
	return userfiles,nil
}