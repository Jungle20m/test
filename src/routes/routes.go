package routes

import (
	"golang-structure/src/controllers"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	_ "golang-structure/docs"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/audit")

	v1 := api.Group("/v1")

	// Monitoring
	v1.Get("/monitor", monitor.New())

	// Swagger
	v1.Use("/docs", swagger.HandlerDefault)

	// Mobile Card
	mobile_cards := v1.Group("/mobile-cards")
	mobile_cards.Get("/details", controllers.GetAllMobileCardExchange)
	mobile_cards.Get("/gmv", controllers.GetMobileCardGMV)

	// Topup
	topup := v1.Group("/topup")
	topup.Get("/details", controllers.GetTopupDetails)
	topup.Get("/gmv", controllers.GetTopupGMV)
}
