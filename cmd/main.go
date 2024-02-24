package cmd

import (
	"github.com/gin-gonic/gin"
	categoryRoutes "github.com/noo-dev/example/internal/category/routes"
	postRoutes "github.com/noo-dev/example/internal/category/routes"
)

func main() {
	app := gin.Default()
	DB, err := getDBConnection()
	api := app.Group("api")
	postRoutes.InitPostRoutes(api, DB)
	categoryRoutes.InitCategoryRoutes(api, DB)

	app.Run("8000")
}
