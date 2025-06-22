package main

import (
	"fmt"
	"log"
	"os"

	"crm-backend/internal/handlers"
	"crm-backend/internal/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title CRM API
// @version 1.0
// @description Мини-CRM/ERP для малого бизнеса
// @host localhost:8080
// @BasePath /

func main() {
	// Получаем параметры подключения из переменных окружения
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=crm port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	fmt.Println("Успешное подключение к базе данных")

	// Миграция моделей
	db.AutoMigrate(
		&models.Customer{},
		&models.Deal{},
		&models.Status{},
		&models.Comment{},
		&models.Tag{},
		&models.User{},
	)

	r := gin.Default()

	// Auth
	r.POST("/auth/register", handlers.Register(db))
	r.POST("/auth/login", handlers.Login(db))
	r.GET("/auth/me", handlers.JWTAuthMiddleware(), handlers.Me(db))

	// CRUD для клиентов
	r.GET("/customers", handlers.GetCustomers(db))
	r.GET("/customers/:id", handlers.GetCustomer(db))
	r.POST("/customers", handlers.CreateCustomer(db))
	r.PUT("/customers/:id", handlers.UpdateCustomer(db))
	r.DELETE("/customers/:id", handlers.DeleteCustomer(db))

	// CRUD для сделок
	r.GET("/deals", handlers.GetDeals(db))
	r.GET("/deals/:id", handlers.GetDeal(db))
	r.POST("/deals", handlers.CreateDeal(db))
	r.PUT("/deals/:id", handlers.UpdateDeal(db))
	r.DELETE("/deals/:id", handlers.DeleteDeal(db))

	// CRUD для статусов
	r.GET("/statuses", handlers.GetStatuses(db))
	r.GET("/statuses/:id", handlers.GetStatus(db))
	r.POST("/statuses", handlers.CreateStatus(db))
	r.PUT("/statuses/:id", handlers.UpdateStatus(db))
	r.DELETE("/statuses/:id", handlers.DeleteStatus(db))

	// CRUD для тегов
	r.GET("/tags", handlers.GetTags(db))
	r.GET("/tags/:id", handlers.GetTag(db))
	r.POST("/tags", handlers.CreateTag(db))
	r.PUT("/tags/:id", handlers.UpdateTag(db))
	r.DELETE("/tags/:id", handlers.DeleteTag(db))

	// CRUD для пользователей (требует авторизации)
	u := r.Group("/users")
	u.Use(handlers.JWTAuthMiddleware())
	u.GET("", handlers.GetUsers(db))
	u.GET(":id", handlers.GetUser(db))
	u.POST("", handlers.CreateUser(db))
	u.PUT(":id", handlers.UpdateUser(db))
	u.DELETE(":id", handlers.DeleteUser(db))

	// CRUD для комментариев (требует авторизации)
	cmt := r.Group("/comments")
	cmt.Use(handlers.JWTAuthMiddleware())
	cmt.GET("", handlers.GetComments(db))
	cmt.GET(":id", handlers.GetComment(db))
	cmt.POST("", handlers.CreateComment(db))
	cmt.PUT(":id", handlers.UpdateComment(db))
	cmt.DELETE(":id", handlers.DeleteComment(db))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080") // Запуск сервера на порту 8080
}
