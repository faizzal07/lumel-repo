package main

import (
	"log"
	"lumel/internal/handlers"
	repository "lumel/internal/repositories"
	"lumel/internal/routers"
	"lumel/internal/services"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load("/Users/khalith/Documents/lumel/.env")
	logger := initLogger()
	if err != nil {
		logger.Error("Error loading env file")
	}
	connectString := os.Getenv("dbConnectString")
	db := repository.NewConnection(connectString)
	repo := repository.NewRepository(db, logger)
	service := services.NewService(logger, repo)
	handler := handlers.NewHandler(logger, service)
	if err != nil {
		return
	}

	r := routers.InitRouter(handler)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logLevel, _ := strconv.Atoi(os.Getenv("LogLevel"))
	logger.SetLevel(logrus.Level(logLevel))
	return logger
}
