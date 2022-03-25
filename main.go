package main

import (
	"golang-structure/src/common/logger"

	configuration "golang-structure/src/configs"

	"golang-structure/src/database"
	"golang-structure/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	configuration.LoadConfig()

	logger.InitLogger()

	logger.ProcessLog.Info().Msg("hello anh em")

	database.ConnectMongo()

	app.Use(recover.New())

	routes.SetupRoutes(app)

	app.Listen(configuration.Config.GetString("uri"))
}
