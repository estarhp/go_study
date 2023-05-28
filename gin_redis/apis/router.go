package apis

import (
	"gin_redis/apis/middleWare"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleWare.Cors())

	r.POST("/setUser", SetUser)
	r.POST("/getUser", GetUser)

	err := r.Run(`localhost:8000`)
	if err != nil {
		panic(err)
	}

}
