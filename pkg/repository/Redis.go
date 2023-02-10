package repository

import "github.com/go-redis/redis/v8"

type Redis struct {
	Client redis.Client
}
