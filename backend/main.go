package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supakanboss/sa-G08/controller"
	"github.com/supakanboss/sa-G08/entity"
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
	//ให้id
	r.GET("/member/:id", controller.GetMember)
	//แสดงข้อมูลการจอง
	r.GET("/Location_Reservations", controller.ListLocation_Reservation)

	//login
	r.POST("/Login", controller.Login)

	//รับค่าเข้าตาราง
	r.POST("/Location_Reservation", controller.CreateLocation_Reservation)

	r.Run()

}
