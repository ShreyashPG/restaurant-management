package main


import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/ShreyashPG/go-restaurant-backend/database"
	routes "github.com/ShreyashPG/go-restaurant-backend/routes"
	middleware "github.com/ShreyashPG/go-restaurant-backend/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.MenuRoutes(router)

	router.Run(":" + port)
}