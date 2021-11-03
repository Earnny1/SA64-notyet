package controller

import (
	"net/http"

	"github.com/Earnny/sa-64/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// POST /professors
func CreateProfessor(c *gin.Context) {
	var professor entity.Professor
	if err := c.ShouldBindJSON(&professor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(professor.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	professor.Password = string(bytes)
	if err := entity.DB().Create(&professor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": professor})
}

// GET /professor/:id
func GetProfessor(c *gin.Context) {
	var professor entity.Professor
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM professors WHERE id = ?", id).Find(&professor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": professor})
}

// GET /professors
func ListProfessors(c *gin.Context) {
	var professors []entity.Professor
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM professors").Find(&professors).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": professors})
}

// DELETE /professors/:id
func DeleteProfessor(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM professors WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "professor not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /professors
func UpdateProfessor(c *gin.Context) {
	var professor entity.Professor
	if err := c.ShouldBindJSON(&professor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", professor.ID).First(&professor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "professor not found"})
		return
	}

	if err := entity.DB().Save(&professor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": professor})
}
