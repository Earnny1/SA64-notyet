package controller

import (
	"net/http"

	"github.com/Earnny/sa-64/entity"
	"github.com/gin-gonic/gin"
)

// POST /ManageCourses
func CreateManageCourse(c *gin.Context) {

	var manageCourse entity.ManageCourse
	var professor entity.Professor
	var course entity.Course
	var ta entity.TA
	var room entity.Room

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร manageCourse
	if err := c.ShouldBindJSON(&manageCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา course ด้วย id
	if tx := entity.DB().Where("id = ?", manageCourse.CourseID).First(&course); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course not found"})
		return
	}

	// 10: ค้นหา ta ด้วย id
	if tx := entity.DB().Where("id = ?", manageCourse.TAID).First(&ta); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ta not found"})
		return
	}

	// 11: ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", manageCourse.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	// 12: สร้าง manageCourse
	wv := entity.ManageCourse{
		Professor:        professor,                     // โยงความสัมพันธ์กับ Entity Taecher
		Course:           course,                        // โยงความสัมพันธ์กับ Entity Course
		TA:               ta,                            // โยงความสัมพันธ์กับ Entity TA
		Room:             room,                          // โยงความสัมพันธ์กับ Entity Room
		ManageCourseTime: manageCourse.ManageCourseTime, // ตั้งค่าฟิลด์ ManageCourseTime
		Group:            manageCourse.Group,            // ตั้งค่าฟิลด์ Group
		Term:             manageCourse.Term,             // ตั้งค่าฟิลด์ Trimester
		TeachingTime:     manageCourse.TeachingTime,     // ตั้งค่าฟิลด์ Teaching_Time
		Ungraduated_year: manageCourse.Ungraduated_year, // ตั้งค่าฟิลด์ Ungraduared_year
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /manageCourses/:id
func GetManageCourse(c *gin.Context) {
	var manageCourse entity.ManageCourse
	id := c.Param("id")
	if err := entity.DB().Preload("Professor").Preload("Course").Preload("TA").Preload("Room").Raw("SELECT * FROM manageCourses WHERE id = ?", id).Find(&manageCourse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": manageCourse})
}

// GET /manageCourses
func ListManageCourses(c *gin.Context) {
	var manageCourses []entity.ManageCourse
	if err := entity.DB().Preload("Professor").Preload("Course").Preload("TA").Preload("Room").Raw("SELECT * FROM manageCourses").Find(&manageCourses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": manageCourses})
}

// DELETE /managCourses/:id
func DeleteManagCourse(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM managCourses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "managcourse not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /manageCourses
func UpdateManageCourse(c *gin.Context) {
	var manageCourse entity.ManageCourse
	if err := c.ShouldBindJSON(&manageCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", manageCourse.ID).First(&manageCourse); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "manageCourse not found"})
		return
	}

	if err := entity.DB().Save(&manageCourse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": manageCourse})
}
