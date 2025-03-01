package repository

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheInterface interface {
	Set(ctx context.Context, userName string, otp string, ttl time.Duration) error
	Get(ctx context.Context, userName string) (string, error)
	GenerateRandomOTP() string
}

type CacheRepository struct {
	db *redis.Client
}

func NewCacheRepository(db *redis.Client) *CacheRepository {
	return &CacheRepository{
		db: db,
	}
}

func (c *CacheRepository) Set(ctx context.Context, userName string, otp string, ttl time.Duration) error {
	err := c.db.Set(ctx, userName, otp, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *CacheRepository) Get(ctx context.Context, userName string) (string, error) {
	res, err := c.db.Get(ctx, userName).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (c *CacheRepository) GenerateRandomOTP() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	otp := r.Intn(999999-100000) + 100000

	return strconv.Itoa(otp)
}
