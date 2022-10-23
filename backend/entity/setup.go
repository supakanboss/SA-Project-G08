package entity

import (
	"time"

	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {
	///////////ทดลองให้เวลาเกิดเป็นเวลาปัจุบัน///////////
	t := time.Now()

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema
	database.AutoMigrate(
		&MEMBER{},
		&GENDER{},
		&PACKAGE{},
		&PROVINCE{},
	)

	db = database

	// GENDER --> เซ็ตข้อมูลเพศ //
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

	//PACKAGE --> เซ็ตข้อมูล package //
	package1 := PACKAGE{
		Package_Type:   "รายปี",
		PackageE_Price: 25000,
	}
	db.Model(&PACKAGE{}).Create(&package1)

	package2 := PACKAGE{
		Package_Type:   "รายเดือน",
		PackageE_Price: 2000,
	}
	db.Model(&PACKAGE{}).Create(&package2)

	package3 := PACKAGE{
		Package_Type:   "รายครั้ง",
		PackageE_Price: 80,
	}
	db.Model(&PACKAGE{}).Create(&package3)

	//PROVINCE --> เซ็ตข้อมูลจังหวัด //
	province1 := PROVINCE{
		Province_Type: "กระบี่",
	}
	db.Model(&PROVINCE{}).Create(&province1)

	province2 := PROVINCE{
		Province_Type: "กรุงเทพมหานคร",
	}
	db.Model(&PROVINCE{}).Create(&province2)

	province3 := PROVINCE{
		Province_Type: "กาญจนบุรี",
	}
	db.Model(&PROVINCE{}).Create(&province3)

	province4 := PROVINCE{
		Province_Type: "กาฬสินธุ์",
	}
	db.Model(&PROVINCE{}).Create(&province4)

	province5 := PROVINCE{
		Province_Type: "กำแพงเพชร",
	}
	db.Model(&PROVINCE{}).Create(&province5)

	province6 := PROVINCE{
		Province_Type: "ขอนแก่น",
	}
	db.Model(&PROVINCE{}).Create(&province6)

	province7:= PROVINCE{
		Province_Type: "จันทบุรี",
	}
	db.Model(&PROVINCE{}).Create(&province7)

	province8 := PROVINCE{
		Province_Type: "ฉะเชิงเทรา",
	}
	db.Model(&PROVINCE{}).Create(&province8)

	province9 := PROVINCE{
		Province_Type: "ชลบุรี",
	}
	db.Model(&PROVINCE{}).Create(&province9)

	province10 := PROVINCE{
		Province_Type: "ชัยนาท",
	}
	db.Model(&PROVINCE{}).Create(&province10)

	province11 := PROVINCE{
		Province_Type: "ชัยภูมิ",
	}
	db.Model(&PROVINCE{}).Create(&province11)

	province12 := PROVINCE{
		Province_Type: "ชุมพร",
	}
	db.Model(&PROVINCE{}).Create(&province12)

	province13 := PROVINCE{
		Province_Type: "เชียงราย",
	}
	db.Model(&PROVINCE{}).Create(&province13)

	province14 := PROVINCE{
		Province_Type: "เชียงใหม่",
	}
	db.Model(&PROVINCE{}).Create(&province14)

	province15 := PROVINCE{
		Province_Type: "ตรัง",
	}
	db.Model(&PROVINCE{}).Create(&province15)

	province16 := PROVINCE{
		Province_Type: "ตราด",
	}
	db.Model(&PROVINCE{}).Create(&province16)

	province17:= PROVINCE{
		Province_Type: "ตาก",
	}
	db.Model(&PROVINCE{}).Create(&province17)

	province18 := PROVINCE{
		Province_Type: "นครนายก",
	}
	db.Model(&PROVINCE{}).Create(&province18)

	province19 := PROVINCE{
		Province_Type: "นครปฐม",
	}
	db.Model(&PROVINCE{}).Create(&province19)

	province20 := PROVINCE{
		Province_Type: "นครพนม",
	}
	db.Model(&PROVINCE{}).Create(&province20)

	province21 := PROVINCE{
		Province_Type: "นครราชสีมา",
	}
	db.Model(&PROVINCE{}).Create(&province21)

	province22 := PROVINCE{
		Province_Type: "นครศรีธรรมราช",
	}
	db.Model(&PROVINCE{}).Create(&province22)

	province23 := PROVINCE{
		Province_Type: "นครสวรรค์",
	}
	db.Model(&PROVINCE{}).Create(&province23)

	province24 := PROVINCE{
		Province_Type: "นนทบุรี",
	}
	db.Model(&PROVINCE{}).Create(&province24)

	province25 := PROVINCE{
		Province_Type: "นราธิวาส",
	}
	db.Model(&PROVINCE{}).Create(&province25)

	province26 := PROVINCE{
		Province_Type: "น่าน",
	}
	db.Model(&PROVINCE{}).Create(&province26)

	province27:= PROVINCE{
		Province_Type: "บึงกาฬ",
	}
	db.Model(&PROVINCE{}).Create(&province27)

	province28 := PROVINCE{
		Province_Type: "บุรีรัมย์",
	}
	db.Model(&PROVINCE{}).Create(&province28)

	province29 := PROVINCE{
		Province_Type: "ปทุมธานี",
	}
	db.Model(&PROVINCE{}).Create(&province29)

	province30 := PROVINCE{
		Province_Type: "ประจวบคีรีขันธ์",
	}
	db.Model(&PROVINCE{}).Create(&province30)

	province31 := PROVINCE{
		Province_Type: "ปราจีนบุรี",
	}
	db.Model(&PROVINCE{}).Create(&province31)

	province32 := PROVINCE{
		Province_Type: "ปัตตานี",
	}
	db.Model(&PROVINCE{}).Create(&province32)

	province33 := PROVINCE{
		Province_Type: "พะเยา",
	}
	db.Model(&PROVINCE{}).Create(&province33)

	province34 := PROVINCE{
		Province_Type: "พระนครศรีอยุธยา",
	}
	db.Model(&PROVINCE{}).Create(&province34)

	province35 := PROVINCE{
		Province_Type: "พังงา",
	}
	db.Model(&PROVINCE{}).Create(&province35)

	province36 := PROVINCE{
		Province_Type: "พัทลุง",
	}
	db.Model(&PROVINCE{}).Create(&province36)

	province37:= PROVINCE{
		Province_Type: "พิจิตร",
	}
	db.Model(&PROVINCE{}).Create(&province37)

	province38 := PROVINCE{
		Province_Type: "พิษณุโลก",
	}
	db.Model(&PROVINCE{}).Create(&province38)

	province39 := PROVINCE{
		Province_Type: "เพชรบุรี",
	}
	db.Model(&PROVINCE{}).Create(&province39)

	province40 := PROVINCE{
		Province_Type: "เพชรบูรณ์",
	}
	db.Model(&PROVINCE{}).Create(&province40)

	province41 := PROVINCE{
		Province_Type: "แพร่",
	}
	db.Model(&PROVINCE{}).Create(&province41)

	province42 := PROVINCE{
		Province_Type: "ภูเก็ต",
	}
	db.Model(&PROVINCE{}).Create(&province42)

	province43 := PROVINCE{
		Province_Type: "มหาสารคาม",
	}
	db.Model(&PROVINCE{}).Create(&province43)

	province44 := PROVINCE{
		Province_Type: "มุกดาหาร",
	}
	db.Model(&PROVINCE{}).Create(&province44)

	province45 := PROVINCE{
		Province_Type: "แม่ฮ่องสอน",
	}
	db.Model(&PROVINCE{}).Create(&province45)

	province46 := PROVINCE{
		Province_Type: "ยโสธร",
	}
	db.Model(&PROVINCE{}).Create(&province46)

	province47:= PROVINCE{
		Province_Type: "ยะลา",
	}
	db.Model(&PROVINCE{}).Create(&province47)

	province48 := PROVINCE{
		Province_Type: "ร้อยเอ็ด",
	}
	db.Model(&PROVINCE{}).Create(&province48)

	province49 := PROVINCE{
		Province_Type: "ระนอง",
	}
	db.Model(&PROVINCE{}).Create(&province49)

	province50 := PROVINCE{
		Province_Type: "ระยอง",
	}
	db.Model(&PROVINCE{}).Create(&province50)

	province51 := PROVINCE{
		Province_Type: "ราชบุรี",
	}
	db.Model(&PROVINCE{}).Create(&province51)

	province52 := PROVINCE{
		Province_Type: "ลพบุรี",
	}
	db.Model(&PROVINCE{}).Create(&province52)

	province53 := PROVINCE{
		Province_Type: "ลำปาง",
	}
	db.Model(&PROVINCE{}).Create(&province53)

	province54 := PROVINCE{
		Province_Type: "ลำพูน",
	}
	db.Model(&PROVINCE{}).Create(&province54)

	province55 := PROVINCE{
		Province_Type: "เลย",
	}
	db.Model(&PROVINCE{}).Create(&province55)

	province56 := PROVINCE{
		Province_Type: "ศรีสะเกษ",
	}
	db.Model(&PROVINCE{}).Create(&province56)

	province57:= PROVINCE{
		Province_Type: "สกลนคร",
	}
	db.Model(&PROVINCE{}).Create(&province57)

	province58 := PROVINCE{
		Province_Type: "สงขลา",
	}
	db.Model(&PROVINCE{}).Create(&province58)

	province59 := PROVINCE{
		Province_Type: "สตูล",
	}
	db.Model(&PROVINCE{}).Create(&province59)

	province60 := PROVINCE{
		Province_Type: "สมุทรปราการ",
	}
	db.Model(&PROVINCE{}).Create(&province60)

	province61 := PROVINCE{
		Province_Type: "สมุทรสงคราม",
	}
	db.Model(&PROVINCE{}).Create(&province61)

	province62 := PROVINCE{
		Province_Type: "สมุทรสาคร",
	}
	db.Model(&PROVINCE{}).Create(&province62)

	province63 := PROVINCE{
		Province_Type: "สระแก้ว",
	}
	db.Model(&PROVINCE{}).Create(&province63)

	province64 := PROVINCE{
		Province_Type: "สระบุรี",
	}
	db.Model(&PROVINCE{}).Create(&province64)

	province65 := PROVINCE{
		Province_Type: "สิงห์บุรี",
	}
	db.Model(&PROVINCE{}).Create(&province65)

	province66 := PROVINCE{
		Province_Type: "สุโขทัย",
	}
	db.Model(&PROVINCE{}).Create(&province66)

	province67:= PROVINCE{
		Province_Type: "สุพรรณบุรี",
	}
	db.Model(&PROVINCE{}).Create(&province67)

	province68 := PROVINCE{
		Province_Type: "สุราษฎร์ธานี",
	}
	db.Model(&PROVINCE{}).Create(&province68)

	province69 := PROVINCE{
		Province_Type: "สุรินทร์",
	}
	db.Model(&PROVINCE{}).Create(&province69)

	province70 := PROVINCE{
		Province_Type: "หนองคาย",
	}
	db.Model(&PROVINCE{}).Create(&province70)

	province71 := PROVINCE{
		Province_Type: "หนองบัวลำภู",
	}
	db.Model(&PROVINCE{}).Create(&province71)

	province72 := PROVINCE{
		Province_Type: "อ่างทอง",
	}
	db.Model(&PROVINCE{}).Create(&province72)

	province73 := PROVINCE{
		Province_Type: "อำนาจเจริญ",
	}
	db.Model(&PROVINCE{}).Create(&province73)

	province74 := PROVINCE{
		Province_Type: "อุดรธานี",
	}
	db.Model(&PROVINCE{}).Create(&province74)

	province75 := PROVINCE{
		Province_Type: "อุตรดิตถ์",
	}
	db.Model(&PROVINCE{}).Create(&province75)

	province76 := PROVINCE{
		Province_Type: "อุทัยธานี",
	}
	db.Model(&PROVINCE{}).Create(&province76)

	province77:= PROVINCE{
		Province_Type: "อุบลราชธานี",
	}
	db.Model(&PROVINCE{}).Create(&province77)

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
		Gender:      genM,
		Province:    province69,
		Package:     package2,
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
		Gender:      genF,
		Province:    province14,
		Package:     package1,
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
		Gender:      genM,
		Province:    province2,
		Package:     package3,
	})
}
