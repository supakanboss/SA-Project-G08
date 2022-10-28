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
	r.GET("/staff/:id", controller.GetStaff)
	//แสดงข้อมูลการจอง
	r.GET("/Location_Reservations", controller.ListLocation_Reservation)

	//รับค่าเข้าตาราง
	r.POST("/Location_Reservation", controller.CreateLocation_Reservation)

	//combobox Location
	r.GET("/gender", controller.ListGender)
	//combobox sport type
	r.GET("/package", controller.ListPackage)
	//combobox  sport equipment type
	r.GET("/province", controller.ListProvine)
	r.GET("/member", controller.ListMember)
	r.POST("/createmember", controller.CreateMembers)

	//login
	r.POST("/Login", controller.Login)
	r.POST("/StaffLogin", controller.StaffLogin)

	//combobox  sport equipment type
	r.GET("/sport_equipment_type", controller.ListSport_Equipment_Type)
	r.POST("/report", controller.CreateReport)

	// แสดงข้อมูลอุปกรณ์ที่บันทึก
	r.GET("/sport_equipment_data", controller.ListSport_Equipment)
	// sport equipment
	r.POST("/Create_Sports_Equipment", controller.CreateSport_Equipment)

	r.GET("/borrow-sport-eqiupments", controller.ListBorrow_Sport_Eqiupment)
	r.POST("/borrow-sport-eqiupment", controller.CreateBorrow_Sport_Eqiupment)

	//combobox payment_type
	r.GET("/payment_type", controller.ListPayment_Type)
	//combobox bank
	r.GET("/bank", controller.ListBank)
	///////////////////////////////////////////////////
	r.POST("/payment", controller.CreatePayment)

	r.GET("/statuscheck", controller.ListStatus)

	r.GET("/locationReservation/", controller.ListLocationReservation)
	r.GET("/locationReservation/:id", controller.GetLocationReservation)
	r.GET("/locationReservationField/:id", controller.GetLocationReservationIDEA)

	//GET รับค่ามาแสดง POST ใช้ในการสร้าง (CREATE) จะมีการรับส่งค่า
	r.GET("/ciocheck", controller.ListCIO)
	r.GET("/cioChStatus/:id", controller.CheckDoubleCIO)
	r.POST("/createcio", controller.CreateCheckInOut)

	r.Run()
}
