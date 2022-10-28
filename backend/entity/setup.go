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
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&CHECK_IN_OUT{},
		&GENDER{},
		&PACKAGE{},
		&PROVINCE{},
		&MEMBER{},
		&LOCATION{},
		&SPORT_TYPE{},
		&LOCATION_RESERVATION{},
		&STAFF{},
		&STATUS{},
	)

	db = database

	///////////////ข้อมูลใน entity GENDER//////////
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
		Package_Type:  "year",
		Package_Price: 25000,
	}
	db.Model(&PACKAGE{}).Create(&Package1)

	Package2 := PACKAGE{
		Package_Type:  "month",
		Package_Price: 2000,
	}
	db.Model(&PACKAGE{}).Create(&Package2)

	Package3 := PACKAGE{
		Package_Type:  "daily",
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

	Member1 := MEMBER{
		Member_Name: "สุภานัน เรืองสุข",
		Email:       "supanan@gmail.com",
		Password:    "12348888",
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
		Password:    "55555555",
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
		Password:    "12345678",
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

	/////ข้อมูลใน entity LOCATION_RESERVATION/////
	LocationR1 := LOCATION_RESERVATION{
		Member:     Member1,
		Location:   Location3,
		Sport_Type: Sport_type3,
		Time_In:    t,
		Time_Out:   t,
	}
	db.Model(&LOCATION_RESERVATION{}).Create(&LocationR1)

	LocationR2 := LOCATION_RESERVATION{
		Member:     Member2,
		Location:   Location1,
		Sport_Type: Sport_type1,
		Time_In:    t,
		Time_Out:   t,
	}
	db.Model(&LOCATION_RESERVATION{}).Create(&LocationR2)

	LocationR3 := LOCATION_RESERVATION{
		Member:     Member3,
		Location:   Location3,
		Sport_Type: Sport_type3,
		Time_In:    t,
		Time_Out:   t,
	}
	db.Model(&LOCATION_RESERVATION{}).Create(&LocationR3)

	///////////////ข้อมูลใน entity STATUS///////////
	Status1 := STATUS{
		Discribe: "Check-in",
	}
	db.Model(&STATUS{}).Create(&Status1)
	Status2 := STATUS{
		Discribe: "Check-out",
	}
	db.Model(&STATUS{}).Create(&Status2)

	///////////////ข้อมูลใน entity STAFF///////////
	password4, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password5, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password6, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)

	Staff1 := STAFF{
		Staff_name:     "นฤมล เกษรบัว",
		Staff_email:    "Kaenniza55@gmail.com",
		Staff_password: string(password4),
		Gender:         GenderFemale,
		Staff_address:  "114 ม.10 ต.ในเมือง จ.สุรินทร์",
		Staff_tel:      "0954101589",
		Staff_BOD: t,
	}
	db.Model(&STAFF{}).Create(&Staff1)

	Staff2 := STAFF{
		Staff_name:     "ใจดี สมใจ",
		Staff_email:    "Jaideehihi@gmail.com",
		Staff_password: string(password5),
		Gender:         GenderFemale,
		Staff_address:  "555 ม.1 ต.ในเมือง จ.ขอนแก่น",
		Staff_tel:      "0973340404",
		Staff_BOD: t,
	}
	db.Model(&STAFF{}).Create(&Staff2)

	Staff3 := STAFF{
		Staff_name:     "บุญมี กรรมบัง",
		Staff_email:    "Boomee@gmail.com",
		Staff_password: string(password6),
		Gender:         GenderMale,
		Staff_address:  "561 ม.1 ต.ในเมือง จ.ขอนแก่น",
		Staff_tel:      "0983380848",
		Staff_BOD: t,
	}
	db.Model(&STAFF{}).Create(&Staff3)

	/////ข้อมูลใน entity CHECK_IN_OUT/////
	db.Model(&CHECK_IN_OUT{}).Create(&CHECK_IN_OUT{
		Staff:               Staff2,
		Status:              Status1,
		LocationReservation: LocationR3,
	})
	db.Model(&CHECK_IN_OUT{}).Create(&CHECK_IN_OUT{
		Staff:               Staff1,
		Status:              Status2,
		LocationReservation: LocationR3,
	})
	db.Model(&CHECK_IN_OUT{}).Create(&CHECK_IN_OUT{
		Staff:               Staff3,
		Status:              Status1,
		LocationReservation: LocationR1,
	})
}
