package main

import (
	"cloudDisk/service/apigateway/route"
)

func main() {

	router := route.Router()
	router.Run(":8080")
}
