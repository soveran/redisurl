package redisurl

import (
	"os"
	"fmt"
	"net/url"
	"strings"
	"github.com/garyburd/redigo/redis"
	"time"
)

func Connect() (redis.Conn, error) {
	return ConnectToURL(os.Getenv("REDIS_URL"))
}

func ConnectToURL(s string) (c redis.Conn, err error) {
	redisURL, err := url.Parse(s)

	if err != nil {
		return
	}

	auth := ""

	if redisURL.User != nil {
		if password, ok := redisURL.User.Password(); ok {
			auth = password
		}
	}

	c, err = redis.Dial("tcp", redisURL.Host)

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(auth) > 0 {
		_, err = c.Do("AUTH", auth)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if len(redisURL.Path) > 1 {
		db := strings.TrimPrefix(redisURL.Path, "/")
		c.Do("SELECT", db)
	}

	return
}

func NewPool(maxIdle, maxActive int, idleTimeout string) (p *redis.Pool, err error) {
	return NewPoolWithURL(os.Getenv("REDIS_URL"), maxIdle, maxActive, idleTimeout)
}

func NewPoolWithURL(redisURL string, maxIdle, maxActive int, idleTimeout string) (p *redis.Pool, err error) {
	timeout, err := time.ParseDuration(idleTimeout)
	if err != nil {
		return
	}

	p =  &redis.Pool{
		MaxIdle: maxIdle,
		MaxActive: maxActive,
		IdleTimeout: timeout,
		Dial: func () (redis.Conn, error) {
			return ConnectToURL(redisURL)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	c := p.Get()
	defer c.Close()
	_, err = c.Do("PING")

	return p, err
}

