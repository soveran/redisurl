package redisurl

import (
	"os"
	"fmt"
	"net/url"
	"github.com/garyburd/redigo/redis"
)

func Connect() (redis.Conn, error) {
	return ConnectToURL(os.Getenv("REDIS_URL"))
}

func ConnectToURL(s string) (c redis.Conn, err error) {
	redis_url, _ := url.Parse(s)
	auth := ""

	if password, ok := redis_url.User.Password(); ok {
		auth = password
	}

	c, err = redis.Dial("tcp", redis_url.Host)

	if err != nil {
		fmt.Println(err)
		return
	}

	if auth != "" {
		_, err = c.Do("AUTH", auth)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	return
}
