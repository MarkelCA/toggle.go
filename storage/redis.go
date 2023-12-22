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

func NewRedisClient(host string, port int) CacheClient {
    hostStr := fmt.Sprintf("%v:%v",host,port)
    rdb := redis.NewClient(&redis.Options{
        Addr:     hostStr,
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    return RedisClient{rdb}
}

func(r RedisClient) Keys() ([]string, error) {
    return r.client.Keys(ctx, "*").Result()
}

func(r RedisClient) Get(key string) (string, error) {
    val, err := r.client.Get(ctx, key).Result()
    
    if err == redis.Nil {
        return "",Nil
    }
    return val,err
}

func (r RedisClient) Expire(key string, expiration time.Duration) error {
    return r.client.Expire(ctx,key,expiration).Err()
}

func (r RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
    return r.client.Set(ctx, key, value, expiration).Err()
}

func (r RedisClient) Exists(name string) (bool,error) {
    result,err := r.client.Exists(ctx,name).Result()
    if err != nil {
        return false,err
    }
    return result != 0, nil
}
