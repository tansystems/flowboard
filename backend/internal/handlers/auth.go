package handlers

import (
	"net/http"
	"os"
	"time"

	"crm-backend/internal/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET не задан в переменных окружения")
	}
	jwtKey = []byte(secret)
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// Register godoc
// @Summary      Регистрация пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "Данные пользователя"
// @Success      201   {object}  models.User
// @Failure      400   {object}  map[string]string
// @Router       /auth/register [post]
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if input.PasswordHash == "" || input.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email и пароль обязательны"})
			return
		}
		// Проверка на уникальность email
		var exists models.User
		if err := db.Where("email = ?", input.Email).First(&exists).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email уже зарегистрирован"})
			return
		}
		// Хешируем пароль
		hash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеширования пароля"})
			return
		}
		input.PasswordHash = string(hash)
		if input.Role == "" {
			input.Role = "user"
		}
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		input.PasswordHash = "" // не возвращаем хеш
		c.JSON(http.StatusCreated, input)
	}
}

// Login godoc
// @Summary      Логин пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body  map[string]string  true  "Email и пароль"
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /auth/login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var user models.User
		if err := db.Where("email = ?", creds.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
			return
		}
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			UserID: user.ID,
			Email:  user.Email,
			Role:   user.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

// JWT middleware
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Требуется токен авторизации"})
			return
		}
		// Bearer <token>
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}

// Проверка, является ли пользователь админом
func IsAdmin(c *gin.Context) bool {
	role, ok := c.Get("user_role")
	if !ok {
		return false
	}
	return role == "admin"
}

// Me godoc
// @Summary      Получить информацию о текущем пользователе
// @Tags         auth
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      401  {object}  map[string]string
// @Router       /auth/me [get]
func Me(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("user_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось определить пользователя"})
			return
		}
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
			return
		}
		user.PasswordHash = ""
		c.JSON(http.StatusOK, user)
	}
}
