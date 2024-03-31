package main

import (
	"fmt"
	"context"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	hashkeys := "user_profile"
	fieldandvalues := map[string]interface{}{
		"name":"alna",
		"emile":"2222@111",
		"age":30,
	}
	err := rdb.HMSet(ctx,hashkeys,fieldandvalues).Err()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("Hash data stored successfully")

	result,err := rdb.HGetAll(ctx,hashkeys).Result()
	if err!= nil{
		fmt.Println(err)
		return
	}

	fmt.Println("Hash data:")
    for field, value := range result {
        fmt.Printf("%s: %s\n", field, value)
    }
/*
	err := rdb.Set(ctx,"address","长沙",0).Err()
	if err != nil{
		print(err)
	}
	val , err := rdb.Get(ctx,"address").Result()
	if err!= nil{
		print(err)
	}
	fmt.Println("address",val)
*/
}