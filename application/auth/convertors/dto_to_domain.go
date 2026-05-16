package convertors

import (
	"time"

	"github.com/likhithkp/clip/application/auth/dto"
	"github.com/likhithkp/clip/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertSignUpDtoToDomain(dto *dto.SignUpDto) *user.UserDomain {
	return &user.UserDomain{
		Id:        primitive.NewObjectID().Hex(),
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
