package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func serveAPI() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/menu", getAllMenu)
		api.GET("/menu/today", getTodayMenu)
		api.GET("/menu/weekday/:day", getMenuByWeekday)
		api.GET("/menu/fetch", fetchMenu)
	}

	router.Run("localhost:8000")
}

func getAllMenu(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, loadedMenu)
}

func getMenuByWeekday(c *gin.Context) {
	day, err := strconv.Atoi(c.Param("day"))
	if err != nil || day < 1 || day > 7 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Allowed weekdays is 1->7 for Monday->Sunday"})
		return
	}
	c.IndentedJSON(http.StatusOK, loadedMenu.menuOfWeekday(Weekday(day)))
}

func getTodayMenu(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, loadedMenu.menuOfWeekday(getTodayWeekday()))
}

func fetchMenu(c *gin.Context) {
	ok := updateMenu(loadedMenu)
	if ok {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Fetch new menu succeed"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Fetch new menu failed"})
	}
}
