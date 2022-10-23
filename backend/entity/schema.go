package entity

import (
	"gorm.io/gorm"
)

type LOCATION struct {
	gorm.Model
	Location_Name string
	Report        []REPORT `gorm:"foreignKey:LocationID"` // เชื่อมตาราง
}

type SPORT_TYPE struct {
	gorm.Model
	Sport_Type_Name string
	Report          []REPORT `gorm:"foreignKey:Sport_TypeID"` // เชื่อมตาราง
}

type SPORT_EQUIPMENT_TYPE struct {
	gorm.Model
	SPORT_EQUIPMENT_TYPE_Name string
	Report                    []REPORT `gorm:"foreignKey:Sport_Equipment_TypeID"` // เชื่อมตาราง
}

type REPORT struct {
	gorm.Model
	Detail string

	LocationID *uint
	Location   LOCATION //จะได้ join ง่ายขึ้น

	Sport_Equipment_TypeID *uint
	Sport_Equipment_Type   SPORT_EQUIPMENT_TYPE //จะได้ join ง่ายขึ้น

	Sport_TypeID *uint
	Sport_Type   SPORT_TYPE
}
