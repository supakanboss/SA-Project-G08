package entity

import (
	"time"

	"gorm.io/gorm"
)

type GENDER struct {
	gorm.Model
	Gender_Type string
	// 1 gender มีได้หลาย staff
	STAFF []STAFF `gorm:"foreignKey:GenderID"`
}

type STAFF struct {
	gorm.Model
	Staff_name     string
	Staff_email    string
	Staff_password string
	Staff_address  string
	Staff_tel      string
	Staff_BOD      time.Time

	// GenderID is a fk from GENDER
	GenderID *uint
	Gender   GENDER

	// 1 staff มีได้หลาย sport equipment
	Sport_Equipment []SPORT_EQUIPMENT `gorm:"foreignKey:StaffID"`
}

type SPORT_TYPE struct {
	gorm.Model
	Sport_Type_Name string

	// 1 sportType มีได้หลาย sport equipment
	Sport_Equipment []SPORT_EQUIPMENT `gorm:"foreignKey:SportTypeID"`
}

type SPORT_EQUIPMENT_TYPE struct {
	gorm.Model
	SPORT_EQUIPMENT_TYPE_Name string

	// 1 SportEquipmentType มีได้หลาย sport equipment
	Sport_Equipment []SPORT_EQUIPMENT `gorm:"foreignKey:SportEquipmentTypeID"`
}

type SPORT_EQUIPMENT struct {
	gorm.Model
	Sport_Equipment_Name   string
	Sport_Equipment_Amount uint
	// StaffID fk จาก staff
	StaffID *uint
	Staff   STAFF

	// SportTypeID fk จาก SportType
	SportTypeID *uint
	Sport_type  SPORT_TYPE

	// SportEquipmentTypeID จาก SportEquipmentType
	SportEquipmentTypeID *uint
	Sport_equipment_type SPORT_EQUIPMENT_TYPE

}

