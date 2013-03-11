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

	return
}
