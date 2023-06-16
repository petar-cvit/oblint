package internal

import (
	"context"
	"math/rand"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Storage struct {
	client *redis.Client
}

func NewStorage() (Storage, error) {
	host := os.Getenv("REDIS_HOST")

	client := redis.NewClient(&redis.Options{
		Addr: host,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return Storage{}, err
	}

	return Storage{
		client: client,
	}, nil
}

func (s Storage) Random() error {
	random := rand.Int()

	return s.client.Set(context.Background(), strconv.Itoa(random), random, 0).Err()
}
