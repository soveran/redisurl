package redisurl

import (
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestConnect(t *testing.T) {
	c, err := Connect()

	if err != nil {
		t.Errorf("Error returned")
	}

	pong, err := redis.String(c.Do("PING"))

	if err != nil {
		t.Errorf("Call to PING returned an error: %v", err)
	}

	if pong != "PONG" {
		t.Errorf("Wanted PONG, got %v\n", pong)
	}
}

func TestNewPool(t *testing.T) {
	pool, err := NewPool(3, 200, "240s")

	if err != nil {
		t.Error(err)
	}

	c := pool.Get()
	defer c.Close()
	pong, err := redis.String(c.Do("PING"))

	if err != nil {
		t.Errorf("Call to PING returned an error: %v", err)
	}

	if pong != "PONG" {
		t.Errorf("Wanted PONG, got %v\n", pong)
	}
}
