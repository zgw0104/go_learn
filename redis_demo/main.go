package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100, // 连接池大小，根据实际情况决定
	})

	_, err = rdb.Ping().Result()
	return err
}

func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Println("get name failed")
		return
	} else {
		fmt.Println("name", val2)
	}
}

func hsetDemo() {
	rdb.HSet("user", "name", "gw")
	rdb.HSet("user", "age", 18)

	data := map[string]interface{}{
		"name": "my",
		"age":  18,
	}

	rdb.HMSet("user2", data)
}

func hgetDemo() {
	val, err := rdb.HGetAll("user").Result()
	if err != nil {
		fmt.Println("hgetAll failed")
		return
	}
	fmt.Println("hgetAll:", val)

	v2 := rdb.HGet("user", "name").Val()
	fmt.Println("hget:", v2)

	v3 := rdb.HMGet("user", "name", "age").Val()
	fmt.Println("hmget:", v3)

}

func zsetDemo() {
	zsetKey := "language_Rank"
	language := []redis.Z{
		redis.Z{Score: 90.0, Member: "Go"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	num, err := rdb.ZAdd(zsetKey, language...).Result()
	if err != nil {
		fmt.Println("zadd failed", err)
		return
	}
	fmt.Println("zadd num:", num)

	//给某一个元素添加值
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Go").Result()
	if err != nil {
		fmt.Println("zincr failed", err)
		return
	}
	fmt.Println("Go's score:", newScore)

	//取前三个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Println("zrevrange failed", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	//取分数95-100之间的
	op := redis.ZRangeBy{ //定义标志
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Println("zrangebyscore failed", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	if err := initRedis(); err != nil {
		fmt.Println("init redis client failed ,err:", err)
		return
	}

	fmt.Println("redis client init success")
	defer rdb.Close()
	//redisExample()

	//hsetDemo()
	//hgetDemo()
	zsetDemo()
}
