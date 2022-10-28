package entity

import (
	"time"

	"gorm.io/gorm"
)

// สร้างตาราง gender //
type GENDER struct {
	gorm.Model
	Gender_Type string
	Member      []MEMBER `gorm:"foreignKey:GenderID"`
}

// สร้างตาราง package //
type PACKAGE struct {
	gorm.Model
	Package_Type   string
	Package_Price uint
	Member         []MEMBER `gorm:"foreignKey:PackageID"`
}

// สร้างตาราง province //
type PROVINCE struct {
	gorm.Model
	Province_Type string
	Member        []MEMBER `gorm:"foreignKey:ProvinceID"`
}

// สร้างตาราง member //
type MEMBER struct {
	gorm.Model
	Member_Name	string
	Email       string `gorm:"uniqueIndex"`
	Password    string
	Age         uint
	Height      uint
	Weight      uint
	Tel         string `gorm:"uniqueIndex"`
	BirthDay	time.Time

	// เป็นการไปหาข้อมูลในตารางย่อยต่างๆ //
	GenderID    *uint
	ProvinceID  *uint
	PackageID   *uint

	// เป็นข้อมูล member ใช้เพื่อให้ join ง่ายขึ้น
	Gender   GENDER
	Province PROVINCE
	Package  PACKAGE
}
