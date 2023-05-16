package apis

import (
	"gin_demo/apis/middleware"
	"gin_demo/dao"
	"gin_demo/model"
	"gin_demo/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func Register(c *gin.Context) {

	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespSuccess(c, "verification failed")
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	println(username, password)

	flag := dao.SelectUser(username)

	if flag {
		utils.RespFail(c, "user already exists")

		return
	}

	dao.AddUser(username, password)
	utils.RespSuccess(c, "add user successful")

}

func Login(c *gin.Context) {

	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}

	username := c.Query("username")
	password := c.Query("password")
	println(username, password)

	flag := dao.SelectUser(username)

	if !flag {
		utils.RespFail(c, "user doesn't exists")

		return
	}

	realPassword := dao.SelectPasswordByUsername(username)

	if password != realPassword {
		utils.RespFail(c, "wrong password")

		return
	}

	claim := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "Yxh",                                // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, tokenString)

}

func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}
