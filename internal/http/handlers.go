package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kordape/redbrain-challenge/pkg/encryptorapi"
	"github.com/kordape/redbrain-challenge/pkg/gatewayapi"
)

type routes struct{}

func (r *routes) hashMovieNameHandler(encryptor encryptorapi.EncryptorClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var hashMovieNameRequest gatewayapi.HashMovieNameRequest
		if err := json.Unmarshal(requestBody, &hashMovieNameRequest); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		resp, err := encryptor.Encrypt(context.Background(), &encryptorapi.EncryptRequest{
			Data: hashMovieNameRequest.Movie,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gatewayapi.HashMovieNameResponse{
				Error: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gatewayapi.HashMovieNameResponse{
			Encrypted: resp.Encrypted,
		})
	}
}
