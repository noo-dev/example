package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	categoryRoutes "github.com/noo-dev/example/internal/category/routes"
	postRoutes "github.com/noo-dev/example/internal/category/routes"
	"github.com/noo-dev/example/pkg/database/dbcon"
)

func main() {
	app := gin.Default()
	DB, err := dbcon.ConnectToDB()
	if err != nil {
		log.Fatal("db error")
	}
	api := app.Group("api")
	postRoutes.InitPostRoutes(api, DB)
	categoryRoutes.InitCategoryRoutes(api, DB)

	app.Run("8000")
}
