package dao

import (
	"database/sql"
	"fmt"
	"gin_demo/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 假数据库，用 map 实现
var database = make(map[string]string)

var DB *sql.DB

//尝试失败我实在是不想写了
//var database = make(map[string]string)

//func InitDataBase() (err error) {
//
//	fileData, err := os.ReadFile("data.json")
//
//	if err != nil {
//		return fmt.Errorf("open file error %v", err)
//	}
//
//	err = json.Unmarshal(fileData, &database)
//
//	if err != nil {
//
//		return fmt.Errorf("error decoding JSON: %v", err)
//	}
//
//	return nil
//}

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_demo?charset=utf8mb4"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	log.Println("DB connect successfully")

	DB = db

}

func AddUser(username, password string) (err error) {
	sqlStr := "insert into user(username,password) values (?,?)"

	_, err = DB.Exec(sqlStr, username, password)
	if err != nil {
		return err
	}
	log.Println("insert successfully")

	return nil
}

func SelectUser(username string) bool {
	var u model.User
	sqlStr := "select username,password from user where username = ?"
	log.Println(username)
	err := DB.QueryRow(sqlStr, username).Scan(&u.Password, &u.Username)
	if err != nil {
		println(err)
		return false
	}
	return true

}

func SelectPasswordByUsername(username string) string {
	var u model.User
	sqlStr := "select username,password from user where username = ?"
	DB.QueryRow(sqlStr, username).Scan(&u.Username, &u.Password)
	fmt.Println(u.Password)

	return u.Password

}
