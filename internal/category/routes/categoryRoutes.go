package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/noo-dev/example/internal/category/handler"
	"github.com/noo-dev/example/internal/category/repository"
	"github.com/noo-dev/example/internal/category/service"
)

func InitCategoryRoutes(router *gin.RouterGroup, DB *sqlx.DB) {

	categoryRepo := repository.NewCategoryRepository(DB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	categoryRoutes := router.Group("/category")
	{
		categoryRoutes.GET("/", categoryHandler.GetAllCategories)
		categoryRoutes.GET("/:id", categoryHandler.GetOneCategory)
		// create
		// update
		// delete
	}
}
