package entity

import (
	"time"

	"gorm.io/gorm"
)

// /ในส่วนของmember////////////
type GENDER struct {
	gorm.Model
	Gender_Type string   `gorm:"uniqueIndex"`
	Member      []MEMBER `gorm:"foreignKey:GenderID"`
}

type PACKAGE struct {
	gorm.Model
	Package_Type  string `gorm:"uniqueIndex"`
	Package_Price uint
	Member        []MEMBER  `gorm:"foreignKey:PackageID"`
	Payment       []PAYMENT `gorm:"foreignKey:PackageID"`
}

type PROVINCE struct {
	gorm.Model
	Province_Type string   `gorm:"uniqueIndex"`
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
	Tel         string `gorm:"uniqueIndex"`
	BirthDay    time.Time
	GenderID    *uint
	ProvinceID  *uint
	PackageID   *uint
	// เป็นข้อมูล member ใช้เพื่อให้ join ง่ายขึ้น
	Gender   GENDER
	Province PROVINCE
	Package  PACKAGE
	Payment  []PAYMENT `gorm:"foreignKey:MemberID"`
}

/////////////member////////////////

type PAYMENT_TYPE struct {
	gorm.Model
	Payment_Type_Name string    `gorm:"uniqueIndex"`
	Payment           []PAYMENT `gorm:"foreignKey:Payment_TypeID"`
}

type BANK struct {
	gorm.Model
	Bank_Name string    `gorm:"uniqueIndex"`
	Payment   []PAYMENT `gorm:"foreignKey:BankID"`
}

type PAYMENT struct {
	gorm.Model
	MemberID       *uint
	Payment_TypeID *uint
	PackageID      *uint
	BankID         *uint
	////เป็นข้อมูล member ใช้เพื่อให้ join ง่ายขึ้น////////////
	Member       MEMBER
	Payment_Type PAYMENT_TYPE
	Package      PACKAGE
	Bank         BANK
	/////////////////////////////////////////////////
	Datetime     time.Time
	Price        uint
	Bank_Account string `gorm:"uniqueIndex"`
}
