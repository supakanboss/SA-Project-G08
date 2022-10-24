package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pradthana7/sa-65-sq/entity"
	"github.com/pradthana7/sa-65-sq/service"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload login body
type LoginPayload struct {
	Staff_email    string `json:"Staff_email"`
	Staff_password string `json:"Staff_password"`
}

// SignUpPayload signup body
type SignUpPayload struct {
	Staff_name     string `json:"Staff_name"`
	Staff_email    string `json:"Staff_email"`
	Staff_password string `json:"Staff_password"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

// POST /login
func StaffLogin(c *gin.Context) {
	var payload LoginPayload
	var staff entity.STAFF

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา staff ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM staffs WHERE staff_email = ?", payload.Staff_email).Scan(&staff).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(staff.Staff_password), []byte(payload.Staff_password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incerrect"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(staff.Staff_email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
		ID:    staff.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

// POST /create
func CreateUser(c *gin.Context) {
	var payload SignUpPayload
	var staff entity.STAFF

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Staff_password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}

	staff.Staff_name = payload.Staff_name
	staff.Staff_email = payload.Staff_email
	staff.Staff_password = string(hashPassword)

	if err := entity.DB().Create(&staff).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": staff})
}
