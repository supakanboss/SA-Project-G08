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
		&GENDER{},   //เพิ่มค่าเข้าไป
		&PACKAGE{},  //เพิ่มค่าเข้าไป
		&PROVINCE{}, //เพิ่มค่าเข้าไป
		&MEMBER{},
		&LOCATION{},   //เพิ่มค่าเข้าไป
		&SPORT_TYPE{}, //เพิ่มค่าเข้าไป
		&LOCATION_RESERVATION{},
		&STAFF{},                //เพิ่มค่าเข้าไป
		&SPORT_EQUIPMENT_TYPE{}, //เพิ่มค่าเข้าไป
		&REPORT{},
		&SPORT_EQUIPMENT{},
		&BORROW_SPORT_EQUIPMENT{},
		&PAYMENT_TYPE{}, //เพิ่มค่าเข้าไป
		&BANK{},         //เพิ่มค่าเข้าไป
		&PAYMENT{},
		&STATUS{}, //เพิ่มค่าเข้าไป
		&CHECK_IN_OUT{},
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

	password1, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)

	Staff1 := STAFF{
		Staff_name:     "นฤมน เกษรบัว",
		Staff_email:    "Kaenniza55@gmail.com",
		Staff_password: string(password1),
		Staff_address:  "114 ม.10 ต.ในเมือง จ.สุรินทร์",
		Staff_tel:      "0954101589",
		Staff_BOD:      t,
		Gender:         GenderFemale,
	}
	db.Model(&STAFF{}).Create(&Staff1)

	Staff2 := STAFF{
		Staff_name:     "ใจดี สมใจ",
		Staff_email:    "Jaideehihi@gmail.com",
		Staff_password: string(password2),
		Staff_address:  "555 ม.1 ต.ในเมือง จ.ขอนแก่น",
		Staff_tel:      "0973340404",
		Staff_BOD:      t,
		Gender:         GenderMale,
	}
	db.Model(&STAFF{}).Create(&Staff2)

	Staff3 := STAFF{
		Staff_name:     "บุญมี กรรมบัง",
		Staff_email:    "Boomee@gmail.com",
		Staff_password: string(password3),
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

	///////////////ข้อมูลใน entity STATUS///////////
	Status1 := STATUS{
		Discribe: "Check-in",
	}
	db.Model(&STATUS{}).Create(&Status1)
	Status2 := STATUS{
		Discribe: "Check-out",
	}
	db.Model(&STATUS{}).Create(&Status2)

}
