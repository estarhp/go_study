package dao

import (
	"github.com/go-redis/redis"
)

var RedisDB *redis.Client

func InitDB() (err error) {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //密码，没有就写空字符串""
		DB:       0,
	})

	_, err = redisDB.Ping().Result()
	if err != nil {
		return err
	}

	RedisDB = redisDB
	return nil

}

//	func Set(key string, val string) error {
//		return RedisDB.Set(key, val, 0).Err()
//	}
//
//	func Get(key string) (string, error) {
//		return RedisDB.Get(key).Result()
//	}
func Exists(key string) (int64, error) {
	return RedisDB.Exists(key).Result()
}

func Hset(key string, user map[string]interface{}) (string, error) {
	return RedisDB.HMSet(key, user).Result()
}

func Hget(key string) (map[string]string, error) {
	return RedisDB.HGetAll(key).Result()
}

//func Zset(key string, val []redis.Z) error {
//	_, err := RedisDB.ZAdd(key, val...).Result()
//	if err != nil {
//		fmt.Printf("Zadd failed,err:%v\n", err)
//		return err
//	}
//
//	return err
//}

//key := "language_rank"
//language := []redis.Z{
//	redis.Z{Score: 80, Member: "go"},
//	redis.Z{
//		Score:  88,
//		Member: "java",
//	},
//	redis.Z{
//		Score:  89,
//		Member: "c++",
//	},
//}
//
//err = dao.Zset(key, language)
//
//if err != nil {
//	log.Fatalln("zset redis is err:", err)
//}
//
//op := redis.ZRangeBy{
//	Min: "60",
//	Max: "85",
//}
//
//ret, err := dao.RedisDB.ZRangeByScoreWithScores(key, op).Result()
//
//if err != nil {
//	fmt.Println("zrangebyscore failed ,err :", err)
//	return
//}
//
//for _, z := range ret {
//	fmt.Println(z.Member, z.Score)
//}
