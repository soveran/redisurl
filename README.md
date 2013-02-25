redisurl
========

Connect to Redis using a REDIS_URL and the redigo client.

Usage
=====

It uses Redigo[1] under the hood:

    import "redisurl"

    # Connect using os.Getenv("REDIS_URL").
    c, err := redisurl.Connect()

    # Alternatively, connect using a custom Redis URL.
    c, err := redisurl.ConnectToURL("redis://...")

In both cases you will get the result values of `redis.Dial(...)`,
that is, an instance of `redis.Conn` and an error.

Installation
============

Install it using the "go get" command:

go get github.com/soveran/redisurl


