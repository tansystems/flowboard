package handlers

import (
	"net/http"

	"crm-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetCustomers godoc
// @Summary      Получить список клиентов
// @Description  Возвращает всех клиентов
// @Tags         customers
// @Produce      json
// @Success      200  {array}   models.Customer
// @Failure      500  {object}  map[string]string
// @Router       /customers [get]
func GetCustomers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customers []models.Customer
		if err := db.Find(&customers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, customers)
	}
}

// GetCustomer godoc
// @Summary      Получить клиента по ID
// @Description  Возвращает клиента по идентификатору
// @Tags         customers
// @Produce      json
// @Param        id   path      int  true  "ID клиента"
// @Success      200  {object}  models.Customer
// @Failure      404  {object}  map[string]string
// @Router       /customers/{id} [get]
func GetCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customer models.Customer
		id := c.Param("id")
		if err := db.First(&customer, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Клиент не найден"})
			return
		}
		c.JSON(http.StatusOK, customer)
	}
}

// CreateCustomer godoc
// @Summary      Создать клиента
// @Description  Создаёт нового клиента
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        customer  body      models.Customer  true  "Данные клиента"
// @Success      201       {object}  models.Customer
// @Failure      400       {object}  map[string]string
// @Router       /customers [post]
func CreateCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customer models.Customer
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&customer).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, customer)
	}
}

// UpdateCustomer godoc
// @Summary      Обновить клиента
// @Description  Обновляет данные клиента по ID
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        id        path      int              true  "ID клиента"
// @Param        customer  body      models.Customer  true  "Данные клиента"
// @Success      200       {object}  models.Customer
// @Failure      400       {object}  map[string]string
// @Failure      404       {object}  map[string]string
// @Router       /customers/{id} [put]
func UpdateCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customer models.Customer
		id := c.Param("id")
		if err := db.First(&customer, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Клиент не найден"})
			return
		}
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&customer).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, customer)
	}
}

// DeleteCustomer godoc
// @Summary      Удалить клиента
// @Description  Удаляет клиента по ID
// @Tags         customers
// @Produce      json
// @Param        id   path      int  true  "ID клиента"
// @Success      204  {object}  nil
// @Failure      500  {object}  map[string]string
// @Router       /customers/{id} [delete]
func DeleteCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Customer{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
