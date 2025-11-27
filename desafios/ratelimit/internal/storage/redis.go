package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisStorage implementa LimiterStorage usando Redis.
type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr, password string, db int) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisStorage{client: client}
}

func (r *RedisStorage) Increment(key string, limit int, expireSeconds int) (int, int, error) {
	ctx := context.Background()
	pipe := r.client.TxPipeline()
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, time.Duration(expireSeconds)*time.Second)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, 0, err
	}
	count := int(incr.Val())
	ttl, err := r.client.TTL(ctx, key).Result()
	if err != nil {
		return count, 0, err
	}
	return count, int(ttl.Seconds()), nil
}

func (r *RedisStorage) Get(key string) (int, int, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Int()
	if err == redis.Nil {
		return 0, 0, nil
	}
	if err != nil {
		return 0, 0, err
	}
	ttl, err := r.client.TTL(ctx, key).Result()
	if err != nil {
		return val, 0, err
	}
	return val, int(ttl.Seconds()), nil
}

func (r *RedisStorage) Block(key string, blockSeconds int) error {
	ctx := context.Background()
	blockKey := fmt.Sprintf("block:%s", key)
	return r.client.Set(ctx, blockKey, 1, time.Duration(blockSeconds)*time.Second).Err()
}

func (r *RedisStorage) IsBlocked(key string) (bool, int, error) {
	ctx := context.Background()
	blockKey := fmt.Sprintf("block:%s", key)
	val, err := r.client.Get(ctx, blockKey).Result()
	if err == redis.Nil {
		return false, 0, nil
	}
	if err != nil {
		return false, 0, err
	}
	ttl, err := r.client.TTL(ctx, blockKey).Result()
	if err != nil {
		return true, 0, err
	}
	return val == "1", int(ttl.Seconds()), nil
}
