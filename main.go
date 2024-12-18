package main

import (
	"api/first-go/auth"
	"api/first-go/common"
	"api/first-go/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
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

	auth.Init(router, db)

	router.Listen(configs.Cfg.App.Port)
}