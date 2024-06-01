package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(host string, port uint) CacheClient {
	hostStr := fmt.Sprintf("%v:%v", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     hostStr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return RedisClient{rdb}
}

func (r RedisClient) Delete(key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r RedisClient) Keys() ([]string, error) {
	return r.client.Keys(ctx, "*").Result()
}

func (r RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", Nil
	}
	return val, err
}

func (r RedisClient) Expire(key string, expiration time.Duration) error {
	return r.client.Expire(ctx, key, expiration).Err()
}

func (r RedisClient) GetList(key string) ([]string, error) {
	return r.client.LRange(ctx, key, 0, -1).Result()
}

func (r RedisClient) AppendToList(key string, expiration time.Duration, values ...any) error {
	err := r.client.LPush(ctx, key, values).Err()
	r.client.Expire(ctx, key, expiration).Err()
	return err
}

func (r RedisClient) Set(key string, value any, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r RedisClient) Exists(name string) (bool, error) {
	result, err := r.client.Exists(ctx, name).Result()
	if err != nil {
		return false, err
	}
	return result != 0, nil
}
