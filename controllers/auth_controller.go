package controllers

import (
	"lab-inventaris/config"
	"lab-inventaris/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {
	var user models.User

	if err := config.DB.First(&user).Error; err != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := models.User{
		Username: "admin",
		Password: string(hashedPassword),
		Role: "Super Admin",
	}
	config.DB.Create(&admin)
	}
}

func ShowLoginPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Login(c *gin.Context)	{
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User

	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error" : "Username tidak ditemukan"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error" : "Password salah"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("username", user.Username)
	session.Save()

	c.Redirect(http.StatusFound, "labs/1/check")
}

func Logout(c *gin.Context)	{
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}