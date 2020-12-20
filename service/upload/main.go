package main

import (
	"cloudDisk/handler"
	"cloudDisk/router"
	"net/http"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	//http.HandleFunc("/file/upload", handler.UploadHandler)
	//http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	//http.HandleFunc("/file/upload/get", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	//http.HandleFunc("/user/signup", handler.SignUpHandler)
	//http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))
	http.HandleFunc("/health", handler.HealthCheckHandler)
	http.HandleFunc("/file/query", handler.QueryUseFileMetaHandler)
	http.HandleFunc("/file/fastupload", handler.TryFastUploadHandler)
	http.HandleFunc("/file/mpupload/init", handler.InitialMultipartUploadHandler)
	http.HandleFunc("/file/mpupload/uppart", handler.UploadPartHandler)
	http.HandleFunc("/file/mpupload/complete", handler.CompleteUploadHandler)
	router.Router().Run(":8080")

}
