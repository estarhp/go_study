package dao

// 假数据库，用 map 实现
var database = map[string]string{
	"yxh": "123456",
	"wx":  "654321",
}

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

func AddUser(username, password string) {

	database[username] = password

}

func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordByUsername(username string) string {
	return database[username]
}
