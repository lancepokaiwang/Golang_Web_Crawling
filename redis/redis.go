/*
	First run:
	docker pull redis


	Running docker:
	docker run --name redis-test-instance -p 6379:6379 -d redis
*/
package redis

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
	"github.com/pkg/errors"
)

type Redis struct {
	client *redis.Client
}

type ProductResponse struct {
}

func (s ProductResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s ProductResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

func NewClient() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return Redis{client: client}
}

func (r *Redis) Insert(keyword string, value []productPB.ProductResponse) error {
	return r.client.Set(keyword, value, 300).Err()
}

func (r *Redis) Query(keyword string) (string, error) {
	val2, err := r.client.Get(keyword).Result()
	if err == redis.Nil {
		return "", errors.Wrapf(err, "keyword %v does not exist", keyword)
	} else if err != nil {
		return "", errors.Wrap(err, "failed to query keyword")
	}
	return val2, nil
}
