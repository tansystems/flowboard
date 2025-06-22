package handlers

import (
	"crm-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetTags godoc
// @Summary      Получить список тегов
// @Description  Возвращает все теги
// @Tags         tags
// @Produce      json
// @Success      200  {array}   models.Tag
// @Failure      500  {object}  map[string]string
// @Router       /tags [get]
func GetTags(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tags []models.Tag
		if err := db.Preload("Deals").Find(&tags).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tags)
	}
}

// GetTag godoc
// @Summary      Получить тег по ID
// @Description  Возвращает тег по идентификатору
// @Tags         tags
// @Produce      json
// @Param        id   path      int  true  "ID тега"
// @Success      200  {object}  models.Tag
// @Failure      404  {object}  map[string]string
// @Router       /tags/{id} [get]
func GetTag(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tag models.Tag
		id := c.Param("id")
		if err := db.Preload("Deals").First(&tag, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Тег не найден"})
			return
		}
		c.JSON(http.StatusOK, tag)
	}
}

// CreateTag godoc
// @Summary      Создать тег
// @Description  Создаёт новый тег
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        tag  body      models.Tag  true  "Данные тега"
// @Success      201  {object}  models.Tag
// @Failure      400  {object}  map[string]string
// @Router       /tags [post]
func CreateTag(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tag models.Tag
		if err := c.ShouldBindJSON(&tag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&tag).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, tag)
	}
}

// UpdateTag godoc
// @Summary      Обновить тег
// @Description  Обновляет данные тега по ID
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        id   path      int       true  "ID тега"
// @Param        tag  body      models.Tag  true  "Данные тега"
// @Success      200  {object}  models.Tag
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /tags/{id} [put]
func UpdateTag(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tag models.Tag
		id := c.Param("id")
		if err := db.First(&tag, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Тег не найден"})
			return
		}
		if err := c.ShouldBindJSON(&tag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&tag).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tag)
	}
}

// DeleteTag godoc
// @Summary      Удалить тег
// @Description  Удаляет тег по ID
// @Tags         tags
// @Produce      json
// @Param        id   path      int  true  "ID тега"
// @Success      204  {object}  nil
// @Failure      500  {object}  map[string]string
// @Router       /tags/{id} [delete]
func DeleteTag(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Tag{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
