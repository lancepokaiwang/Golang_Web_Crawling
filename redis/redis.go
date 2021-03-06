/*
	First run:
	docker pull redis


	Running docker:
	docker run --name redis-test-instance -p 6379:6379 -d redis
*/

/*
	How-to:
	(0) Initialize redis instance
	r := redis.NewClient()

	(1) Insert Data
	if err := r.Insert(<KEYWORD>, <PRODUCTS>); err != nil {
		// Error handling here
	}

	(2) Query Data
	products, err := r.Query(<KEYWORD>)
	if err != nil {
		// Error handling here
	}
*/
package redis

import (
	"encoding/json"
	"time"

	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type Redis struct {
	client *redis.Client
}

func NewClient() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return Redis{client: client}
}

func (r *Redis) Insert(keyword string, value ProductSlice) error {
	data, err := value.marshalBinary()
	if err != nil {
		return errors.Wrap(err, "failed to marshal value")
	}
	return r.client.Set(keyword, data, 180*time.Second).Err()
}

func (r *Redis) Query(keyword string) ([]productPB.ProductResponse, error) {
	data, err := r.client.Get(keyword).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "failed to query keyword")
	}

	var results ProductSlice

	if err := results.unmarshalBinary(data); err != nil {
		return nil, errors.Wrap(err, "failed to unmarsh data")
	}

	return results.Products, nil
}

func (r *Redis) Update(keyword string, value ProductSlice) error {
	data_exist, err := r.Query(keyword)
	if err != nil {
		return errors.Wrap(err, "failed to query existing data")
	}

	// Combine exist data.
	if data_exist != nil {
		value.Products = append(value.Products, data_exist...)
	}

	data, err := value.marshalBinary()
	if err != nil {
		return errors.Wrap(err, "failed to marshal value")
	}

	return r.client.Set(keyword, data, 180*time.Second).Err()
}

type ProductSlice struct {
	Products []productPB.ProductResponse
}

// MarshalBinary translate ProductResponse to []byte format which is accepted by Redis.
func (ps *ProductSlice) marshalBinary() ([]byte, error) {
	return json.Marshal(ps.Products)
}

func (ps *ProductSlice) unmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &ps.Products)
}
