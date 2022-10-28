package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
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
		&MEMBER{},
		&PAYMENT_TYPE{},
		&PACKAGE{},
		&BANK{},
		&PAYMENT{},
	)

	db = database

	// GENDER
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
	//PACKAGE
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
	//PROVINCE
	province1 := PROVINCE{
		Province_Type: "กรุงเทพมหานคร",
	}
	db.Model(&PROVINCE{}).Create(&province1)

	province2 := PROVINCE{
		Province_Type: "เชียงใหม่",
	}
	db.Model(&PROVINCE{}).Create(&province2)

	province3 := PROVINCE{
		Province_Type: "สุรินทร์",
	}
	db.Model(&PROVINCE{}).Create(&province3)

	province4 := PROVINCE{
		Province_Type: "ภูเก็ต",
	}
	db.Model(&PROVINCE{}).Create(&province4)

	province5 := PROVINCE{
		Province_Type: "ระยอง",
	}
	db.Model(&PROVINCE{}).Create(&province5)

	///////////////ข้อมูลใน entity MEMBER///////////
	password1, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	//คนที่1
	db.Model(&MEMBER{}).Create(&MEMBER{
		Member_Name: "สุภานัน เรืองสุข",
		Email:       "supanan@gmail.com",
		Password:    string(password1),
		Age:         20,
		Height:      190,
		Weight:      45,
		Tel:         "0885870149",
		BirthDay:    t,
		Gender:      GenderMale,
		Province:    province3,
		Package:     Package2,
	})
	//คนที่2
	db.Model(&MEMBER{}).Create(&MEMBER{
		Member_Name: "วนัสนันท์ จันทร์มล",
		Email:       "B6306304@g.sut.ac.th",
		Password:    string(password2),
		Age:         21,
		Height:      170,
		Weight:      48,
		Tel:         "0803299545",
		BirthDay:    t,
		Gender:      GenderFemale,
		Province:    province2,
		Package:     Package1,
	})
	//คนที่3
	db.Model(&MEMBER{}).Create(&MEMBER{
		Member_Name: "ศุภกานต์ แสงจันทร์",
		Email:       "B6311391@g.sut.ac.th",
		Password:    string(password3),
		Age:         20,
		Height:      180,
		Weight:      66,
		Tel:         "0655619892",
		BirthDay:    t,
		Gender:      GenderMale,
		Province:    province1,
		Package:     Package3,
	})

	//PAYMENT_TYPE
	payment_type1 := PAYMENT_TYPE{
		Payment_Type_Name: "Mobile banking",
	}
	db.Model(&PAYMENT_TYPE{}).Create(&payment_type1)

	//BANK
	bank1 := BANK{
		Bank_Name: "ธนาคารไทยพาณิชย์",
	}
	db.Model(&BANK{}).Create(&bank1)

	bank2 := BANK{
		Bank_Name: "ธนาคารกสิกรไทย",
	}
	db.Model(&BANK{}).Create(&bank2)

	bank3 := BANK{
		Bank_Name: "ธนาคารกรุงไทย",
	}
	db.Model(&BANK{}).Create(&bank3)

	bank4 := BANK{
		Bank_Name: "ธนาคารกรุุงเทพ",
	}
	db.Model(&BANK{}).Create(&bank4)

	bank5 := BANK{
		Bank_Name: "ธนาคารทีทีบี (TTB)",
	}
	db.Model(&BANK{}).Create(&bank5)

	bank6 := BANK{
		Bank_Name: "ธนาคารออมสิน",
	}
	db.Model(&BANK{}).Create(&bank6)

	bank7 := BANK{
		Bank_Name: "ธนาคารกรุงศรี",
	}
	db.Model(&BANK{}).Create(&bank7)

	bank8 := BANK{
		Bank_Name: "ธ.ก.ส",
	}
	db.Model(&BANK{}).Create(&bank8)

	bank9 := BANK{
		Bank_Name: "ธนาคารยูบีโอ (UBO)",
	}
	db.Model(&BANK{}).Create(&bank9)

	bank10 := BANK{
		Bank_Name: "ธนาคารอาคารสงเคราะห์",
	}
	db.Model(&BANK{}).Create(&bank10)

	bank11 := BANK{
		Bank_Name: "ธนาคารซีไอเอ็มบี (ICBM)",
	}
	db.Model(&BANK{}).Create(&bank11)

	bank12 := BANK{
		Bank_Name: "ซิตี้แบงก์",
	}
	db.Model(&BANK{}).Create(&bank12)

	bank13 := BANK{
		Bank_Name: "ดอยซ์แบงก์",
	}
	db.Model(&BANK{}).Create(&bank13)

	bank14 := BANK{
		Bank_Name: "ธนาคารเอชเอสบีซี (HSBC)",
	}
	db.Model(&BANK{}).Create(&bank14)

	bank15 := BANK{
		Bank_Name: "ธนาคารไอซีบีซี (ICBC)",
	}
	db.Model(&BANK{}).Create(&bank15)

	bank16 := BANK{
		Bank_Name: "ธนาคารอิสลาม",
	}
	db.Model(&BANK{}).Create(&bank16)

	bank17 := BANK{
		Bank_Name: "ธนาคารเกียรตินาคินภัทร",
	}
	db.Model(&BANK{}).Create(&bank17)

	bank18 := BANK{
		Bank_Name: "ธนาคารแลนด์ แอนด์ เฮ้าส์ (LH)",
	}
	db.Model(&BANK{}).Create(&bank18)

	bank19 := BANK{
		Bank_Name: "ธนาคารมิซูโฮ (MIZUHO)",
	}
	db.Model(&BANK{}).Create(&bank19)

	bank20 := BANK{
		Bank_Name: "ธนาคารสแตนดาร์ดชาร์เตอร์ด (standardchartered)",
	}
	db.Model(&BANK{}).Create(&bank20)

	bank21 := BANK{
		Bank_Name: "ธนาคารซูมิโตโม (SMBC)",
	}
	db.Model(&BANK{}).Create(&bank21)

	bank22 := BANK{
		Bank_Name: "ธนาคารไทยเครดิต",
	}
	db.Model(&BANK{}).Create(&bank22)

	bank23 := BANK{
		Bank_Name: "ธนาคารทิสโก้ (TISCO)",
	}
	db.Model(&BANK{}).Create(&bank23)

	//PAYMENT (main entity)
	//คนที่ 1
	// db.Model(&PAYMENT{}).Create(&PAYMENT{
	// 	Member:       Member1,
	// 	Payment_Type: payment_type1,
	// 	Package:      package2,
	// 	Datetime:     t,
	// 	Price:        2000,
	// 	Bank:         bank3,
	// 	Bank_Account: "832-23430-21",
	// })

	// //คนที่ 2
	// db.Model(&PAYMENT{}).Create(&PAYMENT{
	// 	Member:       Member2,
	// 	Payment_Type: payment_type1,
	// 	Package:      package1,
	// 	Datetime:     t,
	// 	Price:        25000,
	// 	Bank:         bank1,
	// 	Bank_Account: "546-21345-09",
	// })

	// //คนที่ 3
	// db.Model(&PAYMENT{}).Create(&PAYMENT{
	// 	Member:       Member3,
	// 	Payment_Type: payment_type1,
	// 	Package:      package3,
	// 	Datetime:     t,
	// 	Price:        80,
	// 	Bank:         bank2,
	// 	Bank_Account: "121-23880-87",
	// })
}
