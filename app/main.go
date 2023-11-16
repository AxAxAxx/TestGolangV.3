package main

import (
	"log"

	route "github.com/AxAxAxx/gofiber/modules/servers"
	db "github.com/AxAxAxx/gofiber/pkg/databases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	db, err := db.ConnPgSQL()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture postgres book shop!"))
	})
	route.Route(app, db.DB)
	app.Listen(":3000")
}
