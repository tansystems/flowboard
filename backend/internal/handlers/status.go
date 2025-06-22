package handlers

import (
	"crm-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetStatuses godoc
// @Summary      Получить список статусов
// @Description  Возвращает все статусы
// @Tags         statuses
// @Produce      json
// @Success      200  {array}   models.Status
// @Failure      500  {object}  map[string]string
// @Router       /statuses [get]
func GetStatuses(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var statuses []models.Status
		if err := db.Find(&statuses).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, statuses)
	}
}

// GetStatus godoc
// @Summary      Получить статус по ID
// @Description  Возвращает статус по идентификатору
// @Tags         statuses
// @Produce      json
// @Param        id   path      int  true  "ID статуса"
// @Success      200  {object}  models.Status
// @Failure      404  {object}  map[string]string
// @Router       /statuses/{id} [get]
func GetStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var status models.Status
		id := c.Param("id")
		if err := db.First(&status, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Статус не найден"})
			return
		}
		c.JSON(http.StatusOK, status)
	}
}

// CreateStatus godoc
// @Summary      Создать статус
// @Description  Создаёт новый статус
// @Tags         statuses
// @Accept       json
// @Produce      json
// @Param        status  body      models.Status  true  "Данные статуса"
// @Success      201     {object}  models.Status
// @Failure      400     {object}  map[string]string
// @Router       /statuses [post]
func CreateStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var status models.Status
		if err := c.ShouldBindJSON(&status); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, status)
	}
}

// UpdateStatus godoc
// @Summary      Обновить статус
// @Description  Обновляет данные статуса по ID
// @Tags         statuses
// @Accept       json
// @Produce      json
// @Param        id      path      int           true  "ID статуса"
// @Param        status  body      models.Status true  "Данные статуса"
// @Success      200     {object}  models.Status
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Router       /statuses/{id} [put]
func UpdateStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var status models.Status
		id := c.Param("id")
		if err := db.First(&status, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Статус не найден"})
			return
		}
		if err := c.ShouldBindJSON(&status); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, status)
	}
}

// DeleteStatus godoc
// @Summary      Удалить статус
// @Description  Удаляет статус по ID
// @Tags         statuses
// @Produce      json
// @Param        id   path      int  true  "ID статуса"
// @Success      204  {object}  nil
// @Failure      500  {object}  map[string]string
// @Router       /statuses/{id} [delete]
func DeleteStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsAdmin(c) {
			c.JSON(403, gin.H{"error": "Только администратор может удалять статусы"})
			return
		}
		id := c.Param("id")
		if err := db.Delete(&models.Status{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
