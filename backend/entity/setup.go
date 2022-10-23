package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	///////////ทดลองให้เวลาเกิดเป็นเวลาปัจุบัน///////////
	//////////////////////////////////////////////

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&LOCATION{},
		&SPORT_TYPE{},
		&SPORT_EQUIPMENT_TYPE{},
		&REPORT{},
	)

	db = database

	/////////////ข้อมูลใน entity LOCATION///////////
	Location1 := LOCATION{
		Location_Name: "สนามฟุตบอล",
	}
	db.Model(&LOCATION{}).Create(&Location1)

	Location2 := LOCATION{
		Location_Name: "สระว่ายน้ำ",
	}
	db.Model(&LOCATION{}).Create(&Location2)

	Location3 := LOCATION{
		Location_Name: "สนามแบดมินตัน",
	}
	db.Model(&LOCATION{}).Create(&Location3)

	////////////ข้อมูลใน entity SPORT_TYPE/////////
	Sport_type1 := SPORT_TYPE{
		Sport_Type_Name: "ฟุตบอล",
	}
	db.Model(&SPORT_TYPE{}).Create(&Sport_type1)

	Sport_type2 := SPORT_TYPE{
		Sport_Type_Name: "ว่ายน้ำ",
	}
	db.Model(&SPORT_TYPE{}).Create(&Sport_type2)

	Sport_type3 := SPORT_TYPE{
		Sport_Type_Name: "แบดมินตัน",
	}
	db.Model(&SPORT_TYPE{}).Create(&Sport_type3)

	//ข้อมํูลในตาราง sport epuipment tpye
	Sport_Equipment_Type1 := SPORT_EQUIPMENT_TYPE{
		SPORT_EQUIPMENT_TYPE_Name: "รองเท้า",
	}
	db.Model(&SPORT_EQUIPMENT_TYPE{}).Create(&Sport_Equipment_Type1)

	Sport_Equipment_Type2 := SPORT_EQUIPMENT_TYPE{
		SPORT_EQUIPMENT_TYPE_Name: "ลูก",
	}
	db.Model(&SPORT_EQUIPMENT_TYPE{}).Create(&Sport_Equipment_Type2)

	Sport_Equipment_Type3 := SPORT_EQUIPMENT_TYPE{
		SPORT_EQUIPMENT_TYPE_Name: "ไม้",
	}
	db.Model(&SPORT_EQUIPMENT_TYPE{}).Create(&Sport_Equipment_Type3)

	//ข้อมูลตารางreport ตารางหลัก
	db.Model(&REPORT{}).Create(&REPORT{
		Location:             Location3,
		Sport_Equipment_Type: Sport_Equipment_Type3,
		Sport_Type:           Sport_type3,
		Detail:               "หนูตบแรงจนไม้แบดมินตันหัก",
	})

	db.Model(&REPORT{}).Create(&REPORT{
		Location:             Location1,
		Sport_Equipment_Type: Sport_Equipment_Type1,
		Sport_Type:           Sport_type1,
		Detail:               "ผมเตะลูกฟุตบอลแตกครับ",
	})

	db.Model(&REPORT{}).Create(&REPORT{
		Location:             Location3,
		Sport_Equipment_Type: Sport_Equipment_Type2,
		Sport_Type:           Sport_type3,
		Detail:               "ลูกแบดมินตันพังค่ะ",
	})

}
