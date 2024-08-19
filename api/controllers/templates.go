package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Templates interface {
	Index(ctx *gin.Context)
}

type templates struct {
}

func (t *templates) Index(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.tmpl",
		gin.H{
			"title":   "Gitignore service | Home",
			"appname": "Gitignore service",
		},
	)
}

func NewTemplates() Templates {
	return &templates{}
}
