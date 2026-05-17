package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	urlRepository "github.com/likhithkp/clip/data_access/repository/url"
	"github.com/likhithkp/clip/utils/other"
)

type GetUrlHandler struct {
	utils         *other.ResponseStruct
	urlRepository *urlRepository.UrlRepository
}

func NewGetUrlHandler(
	utils *other.ResponseStruct,
	urlRepository *urlRepository.UrlRepository,
) *GetUrlHandler {
	return &GetUrlHandler{
		utils:         utils,
		urlRepository: urlRepository,
	}
}

func (handler *GetUrlHandler) GetUrl(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	if code == "" {
		return handler.utils.Response(ctx, http.StatusBadRequest, false, "Code not provided", nil)
	}

	longUrl, err := handler.urlRepository.GetUrl(ctx.Context(), code)
	if err != nil {
		log.Printf("error[GetUrl(Redis)]: %s", err.Error())
	}

	if longUrl == "" {
		urlDomain, err := handler.urlRepository.GetUrlByCode(ctx.Context(), code)
		if err != nil {
			log.Printf("error[GetUrl(MongoDB)]: %v", err.Error())
			return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
		}
		if urlDomain == nil {
			return handler.utils.Response(ctx, http.StatusNotFound, false, "URL not found", nil)
		}
		longUrl = urlDomain.LongUrl

		if err := handler.urlRepository.SetUrl(ctx.Context(), code, longUrl); err != nil {
			log.Printf("error[SetUrl(Redis)]: %s", err.Error())
		}
	}

	return ctx.Redirect(longUrl, http.StatusMovedPermanently)
}
