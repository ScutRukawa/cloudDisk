package main

import (
	"cloudDisk/store/ceph"
	"fmt"

	"gopkg.in/amz.v1/s3"
)

func main() {
	bucket := ceph.GetCephBucket("testbucket1")

	err := bucket.PutBucket(s3.PublicRead) //创建bucket，指定权限级别

	fmt.Printf("create bucket err:%v\n", err)

	//查询指定条件对象，指定100条
	res, err := bucket.List("", "", "", 100)
	fmt.Printf("object key :%+v\n", res)

	//上传一个对象
	err = bucket.Put("/home/wei/java_error_in_GOLAND_5764.log", []byte("just for test"), "octet-stream", s3.PublicRead)
	fmt.Printf("bucket.Put:%+v\n", err)

	res, err = bucket.List("", "", "", 100)
	fmt.Printf("object key :%+v\n", res)

}
