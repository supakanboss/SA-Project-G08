package controller

import (
	"github.com/bxngearnx/sa-demo/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST
func CreateBorrow_Sport_Eqiupment(c *gin.Context) {

	var Member entity.MEMBER
	var Sport_Equipment_Type entity.SPORT_EQUIPMENT_TYPE
	var Sport_Equipment entity.SPORT_EQUIPMENT
	var Borrow_Sport_Equipment entity.BORROW_SPORT_EQUIPMENT

	if err := c.ShouldBindJSON(&Borrow_Sport_Equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//////////////////////////////// 3: ค้นหา Member ด้วย Email ส่วนของlogin/////////////////////////////////////////////
	if tx := entity.DB().Where("id = ?", Borrow_Sport_Equipment.MemberID).First(&Member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Member not found"})
		return
	}
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//8:ค้นหา Sport Equipment ด้วย id
	if tx := entity.DB().Where("id = ?", Borrow_Sport_Equipment.Sport_EquipmentID).First(&Sport_Equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sport Eqipment not found"})
		return
	}

	// 9: ค้นหา Sport Equipment Type  ด้วย id
	if tx := entity.DB().Where("id = ?", Borrow_Sport_Equipment.Sport_Equipment_TypeID).First(&Sport_Equipment_Type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sport Equipment Type not found"})
		return
	}

	// 10: สร้าง entity Borrow_Sport_Eqiupment
	lr := entity.BORROW_SPORT_EQUIPMENT{
		MemberID:             Borrow_Sport_Equipment.MemberID,
		Member:               Member,                               // โยงความสัมพันธ์กับ Entity Member
		Sport_Equipment_Type: Sport_Equipment_Type,                 // โยงความสัมพันธ์กับ Entity Sport_Equipment_Type
		Sport_Equipment:      Sport_Equipment,                      // โยงความสัมพันธ์กับ Entity Sport_Equipment
		Time_Borrow:          Borrow_Sport_Equipment.Time_Borrow,   //รายละเอียด เวลายืม ที่กรอกเข้ามาใน Borrow_Sport_Eqiupment
		Time_Giveback:        Borrow_Sport_Equipment.Time_Giveback, //รายละเอียด เวลาคืน ที่กรอกเข้ามาใน Borrow_Sport_Eqiupment

	}
	// 11: บันทึก
	if err := entity.DB().Create(&lr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": lr})
}

// 5: GET sport equipment type เตรียมข้อมูลให้ combobox
func ListSport_Equipment_Type(c *gin.Context) {
	var Sport_Equipment_Type []entity.SPORT_EQUIPMENT_TYPE

	if err := entity.DB().Raw("SELECT * FROM sport_eq_ui_pment_types").Scan(&Sport_Equipment_Type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Sport_Equipment_Type})
}

// 6: GET sport equipment เตรียมข้อมูลให้ combobox
func ListSport_Equipment(c *gin.Context) {
	var Sport_Equipment []entity.SPORT_EQUIPMENT

	if err := entity.DB().Raw("SELECT * FROM sport_eq_ui_pments").Scan(&Sport_Equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Sport_Equipment})
}

// ดึงข้อมูลบันทึกอุปกรณ์มาแสดง /borrow-sport-eqiupments
func ListBorrow_Sport_Eqiupment(c *gin.Context) {

	var borrowsporteq []entity.BORROW_SPORT_EQUIPMENT

	if err := entity.DB().Preload("Member").Preload("Sport_Equipment_Type").Preload("Sport_Equipment").Raw("SELECT * FROM borrow_sport_eq_ui_pments").Scan(&borrowsporteq).Find(&borrowsporteq).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": borrowsporteq})

}

// GET /member/:id
func GetMember(c *gin.Context) {
	var member entity.MEMBER
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM members WHERE id = ?", id).Scan(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": member})
}
