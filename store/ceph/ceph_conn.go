package ceph

import (
	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

var cephConn *s3.S3

func GetcephConnection() *s3.S3 {
	if cephConn != nil {
		return cephConn
	}

	//1 初始化ceph
	auth := aws.Auth{
		AccessKey: "B2E5C1SD4YDR99QMNFYJ",
		SecretKey: "ak5NaOS9FkS5MvZi2oIDHIgxK6BAPHX09sVXqrPG",
	}

	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          "http://127.0.0.1:9080", //what mean
		S3Endpoint:           "http://127.0.0.1:9080",
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
	}
	return s3.New(auth, curRegion)
}

//获取bucket对象
func GetCephBucket(bucket string) *s3.Bucket {
	conn := GetcephConnection()
	return conn.Bucket(bucket)
}
