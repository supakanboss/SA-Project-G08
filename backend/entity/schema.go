package entity

import (
	"time"

	"gorm.io/gorm"
)

type GENDER struct {
	gorm.Model
	Gender_Type string
	Member      []MEMBER `gorm:"foreignKey:GenderID"`
	Staff       []STAFF  `gorm:"foreignKey:GenderID"`
}

type PACKAGE struct {
	gorm.Model
	Package_Type  string
	Package_Price uint
	Member        []MEMBER `gorm:"foreignKey:PackageID"`
}

type PROVINCE struct {
	gorm.Model
	Province_Type string
	Member        []MEMBER `gorm:"foreignKey:ProvinceID"`
}

type MEMBER struct {
	gorm.Model
	Member_Name string
	Email       string `gorm:"uniqueIndex"`
	Password    string
	Age         uint
	Height      float32
	Weight      float32
	Tel         string `gorm:"uniqueIndex"`
	BirthDay    time.Time
	GenderID    *uint
	ProvinceID  *uint
	PackageID   *uint

	// เป็นข้อมูล member ใช้เพื่อให้ join ง่ายขึ้น
	Gender   GENDER   `gorm:"references:id"`
	Province PROVINCE `gorm:"references:id"`
	Package  PACKAGE  `gorm:"references:id"`

	Location_Reservation []LOCATION_RESERVATION `gorm:"foreignKey:MemberID"`
}

type LOCATION struct {
	gorm.Model
	Location_Name        string
	Location_Reservation []LOCATION_RESERVATION `gorm:"foreignKey:LocationID"`
}

type SPORT_TYPE struct {
	gorm.Model
	Sport_Type_Name      string
	Location_Reservation []LOCATION_RESERVATION `gorm:"foreignKey:Sport_TypeID"`
}

type LOCATION_RESERVATION struct {
	gorm.Model
	Time_In  time.Time
	Time_Out time.Time

	//id ทำหน้าที่ เป็น FK
	MemberID     *uint
	LocationID   *uint
	Sport_TypeID *uint

	//เป็น ข้อมูล เพื่อให้จอยง่ายขึ้น
	Member     MEMBER     `gorm:"references:id"`
	Location   LOCATION   `gorm:"references:id"`
	Sport_Type SPORT_TYPE `gorm:"references:id"`

	CheckInOut []CHECK_IN_OUT `gorm:"foreignKey:LocationReservationID"`
}
type STAFF struct {
	gorm.Model
	Staff_BOD time.Time

	Staff_name     string
	Staff_email    string
	Staff_password string

	GenderID *uint

	Staff_address string
	Staff_tel     string

	Gender     GENDER         `gorm:"references:id"`
	CheckInOut []CHECK_IN_OUT `gorm:"foreignKey:StaffID"`
}
type STATUS struct {
	gorm.Model
	Discribe   string
	CheckInOut []CHECK_IN_OUT `gorm:"foreignKey:StatusID"`
}
type CHECK_IN_OUT struct {
	gorm.Model

	StaffID               *uint
	StatusID              *uint
	LocationReservationID *uint

	Status              STATUS               `gorm:"references:id"`
	Staff               STAFF                `gorm:"references:id"`
	LocationReservation LOCATION_RESERVATION `gorm:"references:id"`
}
