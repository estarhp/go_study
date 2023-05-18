package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type student struct {
	id   int
	name string
}

var db *sql.DB

func initDB() {
	var err error
	// 设置一下dns charset:编码方式 parseTime:是否解析time类型 loc:时区
	dsn := "root:123456@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开mysql驱动

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}

func insertStudent(st student) {
	sqlStr := "insert into student(id,name) values (?,?)"
	_, err := db.Exec(sqlStr, st.id, st.name)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("insert success")
}

// func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
func queryStudent() {
	sqlStr := "select id,name from student where id = ?"
	var u student
	err := db.QueryRow(sqlStr, 11).Scan(&u.id, &u.name)

	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return

	}

	println(u.id)

	//defer row.Close()
	//fmt.Println(rows.Next())
	//for rows.Next() {
	//	var u student
	//	err := rows.Scan(&u.id, &u.name)
	//
	//	if err != nil {
	//		println(err)
	//		return
	//	}
	//	fmt.Printf("id:%d name:%s \n", u.id, u.name)
	//}

}

func main() {

	initDB()
	//for i := 0; i < 10; i++ {
	//	insertStudent(student{
	//		id:   i,
	//		name: "tang",
	//	})
	//}

	queryStudent()

}
