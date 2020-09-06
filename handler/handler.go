package handler

import (
	"encoding/json"
	"fmt"
	"gopan/db"
	"gopan/meta"
	"gopan/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		data,err:=ioutil.ReadFile("./static/view/index.html")
		log.Println("receive upload file request ")
		if err!=nil{
			io.WriteString(w,err.Error())
			return
		}
		io.WriteString(w,string(data))
	}else if r.Method=="POST"{
		fmt.Println("recieve file")
		file,head,err:=r.FormFile("file")
		defer file.Close()
		if err!=nil{
			fmt.Printf("failed to read data:%s",err) //很多err 的判断 改进
			return
		}
		fileMeta:=meta.FileMeta{
			FileName: head.Filename,
			Location: "C:\\eznewei\\Mydocuments\\agogogogo\\"+head.Filename, ///change
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		newFile,err:=os.Create(fileMeta.Location)
		if err!=nil{
			fmt.Println("create file err:",err)
			return
		}
		fileMeta.FileSize,err=io.Copy(newFile,file)
		if err!=nil{
			fmt.Println("save data to file failed:",err)
			return
		}
		newFile.Seek(0,0)
		fileMeta.FileSha1=util.FileSha1(newFile)
		meta.UpdateFileMetaDb(fileMeta)

		// update user file table
		r.ParseForm()
		username:=r.Form.Get("username")
		suc:=db.OnUserFileUploadFinished(username,fileMeta.FileSha1,fileMeta.FileName,fileMeta.FileSize)
		if suc {
			http.Redirect(w,r,"/static/view/home.html",http.StatusFound)
		}else{
			w.Write([]byte("upload failed"))
		}
	}
}
func UploadSucHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"update success")
}

func GetFileMetaHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fileSha1:=r.Form["filehash"][0]
	//	fileSha1:=r.Form.Get("filehash")
	testData:=r.Form["filehash"]
	fmt.Println("test data:",testData)
	filemata,err:=meta.GetFileMetaDB(fileSha1)
	if err!=nil{
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data,err:=json.Marshal(filemata)
	if err!=nil{
		fmt.Println("json format error:",err)
		io.WriteString(w,"json format error")
		return
	}
	io.WriteString(w,string(data))
	w.Write(data)
}

//func DownloadHandler(w http.ResponseWriter,r *http.Request){
//	r.ParseForm()
//	fileSha1:=r.Form["filehash"][0]
//	filemata:=meta.GetFileMeta(fileSha1)
//	//file,err:=os.OpenFile(filemata.Location+filemata.FileName,os.O_RDONLY,0777)
//	log.Println("file loaction:",filemata.Location)
//	fileData,err:=ioutil.ReadFile(filemata.Location)
//
//	if err!=nil{
//		fmt.Println("read file err:",err)
//		return
//	}
//	io.WriteString(w,string(fileData))
//}
func DownloadHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fileSha1:=r.Form["filehash"][0]
	filemata:=meta.GetFileMeta(fileSha1)
	//file,err:=os.Open(filemata.Location) //和 openfile 的区别？
	file,err:=os.OpenFile(filemata.Location+filemata.FileName,os.O_RDONLY,0777)
	defer file.Close()
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data,err:=ioutil.ReadAll(file)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/octect-stream")
	w.Header().Set("Content-Descrption","attachment;filename=\""+filemata.FileName+"\"")

	w.Write(data)

}
func HealthCheckHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"health!!!")
}
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")

			//验证登录token是否有效
			if len(username) < 3 || !IsTokenValid(token) {
				// w.WriteHeader(http.StatusForbidden)
				// token校验失败则跳转到登录页面
				http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
				return
			}
			h(w, r)
		})
}
func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	// TODO: 判断token的时效性，是否过期
	// TODO: 从数据库表tbl_user_token查询username对应的token信息
	// TODO: 对比两个token是否一致
	return true
}

func QueryUseFileMetaHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	username:=r.Form.Get("username")
	userfiles,err:=db.QueryUseFileMeta(username)
	if err!=nil{
		fmt.Println("QueryUseFileMeta err",err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	data,err:=json.Marshal(userfiles)
	if err!=nil{
		fmt.Println("json.Marshal err",err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println("send data to ajax")
	w.Write(data)
}