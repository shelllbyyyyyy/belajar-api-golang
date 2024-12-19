package main

import (
	"api/first-go/apps/auth/presentation/controller"
	"api/first-go/common"
	"api/first-go/configs"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	pwd, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    
    err = godotenv.Load(filepath.Join(pwd, "configs", ".env"))
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	filename := "./config.yaml"
	if err := configs.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := configs.ConnectPostgres(configs.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: configs.Cfg.App.Name,
	})

	router.Use(common.LoggerMiddleware())

	controller.AuthRoute(router, db)
	controller.UserRoute(router, db)

	router.Listen(configs.Cfg.App.Port)
}