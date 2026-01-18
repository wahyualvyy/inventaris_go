package routes

import (
	"lab-inventaris/config"
	"lab-inventaris/controllers"
	"lab-inventaris/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	controllers.SeedAdmin()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte("secret123"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/login", controllers.ShowLoginPage)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	authorized := r.Group("/")
	authorized.Use(AuthRequired)
	{
		api := authorized.Group("/api/v1")
		{
			api.POST("/items", controllers.CreateItem)
			api.GET("/labs/:lab_id/items", controllers.GetItemsByLab)
			api.PUT("/items/:id/check", controllers.UpdateItemStatus)
			api.PUT("/items/batch-check", controllers.BatchUpdateItems)
		}

		authorized.GET("/labs/:lab_id/check", func(ctx *gin.Context) {
			labId := ctx.Param("lab_id")

			session := sessions.Default(ctx)
			username := session.Get("username")

			var items []models.Item
			var lab models.Lab

			if err := config.DB.First(&lab, labId).Error; err != nil {
				ctx.String(http.StatusNotFound, "Lab tidak ditemukan")
				return
			}

			config.DB.Where("lab_id = ?", labId).Order("id asc").Find(&items)

			ctx.HTML(http.StatusOK, "lab_check.html", gin.H{
				"Lab":      lab,
				"Items":    items,
				"Username": username,
			})
		})
	}

	return r
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user_id")

	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	c.Next()
}