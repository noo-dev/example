package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/noo-dev/example/internal/post/handler"
	"github.com/noo-dev/example/internal/post/repository"
	"github.com/noo-dev/example/internal/post/service"
)

func InitPostRoutes(router *gin.RouterGroup, DB *sqlx.DB) {

	postRepo := repository.NewPostRepository(DB)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	postRoutes := router.Group("/post")
	{
		postRoutes.GET("/", postHandler.GetAllPosts)
		postRoutes.GET("/:id", postHandler.GetOnePost)
		// create
		// update
		// delete
	}
}
