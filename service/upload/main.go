package main

import (
	"cloudDisk/handler"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/upload/get", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/user/signup", handler.SignUpHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))
	http.HandleFunc("/health", handler.HealthCheckHandler)
	http.HandleFunc("/file/query", handler.QueryUseFileMetaHandler)
	http.HandleFunc("/file/fastupload", handler.TryFastUploadHandler)
	http.HandleFunc("/file/mpupload/init", handler.InitialMultipartUploadHandler)
	http.HandleFunc("/file/mpupload/uppart", handler.UploadPartHandler)
	http.HandleFunc("/file/mpupload/complete", handler.CompleteUploadHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("create http server:", err)
	}
}

//scanArgs := make([]interface{}, 10)
//values := make([]interface{}, 10)
//for j := range values {
//	scanArgs[j] = &values[j]
//}
//parse(scanArgs...)
//fmt.Println(reflect.TypeOf(scanArgs[0]))
//str,ok1:=scanArgs[0].(*interface{})
//strReal:=*str
//fmt.Println("reflect.TypeOf(scanArgs[0])  strReal",reflect.TypeOf(scanArgs[0]))
//fmt.Println("ok:",ok1)
//
//	fmt.Println(strReal)
//for index,value:=range scanArgs{
//	fmt.Println(value)
//	fmt.Println(index)
//
//}
//for _,value:=range values{
//	fmt.Println(value)
//}

func parse(in ...interface{}) {
	fmt.Println("parse start")
	for _, value := range in {
		fmt.Println("reflect.TypeOf(value):", reflect.TypeOf(value))
		switch v := value.(type) {
		case *interface{}:
			fmt.Println("i am *interface")
			*v = "string"
		case []interface{}:
			fmt.Println("[]interface {}xxxx")

		}
	}
}
