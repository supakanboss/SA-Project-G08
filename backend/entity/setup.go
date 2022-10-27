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
	///////////ทดลองให้เวลาเกิดเป็นเวลาปัจุบัน///////////
	t := time.Now()
	//////////////////////////////////////////////

	database, err := gorm.Open(sqlite.Open("sa-G08.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&GENDER{},
		&PACKAGE{},
		&PROVINCE{},
		&MEMBER{},
		&LOCATION{},
		&SPORT_TYPE{},
		&LOCATION_RESERVATION{},
	)

	db = database

	///////////////ข้อมูลใน entity GENDER///////////
	GenderMale := GENDER{
		Gender_Type: "ชาย",
	}
	db.Model(&GENDER{}).Create(&GenderMale)

	GenderFemale := GENDER{
		Gender_Type: "หญิง",
	}
	db.Model(&GENDER{}).Create(&GenderFemale)

	GenderOther := GENDER{
		Gender_Type: "อื่นๆ",
	}
	db.Model(&GENDER{}).Create(&GenderOther)

	///////////////ข้อมูลใน entity PACKAGE//////////
	Package1 := PACKAGE{
		Package_Type:  "Year (25000 THB)",
		Package_Price: 25000,
	}
	db.Model(&PACKAGE{}).Create(&Package1)

	Package2 := PACKAGE{
		Package_Type:  "Month (2000 THB)",
		Package_Price: 2000,
	}
	db.Model(&PACKAGE{}).Create(&Package2)

	Package3 := PACKAGE{
		Package_Type:  "daily (80 THB)",
		Package_Price: 80,
	}
	db.Model(&PACKAGE{}).Create(&Package3)

	///////////////ข้อมูลใน entity PROVINCE/////////
	Province1 := PROVINCE{
		Province_Type: "กรุงเทพมหานคร",
	}
	db.Model(&PROVINCE{}).Create(&Province1)

	Province2 := PROVINCE{
		Province_Type: "ขอนแก่น",
	}
	db.Model(&PROVINCE{}).Create(&Province2)

	Province3 := PROVINCE{
		Province_Type: "ชลบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province3)

	Province4 := PROVINCE{
		Province_Type: "เชียงใหม่",
	}
	db.Model(&PROVINCE{}).Create(&Province4)

	Province5 := PROVINCE{
		Province_Type: "นครราชสีมา",
	}
	db.Model(&PROVINCE{}).Create(&Province5)

	Province6 := PROVINCE{
		Province_Type: "ปทุมธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province6)

	Province7 := PROVINCE{
		Province_Type: "พระนครศรีอยุธยา",
	}
	db.Model(&PROVINCE{}).Create(&Province7)

	Province8 := PROVINCE{
		Province_Type: "สุรินทร์",
	}
	db.Model(&PROVINCE{}).Create(&Province8)

	Province9 := PROVINCE{
		Province_Type: "อ่างทอง",
	}
	db.Model(&PROVINCE{}).Create(&Province9)

	Province10 := PROVINCE{
		Province_Type: "อุตรดิตถ์",
	}
	db.Model(&PROVINCE{}).Create(&Province10)

	///////////////ข้อมูลใน entity MEMBER///////////
	password1, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	Member1 := MEMBER{
		Member_Name: "สุภานัน เรืองสุข",
		Email:       "supanan@gmail.com",
		Password:    string(password1),
		Age:         20,
		Height:      190.0,
		Weight:      45.0,
		Tel:         "0885870149",
		BirthDay:    t,
		Gender:      GenderFemale,
		Package:     Package2,
		Province:    Province8,
	}
	db.Model(&MEMBER{}).Create(&Member1)

	Member2 := MEMBER{
		Member_Name: "วนัสนันท์ จันทร์มล",
		Email:       "B6306304@gmail.com",
		Password:    string(password2),
		Age:         21,
		Height:      170.0,
		Weight:      48.0,
		Tel:         "0803299545",
		BirthDay:    t,
		Gender:      GenderFemale,
		Package:     Package1,
		Province:    Province4,
	}
	db.Model(&MEMBER{}).Create(&Member2)

	Member3 := MEMBER{
		Member_Name: "ศุภกานต์ แสงจันทร์",
		Email:       "B6311391@gmail.com",
		Password:    string(password3),
		Age:         20,
		Height:      180.0,
		Weight:      69.0,
		Tel:         "0655619892",
		BirthDay:    t,
		Gender:      GenderMale,
		Package:     Package3,
		Province:    Province1,
	}
	db.Model(&MEMBER{}).Create(&Member3)

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

}
