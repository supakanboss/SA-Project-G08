package entity

import (
	"time"

	"gorm.io/gorm"
)

type GENDER struct {
	gorm.Model
	Gender_Type string   `gorm:"uniqueIndex"`
	Member      []MEMBER `gorm:"foreignKey:GenderID"`
	Staff       []STAFF  `gorm:"foreignKey:GenderID"`
}

type PACKAGE struct {
	gorm.Model
	Package_Type  string `gorm:"uniqueIndex"`
	Package_Price uint
	Member        []MEMBER `gorm:"foreignKey:PackageID"`
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
	Height      float32
	Weight      float32
	Tel         string `gorm:"uniqueIndex"`
	BirthDay         time.Time

	GenderID   *uint
	ProvinceID *uint
	PackageID  *uint
	// เป็นข้อมูล member ใช้เพื่อให้ join ง่ายขึ้น
	Gender   GENDER
	Province PROVINCE
	Package  PACKAGE

	Borrow_Sport_Equipment []BORROW_SPORT_EQUIPMENT `gorm:"foreignKey:MemberID"`
}
type STAFF struct {
	gorm.Model
	Staff_name      string
	Staff_email     string `gorm:"uniqueIndex"`
	Staff_password  string
	Staff_address   string
	Staff_tel       string  `gorm:"uniqueIndex"`
	Staff_BOD       time.Time
	Sport_Equipment []SPORT_EQUIPMENT `gorm:"foreignKey:StaffID"`
	// GenderID is a fk from GENDER
	GenderID *uint
	Gender   GENDER
}

type SPORT_TYPE struct {
	gorm.Model
	Sport_Type_Name string            `gorm:"uniqueIndex"`
	Sport_Equipment []SPORT_EQUIPMENT `gorm:"foreignKey:Sport_TypeID"`
}

type SPORT_EQUIPMENT_TYPE struct {
	gorm.Model
	SPORT_EQUIPMENT_TYPE_Name string                   `gorm:"uniqueIndex"`
	Sport_Equipment           []SPORT_EQUIPMENT        `gorm:"foreignKey:Sport_Equipment_TypeID"`
	Borrow_Sport_Equipment    []BORROW_SPORT_EQUIPMENT `gorm:"foreignKey:Sport_Equipment_TypeID"`
}

type SPORT_EQUIPMENT struct {
	gorm.Model
	Sport_Equipment_Name   string
	Sport_Equipment_Amount uint
	// StaffID fk จาก staff
	StaffID *uint
	Staff   STAFF

	// SportTypeID fk จาก SportType
	Sport_TypeID *uint
	Sport_type   SPORT_TYPE

	// SportEquipmentTypeID จาก SportEquipmentType
	Sport_Equipment_TypeID *uint
	Sport_equipment_type   SPORT_EQUIPMENT_TYPE
	Borrow_Sport_Equipment []BORROW_SPORT_EQUIPMENT `gorm:"foreignKey:Sport_EquipmentID"`
}

type BORROW_SPORT_EQUIPMENT struct {
	gorm.Model
	
	Time_Borrow     time.Time
	Time_Giveback   time.Time

	MemberID               *uint
	Sport_Equipment_TypeID *uint
	Sport_EquipmentID      *uint

	Member               MEMBER
	Sport_Equipment_Type SPORT_EQUIPMENT_TYPE
	Sport_Equipment      SPORT_EQUIPMENT
}
