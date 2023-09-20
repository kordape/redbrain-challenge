// Package http implements routing paths. Each services in own file.
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kordape/redbrain-challenge/pkg/encryptorapi"
)

func NewRouter(
	handler *gin.Engine,
	encryptor encryptorapi.EncryptorClient,
) {
	// Options
	handler.Use(gin.Recovery())

	r := routes{}
	echo := handler.Group("/hash-movie-name")
	{
		echo.POST("", r.hashMovieNameHandler(encryptor))
	}
}
