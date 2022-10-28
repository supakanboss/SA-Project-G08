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
		Location_Name: "สนาม ฟุตบอล 1",
	}
	db.Model(&LOCATION{}).Create(&Location1)

	Location2 := LOCATION{
		Location_Name: "สนาม ฟุตบอล 2",
	}
	db.Model(&LOCATION{}).Create(&Location2)

	Location3 := LOCATION{
		Location_Name: "สระว่ายน้ำ",
	}
	db.Model(&LOCATION{}).Create(&Location3)

	Location4 := LOCATION{
		Location_Name: "สนาม เทนนิส 1",
	}
	db.Model(&LOCATION{}).Create(&Location4)

	Location5 := LOCATION{
		Location_Name: "สนาม เทนนิส 2",
	}
	db.Model(&LOCATION{}).Create(&Location5)

	Location6 := LOCATION{
		Location_Name: "สนาม วอลเลย์บอล 1",
	}
	db.Model(&LOCATION{}).Create(&Location6)

	Location7 := LOCATION{
		Location_Name: "สนาม วอลเลย์บอล 2",
	}
	db.Model(&LOCATION{}).Create(&Location7)

	Location8 := LOCATION{
		Location_Name: "สนาม เปตอง 1",
	}
	db.Model(&LOCATION{}).Create(&Location8)

	Location9 := LOCATION{
		Location_Name: "สนาม เปตอง 2",
	}
	db.Model(&LOCATION{}).Create(&Location9)

	Location10 := LOCATION{
		Location_Name: "สนาม บาสเกตบอล 1",
	}
	db.Model(&LOCATION{}).Create(&Location10)

	Location11 := LOCATION{
		Location_Name: "สนาม บาสเกตบอล 2",
	}
	db.Model(&LOCATION{}).Create(&Location11)

	Location12 := LOCATION{
		Location_Name: "สนาม มวย",
	}
	db.Model(&LOCATION{}).Create(&Location12)

	Location13 := LOCATION{
		Location_Name: "สนาม ฟุตซอล 1",
	}
	db.Model(&LOCATION{}).Create(&Location13)

	Location14 := LOCATION{
		Location_Name: "สนาม ฟุตซอล 2",
	}
	db.Model(&LOCATION{}).Create(&Location14)

	Location15 := LOCATION{
		Location_Name: "โรงยิม (คอร์ดแบดมินตัน1)",
	}
	db.Model(&LOCATION{}).Create(&Location15)

	Location16 := LOCATION{
		Location_Name: "โรงยิม (คอร์ดแบดมินตัน2)",
	}
	db.Model(&LOCATION{}).Create(&Location16)

	Location17 := LOCATION{
		Location_Name: "โรงยิม (คอร์ดแบดมินตัน3)",
	}
	db.Model(&LOCATION{}).Create(&Location17)

	Location18 := LOCATION{
		Location_Name: "โรงยิม (คอร์ดแบดมินตัน4)",
	}
	db.Model(&LOCATION{}).Create(&Location18)

	Location19 := LOCATION{
		Location_Name: "โรงยิม (คอร์ดแบดมินตัน5)",
	}
	db.Model(&LOCATION{}).Create(&Location19)

	Location20 := LOCATION{
		Location_Name: "โรงยิม (คอร์ดแบดมินตัน6)",
	}
	db.Model(&LOCATION{}).Create(&Location20)

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
		Sport_Type_Name: "เทนนิส",
	}
	db.Model(&SPORT_TYPE{}).Create(&Sport_type3)

	Sport_type4 := SPORT_TYPE{
		Sport_Type_Name: "วอลเลย์บอล",
	}
	db.Model(&SPORT_TYPE{}).Create(&Sport_type4)

	Sport_type5 := SPORT_TYPE{
		Sport_Type_Name: "เปตอง",
	}
	db.Model(&SPORT_TYPE{}).Create(&Sport_type5)

	Sport_type6 := SPORT_TYPE{
		Sport_Type_Name: "บาสเกตบอล",
	}
	db.Model(&SPORT_TYPE{}).Create(&Sport_type6)

	Sport_type7 := SPORT_TYPE{
		Sport_Type_Name: "มวย",
	}
	db.Model(&SPORT_TYPE{}).Create(&Sport_type7)

	Sport_type8 := SPORT_TYPE{
		Sport_Type_Name: "ฟุตซอล",
	}
	db.Model(&SPORT_TYPE{}).Create(&Sport_type8)

	////////ข้อมํูลในตาราง sport epuipment type///////////
	Sport_Equipment_Type1 := SPORT_EQUIPMENT_TYPE{
		SPORT_EQUIPMENT_TYPE_Name: "ประเภท รองเท้า",
	}
	db.Model(&SPORT_EQUIPMENT_TYPE{}).Create(&Sport_Equipment_Type1)

	Sport_Equipment_Type2 := SPORT_EQUIPMENT_TYPE{
		SPORT_EQUIPMENT_TYPE_Name: "ประเภท ไม้",
	}
	db.Model(&SPORT_EQUIPMENT_TYPE{}).Create(&Sport_Equipment_Type2)

	Sport_Equipment_Type3 := SPORT_EQUIPMENT_TYPE{
		SPORT_EQUIPMENT_TYPE_Name: "ประเภท ลูก",
	}
	db.Model(&SPORT_EQUIPMENT_TYPE{}).Create(&Sport_Equipment_Type3)

	Sport_Equipment_Type4 := SPORT_EQUIPMENT_TYPE{
		SPORT_EQUIPMENT_TYPE_Name: "เสื้อวอร์ม",
	}
	db.Model(&SPORT_EQUIPMENT_TYPE{}).Create(&Sport_Equipment_Type4)

	Sport_Equipment_Type5 := SPORT_EQUIPMENT_TYPE{
		SPORT_EQUIPMENT_TYPE_Name: "อุปกรณ์อื่นๆ",
	}
	db.Model(&SPORT_EQUIPMENT_TYPE{}).Create(&Sport_Equipment_Type5)

}
