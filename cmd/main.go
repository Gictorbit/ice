package main

import (
	"database/sql"
	"github.com/gictorbit/ice/internal/todo/delivery"
	"github.com/gictorbit/ice/internal/todo/infrastructure/mysql"
	redisinfra "github.com/gictorbit/ice/internal/todo/infrastructure/redis"
	"github.com/gictorbit/ice/internal/todo/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(mysql:3306)/todos?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	todoRepo := mysql.NewTodoMySQL(db)
	stream := redisinfra.NewStreamPublisher(rdb, "todos")
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	uc := usecase.NewTodoUseCase(todoRepo, stream, logger)
	h := delivery.NewTodoHandler(uc, logger)

	http.HandleFunc("/todo", h.CreateTodo)
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
