package entity

import (
	"time"

	"gorm.io/gorm"
)

type GENDER struct {
	gorm.Model
	Gender_Type string
	Member      []MEMBER `gorm:"foreignKey:GenderID"`
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
	Height      uint
	Weight      uint
	Tel         string
	BirthDay    time.Time

	Loaction_Reservation []LOCATION_RESERVATION `gorm:"foreignKey:MemberID"`
	//id ทำหน้าที่ เป็น FK
	GenderID   *uint
	PackageID  *uint
	ProvinceID *uint
	//เป็น ข้อมูล เพื่อให้จอยง่ายขึ้น
	Gender   GENDER
	Package  PACKAGE
	Province PROVINCE
}

type LOCATION struct {
	gorm.Model
	Location_Name        string
	Loaction_Reservation []LOCATION_RESERVATION `gorm:"foreignKey:LocationID"`
}

type SPORT_TYPE struct {
	gorm.Model
	Sport_Type_Name      string
	Loaction_Reservation []LOCATION_RESERVATION `gorm:"foreignKey:Sport_TypeID"`
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
	Member     MEMBER
	Location   LOCATION
	Sport_Type SPORT_TYPE
}
