package handlers

import (
	"crm-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUsers godoc
// @Summary      Получить список пользователей
// @Description  Возвращает всех пользователей
// @Tags         users
// @Produce      json
// @Success      200  {array}   models.User
// @Failure      500  {object}  map[string]string
// @Router       /users [get]
func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// GetUser godoc
// @Summary      Получить пользователя по ID
// @Description  Возвращает пользователя по идентификатору
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "ID пользователя"
// @Success      200  {object}  models.User
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser godoc
// @Summary      Создать пользователя
// @Description  Создаёт нового пользователя
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "Данные пользователя"
// @Success      201   {object}  models.User
// @Failure      400   {object}  map[string]string
// @Router       /users [post]
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	}
}

// UpdateUser godoc
// @Summary      Обновить пользователя
// @Description  Обновляет данные пользователя по ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "ID пользователя"
// @Param        user  body      models.User true  "Данные пользователя"
// @Success      200   {object}  models.User
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /users/{id} [put]
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser godoc
// @Summary      Удалить пользователя
// @Description  Удаляет пользователя по ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "ID пользователя"
// @Success      204  {object}  nil
// @Failure      500  {object}  map[string]string
// @Router       /users/{id} [delete]
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsAdmin(c) {
			c.JSON(403, gin.H{"error": "Только администратор может удалять пользователей"})
			return
		}
		id := c.Param("id")
		if err := db.Delete(&models.User{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
