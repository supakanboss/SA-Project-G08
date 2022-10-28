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
		&GENDER{},
		&PACKAGE{},
		&MEMBER{},
		&PROVINCE{},
		&STAFF{},
		&SPORT_TYPE{},
		&SPORT_EQUIPMENT{},
		&SPORT_EQUIPMENT_TYPE{},
		&BORROW_SPORT_EQUIPMENT{},
	)

	db = database
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
		Province_Type: "กระบี่",
	}
	db.Model(&PROVINCE{}).Create(&Province1)

	Province2 := PROVINCE{
		Province_Type: "กรุงเทพมหานคร",
	}
	db.Model(&PROVINCE{}).Create(&Province2)

	Province3 := PROVINCE{
		Province_Type: "กาญจนบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province3)

	Province4 := PROVINCE{
		Province_Type: "กาฬสินธุ์",
	}
	db.Model(&PROVINCE{}).Create(&Province4)

	Province5 := PROVINCE{
		Province_Type: "กำแพงเพชร",
	}
	db.Model(&PROVINCE{}).Create(&Province5)

	Province6 := PROVINCE{
		Province_Type: "ขอนแก่น",
	}
	db.Model(&PROVINCE{}).Create(&Province6)

	Province7 := PROVINCE{
		Province_Type: "จันทบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province7)

	Province8 := PROVINCE{
		Province_Type: "ฉะเชิงเทรา",
	}
	db.Model(&PROVINCE{}).Create(&Province8)

	Province9 := PROVINCE{
		Province_Type: "ชลบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province9)

	Province10 := PROVINCE{
		Province_Type: "ชัยนาท",
	}
	db.Model(&PROVINCE{}).Create(&Province10)

	Province11 := PROVINCE{
		Province_Type: "ชัยภูมิ",
	}
	db.Model(&PROVINCE{}).Create(&Province11)

	Province12 := PROVINCE{
		Province_Type: "ชุมพร",
	}
	db.Model(&PROVINCE{}).Create(&Province12)

	Province13 := PROVINCE{
		Province_Type: "เชียงราย",
	}
	db.Model(&PROVINCE{}).Create(&Province13)

	Province14 := PROVINCE{
		Province_Type: "เชียงใหม่",
	}
	db.Model(&PROVINCE{}).Create(&Province14)

	Province15 := PROVINCE{
		Province_Type: "ตรัง",
	}
	db.Model(&PROVINCE{}).Create(&Province15)

	Province16 := PROVINCE{
		Province_Type: "ตราด",
	}
	db.Model(&PROVINCE{}).Create(&Province16)

	Province17 := PROVINCE{
		Province_Type: "ตาก",
	}
	db.Model(&PROVINCE{}).Create(&Province17)

	Province18 := PROVINCE{
		Province_Type: "นครนายก",
	}
	db.Model(&PROVINCE{}).Create(&Province18)

	Province19 := PROVINCE{
		Province_Type: "นครปฐม",
	}
	db.Model(&PROVINCE{}).Create(&Province19)

	Province20 := PROVINCE{
		Province_Type: "นครพนม",
	}
	db.Model(&PROVINCE{}).Create(&Province20)

	Province21 := PROVINCE{
		Province_Type: "นครราชสีมา",
	}
	db.Model(&PROVINCE{}).Create(&Province21)

	Province22 := PROVINCE{
		Province_Type: "นครศรีธรรมราช",
	}
	db.Model(&PROVINCE{}).Create(&Province22)

	Province23 := PROVINCE{
		Province_Type: "นครสวรรค์",
	}
	db.Model(&PROVINCE{}).Create(&Province23)

	Province24 := PROVINCE{
		Province_Type: "นนทบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province24)

	Province25 := PROVINCE{
		Province_Type: "นราธิวาส",
	}
	db.Model(&PROVINCE{}).Create(&Province25)

	Province26 := PROVINCE{
		Province_Type: "น่าน",
	}
	db.Model(&PROVINCE{}).Create(&Province26)

	Province27 := PROVINCE{
		Province_Type: "บึงกาฬ",
	}
	db.Model(&PROVINCE{}).Create(&Province27)

	Province28 := PROVINCE{
		Province_Type: "บุรีรัมย์",
	}
	db.Model(&PROVINCE{}).Create(&Province28)

	Province29 := PROVINCE{
		Province_Type: "ปทุมธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province29)

	Province30 := PROVINCE{
		Province_Type: "ประจวบคีรีขันธ์",
	}
	db.Model(&PROVINCE{}).Create(&Province30)

	Province31 := PROVINCE{
		Province_Type: "ปราจีนบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province31)

	Province32 := PROVINCE{
		Province_Type: "ปัตตานี",
	}
	db.Model(&PROVINCE{}).Create(&Province32)

	Province33 := PROVINCE{
		Province_Type: "พะเยา",
	}
	db.Model(&PROVINCE{}).Create(&Province33)

	Province34 := PROVINCE{
		Province_Type: "พระนครศรีอยุธยา",
	}
	db.Model(&PROVINCE{}).Create(&Province34)

	Province35 := PROVINCE{
		Province_Type: "พังงา",
	}
	db.Model(&PROVINCE{}).Create(&Province35)

	Province36 := PROVINCE{
		Province_Type: "พัทลุง",
	}
	db.Model(&PROVINCE{}).Create(&Province36)

	Province37 := PROVINCE{
		Province_Type: "พิจิตร",
	}
	db.Model(&PROVINCE{}).Create(&Province37)

	Province38 := PROVINCE{
		Province_Type: "พิษณุโลก",
	}
	db.Model(&PROVINCE{}).Create(&Province38)

	Province39 := PROVINCE{
		Province_Type: "เพชรบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province39)

	Province40 := PROVINCE{
		Province_Type: "เพชรบูรณ์",
	}
	db.Model(&PROVINCE{}).Create(&Province40)

	Province41 := PROVINCE{
		Province_Type: "แพร่",
	}
	db.Model(&PROVINCE{}).Create(&Province41)

	Province42 := PROVINCE{
		Province_Type: "ภูเก็ต",
	}
	db.Model(&PROVINCE{}).Create(&Province42)

	Province43 := PROVINCE{
		Province_Type: "มหาสารคาม",
	}
	db.Model(&PROVINCE{}).Create(&Province43)

	Province44 := PROVINCE{
		Province_Type: "มุกดาหาร",
	}
	db.Model(&PROVINCE{}).Create(&Province44)

	Province45 := PROVINCE{
		Province_Type: "แม่ฮ่องสอน",
	}
	db.Model(&PROVINCE{}).Create(&Province45)

	Province46 := PROVINCE{
		Province_Type: "ยโสธร",
	}
	db.Model(&PROVINCE{}).Create(&Province46)

	Province47 := PROVINCE{
		Province_Type: "ยะลา",
	}
	db.Model(&PROVINCE{}).Create(&Province47)

	Province48 := PROVINCE{
		Province_Type: "ร้อยเอ็ด",
	}
	db.Model(&PROVINCE{}).Create(&Province48)

	Province49 := PROVINCE{
		Province_Type: "ระนอง",
	}
	db.Model(&PROVINCE{}).Create(&Province49)

	Province50 := PROVINCE{
		Province_Type: "ระยอง",
	}
	db.Model(&PROVINCE{}).Create(&Province50)

	Province51 := PROVINCE{
		Province_Type: "ราชบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province51)

	Province52 := PROVINCE{
		Province_Type: "ลพบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province52)

	Province53 := PROVINCE{
		Province_Type: "ลำปาง",
	}
	db.Model(&PROVINCE{}).Create(&Province53)

	Province54 := PROVINCE{
		Province_Type: "ลำพูน",
	}
	db.Model(&PROVINCE{}).Create(&Province54)

	Province55 := PROVINCE{
		Province_Type: "เลย",
	}
	db.Model(&PROVINCE{}).Create(&Province55)

	Province56 := PROVINCE{
		Province_Type: "ศรีสะเกษ",
	}
	db.Model(&PROVINCE{}).Create(&Province56)

	Province57 := PROVINCE{
		Province_Type: "สกลนคร",
	}
	db.Model(&PROVINCE{}).Create(&Province57)

	Province58 := PROVINCE{
		Province_Type: "สงขลา",
	}
	db.Model(&PROVINCE{}).Create(&Province58)

	Province59 := PROVINCE{
		Province_Type: "สตูล",
	}
	db.Model(&PROVINCE{}).Create(&Province59)

	Province60 := PROVINCE{
		Province_Type: "สมุทรปราการ",
	}
	db.Model(&PROVINCE{}).Create(&Province60)

	Province61 := PROVINCE{
		Province_Type: "สมุทรสงคราม",
	}
	db.Model(&PROVINCE{}).Create(&Province61)

	Province62 := PROVINCE{
		Province_Type: "สมุทรสาคร",
	}
	db.Model(&PROVINCE{}).Create(&Province62)

	Province63 := PROVINCE{
		Province_Type: "สระแก้ว",
	}
	db.Model(&PROVINCE{}).Create(&Province63)

	Province64 := PROVINCE{
		Province_Type: "สระบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province64)

	Province65 := PROVINCE{
		Province_Type: "สิงห์บุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province65)

	Province66 := PROVINCE{
		Province_Type: "สุโขทัย",
	}
	db.Model(&PROVINCE{}).Create(&Province66)

	Province67 := PROVINCE{
		Province_Type: "สุพรรณบุรี",
	}
	db.Model(&PROVINCE{}).Create(&Province67)

	Province68 := PROVINCE{
		Province_Type: "สุราษฎร์ธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province68)

	Province69 := PROVINCE{
		Province_Type: "สุรินทร์",
	}
	db.Model(&PROVINCE{}).Create(&Province69)

	Province70 := PROVINCE{
		Province_Type: "หนองคาย",
	}
	db.Model(&PROVINCE{}).Create(&Province70)

	Province71 := PROVINCE{
		Province_Type: "หนองบัวลำภู",
	}
	db.Model(&PROVINCE{}).Create(&Province71)

	Province72 := PROVINCE{
		Province_Type: "อ่างทอง",
	}
	db.Model(&PROVINCE{}).Create(&Province72)

	Province73 := PROVINCE{
		Province_Type: "อำนาจเจริญ",
	}
	db.Model(&PROVINCE{}).Create(&Province73)

	Province74 := PROVINCE{
		Province_Type: "อุดรธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province74)

	Province75 := PROVINCE{
		Province_Type: "อุตรดิตถ์",
	}
	db.Model(&PROVINCE{}).Create(&Province75)

	Province76 := PROVINCE{
		Province_Type: "อุทัยธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province76)

	Province77 := PROVINCE{
		Province_Type: "อุบลราชธานี",
	}
	db.Model(&PROVINCE{}).Create(&Province77)

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
	password4, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password5, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password6, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)

	Staff1 := STAFF{
		Staff_name:     "นฤมน เกษรบัว",
		Staff_email:    "Kaenniza55@gmail.com",
		Staff_password: string(password4),
		Staff_address:  "114 ม.10 ต.ในเมือง จ.สุรินทร์",
		Staff_tel:      "0954101589",
		Staff_BOD:      t,
		Gender:         GenderFemale,
	}
	db.Model(&STAFF{}).Create(&Staff1)

	Staff2 := STAFF{
		Staff_name:     "ใจดี สมใจ",
		Staff_email:    "Jaideehihi@gmail.com",
		Staff_password: string(password5),
		Staff_address:  "555 ม.1 ต.ในเมือง จ.ขอนแก่น",
		Staff_tel:      "0973340404",
		Staff_BOD:      t,
		Gender:         GenderMale,
	}
	db.Model(&STAFF{}).Create(&Staff2)

	Staff3 := STAFF{
		Staff_name:     "บุญมี กรรมบัง",
		Staff_email:    "Boomee@gmail.com",
		Staff_password: string(password6),
		Staff_address:  "561 ม.1 ต.ในเมือง จ.ขอนแก่น",
		Staff_tel:      "0983380848",
		Staff_BOD:      t,
		Gender:         GenderOther,
	}
	db.Model(&STAFF{}).Create(&Staff3)

	//ข้อมํูลในตาราง sport epuipment tpye
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


	//----SPORT_EQUIPMENT----

	EQUIPMENT1 := SPORT_EQUIPMENT{
		Sport_equipment_type:   Sport_Equipment_Type3,
		Sport_type:             Sport_type1,
		Staff:                  Staff1,
		Sport_Equipment_Name:   "ลูกฟุตบอล",
		Sport_Equipment_Amount: 10,
	}

	db.Model(&SPORT_EQUIPMENT{}).Create(&EQUIPMENT1)

	EQUIPMENT2 := SPORT_EQUIPMENT{
		Sport_equipment_type:   Sport_Equipment_Type3,
		Sport_type:             Sport_type3,
		Staff:                  Staff1,
		Sport_Equipment_Name:   "ลูกขนไก่",
		Sport_Equipment_Amount: 24,
	}
	db.Model(&SPORT_EQUIPMENT{}).Create(&EQUIPMENT2)

	EQUIPMENT3 := SPORT_EQUIPMENT{
		Sport_equipment_type:   Sport_Equipment_Type2,
		Sport_type:             Sport_type3,
		Staff:                  Staff1,
		Sport_Equipment_Name:   "ไม้แบดมินตัน",
		Sport_Equipment_Amount: 5,
	}
	db.Model(&SPORT_EQUIPMENT{}).Create(&EQUIPMENT3)


}
