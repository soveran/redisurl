package redisurl

import (
	"os"
	"fmt"
	"net/url"
	"strings"
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

	if db := strings.Split(redisURL.Path, "/")[1]; db != "" {
		c.Do("SELECT", db)
	}

	return
}
