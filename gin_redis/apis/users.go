package apis

import (
	"encoding/json"
	"gin_redis/dao"
	"gin_redis/model"
	"gin_redis/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func SetUser(c *gin.Context) {
	err := c.ShouldBind(&model.User{})
	if err != nil {
		log.Println(err)
		utils.RespFailed(c, "You missed a parameter")
		return
	}
	id := c.PostForm("id")
	log.Println(id, "被传入")

	isExists, err := dao.Exists("gin_redis:" + id)

	if err != nil {
		log.Fatalln(err)
		return
	}

	if isExists == 1 {
		utils.RespFailed(c, "the user already exist")
		return
	}
	var user = make(map[string]interface{})
	user["id"] = id
	user["name"] = c.PostForm("name")
	user["age"] = c.PostForm("age")
	user["sex"] = c.PostForm("sex")

	_, err = dao.Hset("gin_redis:"+id, user)

	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("add user successfully")
	utils.RespSuccess(c, "add user successfully")

}

func GetUser(c *gin.Context) {
	id := c.PostForm("id")

	isExists, err := dao.Exists("gin_redis:" + id)
	if err != nil {
		log.Println(err)
		return
	}
	if isExists == 0 {
		utils.RespFailed(c, "the user does not exists")
		return
	}
	var user = make(map[string]string)
	user, err = dao.Hget("gin_redis:" + id)
	if err != nil {
		utils.RespFailed(c, "internal err")
		log.Fatalln(err)
		return
	}
	jsonUser, err := json.Marshal(user)

	if err != nil {
		utils.RespFailed(c, "internal err")
	}

	s := string(jsonUser)

	utils.RespSuccess(c, s)

}
