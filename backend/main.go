package main

import (
	"github.com/idearslove/sa-65-demo/controller"
	"github.com/idearslove/sa-65-demo/entity"

	"github.com/gin-gonic/gin"
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

	r.GET("/statuscheck", controller.ListStatus)

	r.GET("/locationReservation/", controller.ListLocationReservation)
	r.GET("/locationReservation/:id", controller.GetLocationReservation)
	r.GET("/locationReservationField/:id", controller.GetLocationReservationIDEA)

	r.POST("/StaffLogin", controller.StaffLogin)
	r.GET("/staff/:id", controller.GetStaff)

	//GET รับค่ามาแสดง POST ใช้ในการสร้าง (CREATE) จะมีการรับส่งค่า
	r.GET("/ciocheck", controller.ListCIO)
	r.GET("/cioChStatus/:id", controller.CheckDoubleCIO)
	r.POST("/createcio", controller.CreateCheckInOut)

	r.Run()
}
