package controller

import (
	"net/http"

	"github.com/Earnny/sa-64/entity"
	"github.com/gin-gonic/gin"
)

// POST /tas
func CreateTA(c *gin.Context) { //Ta
	var ta entity.TA
	if err := c.ShouldBindJSON(&ta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&ta).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ta})
}

// GET /ta/:id
func GetTA(c *gin.Context) {
	var ta entity.TA

	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM tas WHERE id = ?", id).Find(&ta).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ta})
}

// GET /tas
func ListTAs(c *gin.Context) {
	var tas []entity.TA
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM tas").Find(&tas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tas})
}

// DELETE /tas/:id
func DeleteTA(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM tas WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TA not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /tas
func UpdateTA(c *gin.Context) {
	var ta entity.TA
	if err := c.ShouldBindJSON(&ta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", ta.ID).First(&ta); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TA not found"})
		return
	}

	if err := entity.DB().Save(&ta).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ta})
}
