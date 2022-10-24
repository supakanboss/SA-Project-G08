package entity

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	t := time.Now()
	database, err := gorm.Open(sqlite.Open("sa-65-SportEquipment.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&GENDER{},
		&STAFF{},
		&SPORT_TYPE{},
		&SPORT_EQUIPMENT_TYPE{},
		&SPORT_EQUIPMENT{},
	)

	db = database

	// GENDER
	genM := GENDER{
		Gender_Type: "ชาย",
	}
	db.Model(&GENDER{}).Create(&genM)

	genF := GENDER{
		Gender_Type: "หญิง",
	}
	db.Model(&GENDER{}).Create(&genF)

	genO := GENDER{
		Gender_Type: "อื่นๆ",
	}
	db.Model(&GENDER{}).Create(&genO)

	// STAFF
	password1, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("00112233"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("888888888"), 14)

	staff1 := STAFF{
		Staff_name:     "นฤมน เกษรบัว",
		Staff_email:    "Kaenniza55@gmail.com",
		Staff_password: string(password1),
		Staff_address:  "114 ม.10 ต.ในเมือง จ.สุรินทร์",
		Staff_tel:      "0954101589",
		Staff_BOD:      t,
		Gender:         genF,
	}
	db.Model(&STAFF{}).Create(&staff1)

	staff2 := STAFF{
		Staff_name:     "ใจดี สมใจ",
		Staff_email:    "Jaideehihi@gmail.com",
		Staff_password: string(password2),
		Staff_address:  "555 ม.1 ต.ในเมือง จ.ขอนแก่น",
		Staff_tel:      "0973340404",
		Staff_BOD:      t,
		Gender:         genM,
	}
	db.Model(&STAFF{}).Create(&staff2)

	staff3 := STAFF{
		Staff_name:     "บุญมี กรรมบัง",
		Staff_email:    "Boomee@gmail.com",
		Staff_password: string(password3),
		Staff_address:  "561 ม.1 ต.ในเมือง จ.ขอนแก่น",
		Staff_tel:      "0983380848",
		Staff_BOD:      t,
		Gender:         genO,
	}
	db.Model(&STAFF{}).Create(&staff3)

	// SPORT_TYPE
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

	// SPORT_EQUIPMENT_TYPE
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
