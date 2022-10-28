package main

import (
	"github.com/miaw88/sa-65-MEOW/controller"

	"github.com/miaw88/sa-65-MEOW/entity"

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

	//combobox package
	r.GET("/package", controller.ListPackage)
	//combobox payment_type
	r.GET("/payment_type", controller.ListPayment_Type)
	//combobox bank
	r.GET("/bank", controller.ListBank)
	///////////////////////////////////////////////////
	r.GET("/member/:id", controller.GetMember)
	r.POST("/payment", controller.CreatePayment)
	r.POST("/Login", controller.Login)

	r.Run()

}
