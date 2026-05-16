package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/clip/application/auth/convertors"
	"github.com/likhithkp/clip/application/auth/dto"
	userRepository "github.com/likhithkp/clip/data_access/repository/user"
	"github.com/likhithkp/clip/utils/other"
	"golang.org/x/crypto/bcrypt"
)

type SignUpHandler struct {
	utils          *other.ResponseStruct
	userRepository *userRepository.UserRepository
}

func NewSignUpHandler(
	utils *other.ResponseStruct,
	userRepository *userRepository.UserRepository,
) *SignUpHandler {
	return &SignUpHandler{
		utils:          utils,
		userRepository: userRepository,
	}
}

func (handler *SignUpHandler) SignUp(ctx *fiber.Ctx) error {
	newUser := new(dto.SignUpDto)

	err := ctx.BodyParser(newUser)
	if err != nil {
		return handler.utils.Response(ctx, http.StatusUnprocessableEntity, false, "Error while parsing json body", nil)
	}
	if newUser.Email == "" || newUser.FirstName == "" || newUser.Password == "" {
		return handler.utils.Response(ctx, http.StatusBadRequest, false, "Missing fields", nil)
	}
	existingUser, err := handler.userRepository.GetUserByEmail(ctx.Context(), newUser.Email)
	if existingUser != nil {
		return handler.utils.Response(ctx, http.StatusConflict, false, "User with the email "+newUser.Email+" already exists", nil)
	}
	if err != nil {
		log.Printf("error: %v", err.Error())
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}

	newUserDomain := convertors.ConvertSignUpDtoToDomain(newUser)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(newUserDomain.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}

	newUserDomain.Password = string(hashedBytes)
	err = handler.userRepository.UpsertUser(ctx.Context(), newUserDomain)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}

	return handler.utils.Response(ctx, http.StatusCreated, true, "Sign-up successful", nil)
}
