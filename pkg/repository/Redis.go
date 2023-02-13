package repository

import (
	"EFpractic2/models"
	"context"
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client redis.Client
}

func (r *Redis) GetBook(ctx context.Context, bookName string) (models.Book, error) {
	myCache := cache.New(&cache.Options{
		Redis: r.Client,
	})
	book := &models.Book{}
	err := myCache.Get(ctx, bookName, book)
	if err != nil {
		return *book, fmt.Errorf("redis - GetByLogin - Get: %w", err)
	}
	return *book, nil
}

func (r *Redis) CacheBook(ctx context.Context, book *models.Book) error {
	mycache := cache.New(&cache.Options{
		Redis: r.Client,
	})

	err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   book.BookName,
		Value: book,
	})
	if err != nil {
		return fmt.Errorf("redis - CreateBook - Set: %w", err)
	}
	return nil
}

func (c *Redis) RedisStreamInit(ctx context.Context) error {
	_, err := c.Client.XAdd(ctx, &redis.XAddArgs{
		Stream: "example ",
	}).Result()
	if err != nil {
		return fmt.Errorf("redis - RedisStreamInit - XAdd: %w", err)
	}

	_, err = c.Client.XGroupCreate(ctx, "example", "user", "").Result()
	if err != nil {
		return fmt.Errorf("redis - RedisStreamInit - XGroupCreate: %w", err)
	}

	return nil
}
