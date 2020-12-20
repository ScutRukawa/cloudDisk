package main

import (
	"cloudDisk/service/account/handler"
	"cloudDisk/service/account/proto"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	//"github.com/micro/go-micro/registry" //
	//"github.com/micro/go-micro/registry/etcd"
)

var etcdReg registry.Registry

// func init() {
// 	cli, err := clientv3.New(clientv3.Config{
// 		Endpoints:   []string{"127.0.0.1:12379"},
// 		DialTimeout: 5 * time.Second,
// 	})
// 	if err != nil {
// 		fmt.Printf("connect to etcd failed", err)
// 		return
// 	}
// 	fmt.Println("connect to etcd success")
// }

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
	)
	service.Init()

	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
