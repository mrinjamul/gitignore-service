package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type API interface {
	Welcome(ctx *gin.Context)
}

type api struct {
}

func (api *api) Welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the API",
	})
}

func NewAPI() API {
	return &api{}
}
