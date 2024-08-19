package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/gitignore-service/utils"
)

type Gitignore interface {
	List(ctx *gin.Context, metadata *[]utils.Metadata)
	Get(ctx *gin.Context, metadata *[]utils.Metadata)
}

type gitignore struct {
}

func (gitignore *gitignore) List(ctx *gin.Context, metadata *[]utils.Metadata) {
	lang := ctx.Query("for")
	if lang != "" {
		meta, _ := utils.SearchForGi(lang)
		if len(meta) != 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message":    "success",
				"gitignores": meta,
			})
			return
		}

		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no gitignore found",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"gitignores": metadata,
	})
}

func (gitignore *gitignore) Get(ctx *gin.Context, metadata *[]utils.Metadata) {
	lang := ctx.Query("for")
	if lang == "" {
		lang = "go"
	}

	meta, _ := utils.SearchForGi(lang)
	if len(meta) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "success",
			"gitignore": meta[0],
		})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "no gitignore found",
	})
}

func NewGitignore() Gitignore {
	return &gitignore{}
}
