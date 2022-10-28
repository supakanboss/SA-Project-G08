package main

import (

  "github.com/bxngearnx/sa-demo/controller"

  "github.com/bxngearnx/sa-demo/entity"

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

	//combobox sport equipment type
	r.GET("/sport_equipment_type", controller.ListSport_Equipment_Type)
	//combobox sport  equipment 
	
	r.GET("/borrow-sport-eqiupments", controller.ListBorrow_Sport_Eqiupment)
	r.POST("/borrow-sport-eqiupment", controller.CreateBorrow_Sport_Eqiupment)
	
	//ให้id
	r.GET("/member/:id", controller.GetMember)
	//login
	r.POST("/Login", controller.Login)
    // แสดงข้อมูลอุปกรณ์ที่บันทึก
    r.GET("/sport_equipment_data", controller.ListSport_Equipment)
	r.Run()

}
