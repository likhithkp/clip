package url

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type UrlRedisService struct {
	client *redis.Client
}

func NewUrlRedisService(client *redis.Client) *UrlRedisService {
	return &UrlRedisService{client: client}
}

func (s *UrlRedisService) SetURL(ctx context.Context, code, longURL string) error {
	return s.client.Set(ctx, "url:"+code, longURL, 24*time.Hour).Err()
}

func (s *UrlRedisService) GetURL(ctx context.Context, code string) (string, error) {
	longURL, err := s.client.Get(ctx, "url:"+code).Result()
	if err == redis.Nil {
		return "", nil
	}
	return longURL, err
}
