package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/ShreyashPG/go-restaurant-backend/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("menus/:menu_id", controller.UpdateMenu())
}