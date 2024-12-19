package main

import (
	"api/first-go/apps/auth/presentation/controller"
	"api/first-go/common"
	"api/first-go/configs"
	_ "api/first-go/docs/first_api"
	"api/first-go/scripts"
	"log"
	"os"
	"path/filepath"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /
// @schemes http
func main() {
	pwd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

	migrator := scripts.MustGetNewMigrator()
    
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

	defer db.Close()

	err = migrator.ApplyMigrations(db.DB)
 	if err != nil {
		log.Println(err)
 	}

	if db != nil {
		log.Println("db connected")
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
		AppName: configs.Cfg.App.Name,
	})

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(common.LoggerMiddleware())

	app.Get("/", HealthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault)

	controller.AuthRoute(app, db)
	controller.UserRoute(app, db)

	app.Listen(configs.Cfg.App.Port)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}