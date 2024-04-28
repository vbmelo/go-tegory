package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vbmelo/go-tegory/internal/repositories"
	use_cases "github.com/vbmelo/go-tegory/internal/use-cases"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCase := use_cases.NewListCategoryUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success":    true,
		"categories": categories,
	})
}
