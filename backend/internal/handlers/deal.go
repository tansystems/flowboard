package handlers

import (
	"net/http"

	"crm-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetDeals godoc
// @Summary      Получить список сделок
// @Description  Возвращает все сделки
// @Tags         deals
// @Produce      json
// @Success      200  {array}   models.Deal
// @Failure      500  {object}  map[string]string
// @Router       /deals [get]
func GetDeals(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var deals []models.Deal
		if err := db.Preload("Customer").Preload("Status").Find(&deals).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, deals)
	}
}

// GetDeal godoc
// @Summary      Получить сделку по ID
// @Description  Возвращает сделку по идентификатору
// @Tags         deals
// @Produce      json
// @Param        id   path      int  true  "ID сделки"
// @Success      200  {object}  models.Deal
// @Failure      404  {object}  map[string]string
// @Router       /deals/{id} [get]
func GetDeal(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var deal models.Deal
		id := c.Param("id")
		if err := db.Preload("Customer").Preload("Status").First(&deal, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Сделка не найдена"})
			return
		}
		c.JSON(http.StatusOK, deal)
	}
}

// CreateDeal godoc
// @Summary      Создать сделку
// @Description  Создаёт новую сделку
// @Tags         deals
// @Accept       json
// @Produce      json
// @Param        deal  body      models.Deal  true  "Данные сделки"
// @Success      201   {object}  models.Deal
// @Failure      400   {object}  map[string]string
// @Router       /deals [post]
func CreateDeal(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var deal models.Deal
		if err := c.ShouldBindJSON(&deal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&deal).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, deal)
	}
}

// UpdateDeal godoc
// @Summary      Обновить сделку
// @Description  Обновляет данные сделки по ID
// @Tags         deals
// @Accept       json
// @Produce      json
// @Param        id    path      int        true  "ID сделки"
// @Param        deal  body      models.Deal  true  "Данные сделки"
// @Success      200   {object}  models.Deal
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /deals/{id} [put]
func UpdateDeal(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var deal models.Deal
		id := c.Param("id")
		if err := db.First(&deal, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Сделка не найдена"})
			return
		}
		if err := c.ShouldBindJSON(&deal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&deal).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, deal)
	}
}

// DeleteDeal godoc
// @Summary      Удалить сделку
// @Description  Удаляет сделку по ID
// @Tags         deals
// @Produce      json
// @Param        id   path      int  true  "ID сделки"
// @Success      204  {object}  nil
// @Failure      500  {object}  map[string]string
// @Router       /deals/{id} [delete]
func DeleteDeal(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Deal{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
