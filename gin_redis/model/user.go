package model

type User struct {
	Id   int    `form:"id" json:"id"  `
	Name string `form:"name" json:"name" `
	Age  int    `form:"age" json:"age" `
	Sex  string `form:"sex" json:"sex" `
}
