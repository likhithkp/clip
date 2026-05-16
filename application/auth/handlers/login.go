package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/clip/application/auth/dto"
	userRepository "github.com/likhithkp/clip/data_access/repository/user"
	_const "github.com/likhithkp/clip/utils/const"
	"github.com/likhithkp/clip/utils/jwt"
	"github.com/likhithkp/clip/utils/other"
	"golang.org/x/crypto/bcrypt"
)

type SignInHandler struct {
	userRepository *userRepository.UserRepository
	utils          *other.ResponseStruct
	jwt            *jwt.GenerateJwtTokenManager
}

func NewSignInHandler(
	userRepository *userRepository.UserRepository,
	utils *other.ResponseStruct,
	jwt *jwt.GenerateJwtTokenManager,
) *SignInHandler {
	return &SignInHandler{
		userRepository: userRepository,
		utils:          utils,
		jwt:            jwt,
	}
}

func (handler *SignInHandler) SignIn(ctx *fiber.Ctx) error {
	newSignIn := new(dto.SignInDto)

	err := ctx.BodyParser(newSignIn)
	if err != nil {
		return handler.utils.Response(ctx, http.StatusUnprocessableEntity, false, "Error while parsing json body", nil)
	}
	if newSignIn.Email == "" || newSignIn.Password == "" {
		return handler.utils.Response(ctx, http.StatusBadRequest, false, "Missing fields", nil)
	}

	existingUser, err := handler.userRepository.GetUserByEmail(ctx.Context(), newSignIn.Email)
	if err != nil {
		log.Printf("error[SignIn]: %v", err.Error())
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}
	if existingUser == nil {
		return handler.utils.Response(ctx, http.StatusUnauthorized, false, "Invalid email or password", nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(newSignIn.Password))
	if err != nil {
		return handler.utils.Response(ctx, http.StatusUnauthorized, false, "Invalid password", nil)
	}

	token, err := handler.jwt.GenerateJWT(existingUser.Id, existingUser.Email, string(_const.User))
	if err != nil {
		log.Printf("error[SignIn]: %v", err.Error())
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}

	return handler.utils.Response(ctx, http.StatusOK, true, "Sign in successful", fiber.Map{
		"id":    existingUser.Id,
		"token": token,
	})
}
