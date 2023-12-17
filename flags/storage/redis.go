package storage

import (
	"context"
	"fmt"
	"github.com/markelca/toggles/flags"
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
    keys,err := r.client.Keys(ctx, "*").Result()
    if err != nil {
        return nil,err
    }

    result := make([]flags.Flag,len(keys))

    for i,key := range keys {
        val,err := r.Get(key)
        if err != nil {
            return nil,err
        }
        result[i] = flags.Flag{
            Name: key,
            Value: val,
        }
    }
    return result,nil
}

func (r RedisRepository) Create(flag flags.Flag) error {
    exists,err := r.Exists(flag.Name)
    if err != nil {
        return err
    } else if exists {
        return flags.FlagAlreadyExistsError
    } 

    err = r.client.Set(ctx, flag.Name, flag.Value, 0).Err()
    if err != nil {
        return err
    }

    return nil
}

func (r RedisRepository) Exists(name string) (bool,error) {
    result,err := r.client.Exists(ctx,name).Result()
    if err != nil {
        return false,err
    }
    return result != 0, nil
}

func (r RedisRepository) Update(name string, value bool) error {
    exists,err := r.Exists(name)
    if err != nil {
        return err
    } else if !exists {
        return flags.FlagNotFoundError
    }
    err = r.client.Set(ctx,name,value,0).Err()
    if err != nil {
        return err
    }
    return nil
}
