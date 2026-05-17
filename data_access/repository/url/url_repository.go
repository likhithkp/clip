package url

import (
	"context"

	urlMongoService "github.com/likhithkp/clip/data_access/mongo/url"
	"github.com/likhithkp/clip/data_access/repository/url/convertor"
	"github.com/likhithkp/clip/domain/url"
)

type UrlRepository struct {
	urlMongoService *urlMongoService.UrlMongoService
}

func NewUrlRepository(
	urlMongoService *urlMongoService.UrlMongoService,
) *UrlRepository {
	return &UrlRepository{
		urlMongoService: urlMongoService,
	}
}

func (repository *UrlRepository) UpsertUrl(ctx context.Context, domain *url.UrlDomain) error {
	entity, err := convertor.DomainToEntity(domain)
	if err != nil {
		return err
	}

	return repository.urlMongoService.UpsertUrl(ctx, entity)
}

func (repository *UrlRepository) GetUrlByCode(ctx context.Context, email string) (domain *url.UrlDomain, err error) {
	entity, err := repository.urlMongoService.GetUrlByCode(ctx, email)
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
