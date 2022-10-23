package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cutekaennie/sa-65-kn/controller"
	"github.com/cutekaennie/sa-65-kn/entity"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	//combobox Location
	r.GET("/location", controller.ListLocation)
	//combobox sport type
	r.GET("/sport_type", controller.ListSport_Type)
	//combobox  sport equipment type
	r.GET("/sport_equipment_type", controller.ListSport_Equipment_Type)
	r.POST("/report", controller.CreateReport)
	
	r.Run()	
}
