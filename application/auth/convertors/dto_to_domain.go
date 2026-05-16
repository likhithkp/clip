package convertors

import (
	"time"

	"github.com/likhithkp/clip/application/auth/dto"
	"github.com/likhithkp/clip/domain/user"
)

func ConvertSignUpDtoToDomain(dto *dto.SignUpDto) *user.UserDomain {
	return &user.UserDomain{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
