package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/gitignore-service/utils"
)

type Gitignore interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type gitignore struct {
}

func (gitignore *gitignore) List(ctx *gin.Context) {
	metadata, err := utils.GetMetadata("")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "something went wrong",
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"gitignores": metadata,
	})
}

func (gitignore *gitignore) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "return a specific git ignore contents",
	})
}

func NewGitignore() Gitignore {
	return &gitignore{}
}
