package main

import (
	"github.com/pradthana7/sa-65-sq/controller"
	"github.com/pradthana7/sa-65-sq/entity"
	"github.com/pradthana7/sa-65-sq/middlewares"

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

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{

			//combobox sport type
			r.GET("/sport_type", controller.ListSport_Type)
			r.GET("/sport_equipment_type", controller.ListSport_Equipment_Type)
			r.GET("/staff/:id", controller.GetStaff)

			// แสดงข้อมูลอุปกรณ์ที่บันทึก
			r.GET("/sport_equipment_data", controller.ListSport_Equipment)

			// sport equipment
			r.POST("/Create_Sports_Equipment", controller.CreateSport_Equipment)
			
		}
	}

	// login
	r.POST("/StaffLogin", controller.StaffLogin)
	r.Run()

}
