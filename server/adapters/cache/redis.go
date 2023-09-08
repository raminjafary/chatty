package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisClient(redisURL string) (*RedisCache, error) {
	opts, err := redis.ParseURL(redisURL)

	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)
	_, err = client.Ping().Result()

	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &RedisCache{client: client}, nil
}

func (c *RedisCache) Get(key string, value interface{}) error {
	data, err := c.client.Get(key).Result()
	if err == redis.Nil {
		return fmt.Errorf("cache miss for key %q", key)
	} else if err != nil {
		return fmt.Errorf("failed to get value for key %q: %v", key, err)
	}

	if err := json.Unmarshal([]byte(data), value); err != nil {
		return fmt.Errorf("failed to unmarshal cache value for key %q: %v", key, err)
	}

	return nil
}

func (c *RedisCache) Set(key string, value interface{}, duration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal cache value for key %q: %v", key, err)
	}

	if err := c.client.Set(key, data, duration).Err(); err != nil {
		return fmt.Errorf("failed to set value for key %q: %v", key, err)
	}

	return nil
}

func (c *RedisCache) Delete(key string) error {
	if err := c.client.Del(key).Err(); err != nil {
		return fmt.Errorf("failed to delete value for key %q: %v", key, err)
	}
	return nil
}
