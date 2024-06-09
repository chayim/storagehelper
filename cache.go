package storagehelper

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	con *redis.Client
	ctx context.Context
}

func NewCache() *Cache {
	c := os.Getenv("CACHE_URL")
	if c == "" {
		c = "redis://localhost:6379"
	}

	opt, err := redis.ParseURL(c)
	if err != nil {
		log.Fatalf("Error parsing cache URL: %v", err)
	}
	client := redis.NewClient(opt)

	cache := Cache{con: client, ctx: context.Background()}
	return &cache

}

func (c *Cache) Set(key, val string) error {
	return c.con.Set(c.ctx, key, val, 0).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.con.Get(c.ctx, key).Result()
}

func (c *Cache) Delete(key string) error {
	return c.con.Del(c.ctx, key).Err()
}

func (c *Cache) Flush() error {
	return c.con.FlushDB(c.ctx).Err()
}

func (c *Cache) GetKeys() ([]string, error) {
	return c.con.Keys(c.ctx, "*").Result()
}
