package routes

import (
	"lab-inventaris/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/items", controllers.CreateItem)
		api.GET("/labs/:lab_id/items", controllers.GetItemsByLab)
		api.PUT("/items/:id/check", controllers.UpdateItemStatus)
	}
	return r
}