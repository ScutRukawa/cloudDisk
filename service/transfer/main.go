package main

import (
	"cloudDisk/config"
	"cloudDisk/db"
	"cloudDisk/mq"
	"cloudDisk/store/ceph"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/amz.v1/s3"
)

func ProcessTransfer(msg []byte) bool {
	fmt.Println("start ProcessTransfer")
	putData := mq.TransferData{}
	err := json.Unmarshal(msg, &putData)
	if err != nil {
		log.Println("json.Unmarshal failed: ", err)
		return false
	}

	fileFd, err := os.Open(putData.CurLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	bucket := ceph.GetCephBucket("testbucket1")
	fileData, err := ioutil.ReadAll(fileFd)
	if err != nil {
		log.Println("ioutil.ReadAll err:", err)
	}
	cephPath := putData.Destination
	_ = bucket.Put(cephPath, fileData, "octet-stream", s3.PublicRead)

	db.UpdateFileLocation(putData.FileHash, putData.Destination)
	fmt.Println(" ProcessTransfer success")
	return true
}
func main() {
	log.Println("开始监听转移任务队列...")
	mq.StartConsume(
		config.TransOSSQueueName, //to to change type
		"ceph",
		ProcessTransfer,
	)
}
