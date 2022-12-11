package main

import (
	"context"
	"log"

	"github.com/muzfr7/simple-redis-client/cache"
)

const (
	RedisAddr     string = "localhost:6379"
	RedisPassword string = ""
	RedisDB       int    = 0
)

func main() {
	ctx := context.Background()

	// creating a Redis client
	c := cache.New(RedisAddr, RedisPassword, RedisDB)
	if err := c.Ping(ctx); err != nil {
		log.Panic("failed to connect to Redis")
	}

	log.Println("connected to Redis..")

	// creating a key with value in Redis
	if err := c.Set(ctx, "user:name", "muzafar", 0); err != nil {
		log.Println("Error: could not store a value in Redis")
	}

	log.Println("value stored in Redis")

	// retrieving a key from Redis
	res, err := c.Get(ctx, "user:name")
	if err != nil {
		log.Println("Error: could not get a value from Redis")
	}

	log.Printf("Result: %s\n", res)
}
