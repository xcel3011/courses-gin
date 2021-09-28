package global

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var pool *redis.Pool

func init() {
	dial := func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", _config.RedisConn.Address)
		if err != nil {
			return nil, err
		}
		if _config.RedisConn.Password != "" {
			if _, err := c.Do("AUTH", "password01!"); err != nil {
				c.Close()
				return nil, err
			}
		}
		return c, nil
	}
	pool = &redis.Pool{
		Dial:        dial,
		IdleTimeout: time.Duration(_config.RedisConn.IdleTimeout) * time.Second,
		MaxIdle:     _config.RedisConn.MaxIdle,
		MaxActive:   _config.RedisConn.MaxActive,
		Wait:        true,
	}
}

func RedisSet(k, v string) error {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", k, v)

	if err != nil {
		log.Println("redis set error:", err)
		return err
	}
	return nil
}

func RedisSetEX(k, v string, time time.Duration) error {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", k, v, "EX", time.Seconds())

	if err != nil {
		log.Println("redis set error: ", err)
		return err
	}
	return nil
}

func RedisGet(key string) (bool, string, error) {
	conn := pool.Get()
	defer conn.Close()
	result, err := redis.String(conn.Do("GET", key))
	if err == redis.ErrNil {
		return false, "", nil
	}
	if err != nil {
		log.Println("redis get error: ", err)
		return false, "", err
	}
	return true, result, nil
}

func RedisExpire(key string, time time.Duration) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("EXPIRE", key, time.Seconds())
	if err != nil {
		log.Println("redis Expire failed:", err)
		return err
	}
	return nil
}
