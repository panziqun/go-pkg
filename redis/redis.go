package redis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/laughmaker/go-pkg/config"
)

var RedisPool *redis.Pool

func Setup() error {
	RedisPool = &redis.Pool{
		MaxIdle:     config.RedisConfig.MaxIdle,
		MaxActive:   config.RedisConfig.MaxActive,
		IdleTimeout: config.RedisConfig.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			addres := fmt.Sprintf("%s:%d", config.RedisConfig.Host, config.RedisConfig.Port)
			c, err := redis.Dial("tcp", addres)
			if err != nil {
				return nil, err
			}

			if config.RedisConfig.Password != "" {
				if _, err := c.Do("AUTH", config.RedisConfig.Password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", config.RedisConfig.Db); err != nil {
				c.Close()
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func Set(key string, data interface{}, time int) error {
	conn := RedisPool.Get()
	defer conn.Close()

	var value string
	switch data.(type) {
	case string:
		value = data.(string)
	case int, int64:
		value = strconv.Itoa(data.(int))
	default:
		str, err := json.Marshal(data)
		if err != nil {
			return err
		}
		value = string(str)
	}

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := RedisPool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) (string, error) {
	conn := RedisPool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return string(reply), nil
}

func Delete(key string) (bool, error) {
	conn := RedisPool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisPool.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
