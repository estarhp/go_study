package main

import (
	"gin_redis/apis"
	"gin_redis/dao"
	"log"
)

func main() {
	err := dao.InitDB()
	apis.InitRouter()

	if err != nil {
		panic(err)
	}
	log.Println("connect successful")

}
