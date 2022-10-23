package main

import (
	"github.com/fahwanat/sa-65-demo/controller"
	"github.com/fahwanat/sa-65-demo/entity"

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

	//combobox Location
	r.GET("/gender", controller.ListGender)
	//combobox sport type
	r.GET("/package", controller.ListPackage)
	//combobox  sport equipment type
	r.GET("/province", controller.ListProvine)

	r.GET("/member", controller.ListMember)

	r.POST("/createmember", controller.CreateMember)

	r.Run()

}
