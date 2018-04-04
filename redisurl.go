package redisurl

import (
	"os"
	"fmt"
	"net/url"
	"strings"
	"github.com/gomodule/redigo/redis"
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

	options := []redis.DialOption{}

	if redisURL.Scheme == "rediss" {
		options = append(options, redis.DialUseTLS(true))
	}

	c, err = redis.Dial("tcp", redisURL.Host, options...)

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
