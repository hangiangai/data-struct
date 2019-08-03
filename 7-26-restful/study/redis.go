package study

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Redis() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //没有密码
		DB:       0,  //实用默认数据库
	})

	//Ping
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong, err)

	fmt.Println("==============================")

	Get(client)
}

func Get(r *redis.Client) {

}
