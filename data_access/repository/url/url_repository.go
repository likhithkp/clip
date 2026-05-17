package url

import (
	"context"

	urlMongoService "github.com/likhithkp/clip/data_access/mongo/url"
	urlRedisService "github.com/likhithkp/clip/data_access/redis/url"
	"github.com/likhithkp/clip/data_access/repository/url/convertor"
	"github.com/likhithkp/clip/domain/url"
)

type UrlRepository struct {
	urlMongoService *urlMongoService.UrlMongoService
	urlRedisService *urlRedisService.UrlRedisService
}

func NewUrlRepository(
	urlMongoService *urlMongoService.UrlMongoService,
	urlRedisService *urlRedisService.UrlRedisService,
) *UrlRepository {
	return &UrlRepository{
		urlMongoService: urlMongoService,
		urlRedisService: urlRedisService,
	}
}

func (repository *UrlRepository) UpsertUrl(ctx context.Context, domain *url.UrlDomain) error {
	entity, err := convertor.DomainToEntity(domain)
	if err != nil {
		return err
	}

	return repository.urlMongoService.UpsertUrl(ctx, entity)
}

func (repository *UrlRepository) GetUrlByCode(ctx context.Context, code string) (domain *url.UrlDomain, err error) {
	entity, err := repository.urlMongoService.GetUrlByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, nil
	}

	return convertor.EntityToDomain(entity), nil
}

func (repository *UrlRepository) GetUserById(ctx context.Context, id string) (domain *url.UrlDomain, err error) {
	entity, err := repository.urlMongoService.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, nil
	}

	return convertor.EntityToDomain(entity), nil
}

//Redis

func (repository *UrlRepository) SetUrl(ctx context.Context, code string, longUrl string) error {
	err := repository.urlRedisService.SetURL(ctx, code, longUrl)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UrlRepository) GetUrl(ctx context.Context, code string) (string, error) {
	url, err := repository.urlRedisService.GetURL(ctx, code)
	if err != nil {
		return "", err
	}

	return url, nil
}
