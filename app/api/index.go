package api

import (
	. "app-template/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) IndexPage(c *gin.Context) {
	SendJson(c, gin.H{
		"status":  "success",
		"message": "Welcome to app template",
	})
}

func (app *App) HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func (app *App) AddLocation(c *gin.Context) {
	payload := struct{}{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		SendError(c, fmt.Sprintf("Error reading request:%s", err.Error()))
		return
	}
	Log("Inserting a new location...")
	if err := app.DB.Save(&payload).Error; err != nil {
		msg := fmt.Sprintf("Error inserting location: %s", err.Error())
		SendError(c, msg)
		return
	}
	SendJson(c, gin.H{
		"message": "location added successfully",
		"payload": payload,
	})
}
