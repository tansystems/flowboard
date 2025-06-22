package handlers

import (
	"crm-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetComments godoc
// @Summary      Получить список комментариев
// @Description  Возвращает все комментарии
// @Tags         comments
// @Produce      json
// @Success      200  {array}   models.Comment
// @Failure      500  {object}  map[string]string
// @Router       /comments [get]
func GetComments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comments []models.Comment
		if err := db.Preload("User").Preload("Deal").Find(&comments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, comments)
	}
}

// GetComment godoc
// @Summary      Получить комментарий по ID
// @Description  Возвращает комментарий по идентификатору
// @Tags         comments
// @Produce      json
// @Param        id   path      int  true  "ID комментария"
// @Success      200  {object}  models.Comment
// @Failure      404  {object}  map[string]string
// @Router       /comments/{id} [get]
func GetComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment models.Comment
		id := c.Param("id")
		if err := db.Preload("User").Preload("Deal").First(&comment, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Комментарий не найден"})
			return
		}
		c.JSON(http.StatusOK, comment)
	}
}

func CreateComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment models.Comment
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, comment)
	}
}

func UpdateComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment models.Comment
		id := c.Param("id")
		if err := db.First(&comment, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Комментарий не найден"})
			return
		}
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, comment)
	}
}

func DeleteComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Comment{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
