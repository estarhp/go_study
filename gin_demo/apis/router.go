package apis

import (
	"gin_demo/apis/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {

	r := gin.Default()
	r.Use(middleware.Cors())

	r.POST("/register", Register)
	r.GET("/login", Login)

	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.Run(":8000")
}
