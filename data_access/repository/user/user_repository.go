package user

import (
	"context"

	userMongoService "github.com/likhithkp/clip/data_access/mongo/user"
	"github.com/likhithkp/clip/data_access/repository/user/convertor"
	"github.com/likhithkp/clip/domain/user"
)

type UserRepository struct {
	userMongoService *userMongoService.UserMongoService
}

func NewUserRepository(userMongoService *userMongoService.UserMongoService,
) *UserRepository {
	return &UserRepository{
		userMongoService: userMongoService,
	}
}

func (repository *UserRepository) UpsertUser(ctx context.Context, domain *user.UserDomain) error {
	entity, err := convertor.DomainToEntity(domain)
	if err != nil {
		return err
	}

	return repository.userMongoService.UpsertUser(ctx, entity)
}

func (repository *UserRepository) GetUserByEmail(ctx context.Context, email string) (domain *user.UserDomain, err error) {
	entity, err := repository.userMongoService.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, nil
	}

	return convertor.EntityToDomain(entity), nil
}
