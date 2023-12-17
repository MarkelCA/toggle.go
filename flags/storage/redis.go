package storage

import (
	"context"
	"fmt"
	"github.com/markelca/toggle.go/flags"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisRepository struct {
    client *redis.Client 
}

func NewRedisRepository(host string, port int) flags.FlagRepository {
    hostStr := fmt.Sprintf("%v:%v",host,port)
    rdb := redis.NewClient(&redis.Options{
        Addr:     hostStr,
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    return RedisRepository{rdb}
}

func(r RedisRepository) Get(key string) (bool, error) {
    val, err := r.client.Get(ctx, key).Bool()
    if err == redis.Nil {
        return false,flags.FlagNotFoundError
    }
    return val,err
}

func (r RedisRepository) List()([]flags.Flag, error) {
    x := r.client.Keys(ctx, "*")
    fmt.Println(x)
    // r.client.Get("foo")

    return nil,nil
}

func (r RedisRepository) Create(flag flags.Flag) error {
    err := r.client.Set(ctx, "key", "value", 0).Err()

    if err != nil {
        panic(err)
    }

    return nil
}

func (r RedisRepository) Exists(name string) (bool,error) {
    return false,nil
}

func (r RedisRepository) Update(name string, value bool) error {
    val2, err := r.client.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
    return nil
}
