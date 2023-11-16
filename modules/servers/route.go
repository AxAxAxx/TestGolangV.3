package handler

import (
	controller "github.com/AxAxAxx/gofiber/modules/users/controllers"
	repository "github.com/AxAxAxx/gofiber/modules/users/repositories"
	usecase "github.com/AxAxAxx/gofiber/modules/users/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Route(app *fiber.App, db *sqlx.DB) {
	bookRepository := repository.NewBoolRepository(db)
	bookUsecase := usecase.NewUserUsecase(*bookRepository)
	bookHandler := controller.NewBookHandler(*bookUsecase)

	app.Get("/books", bookHandler.GetBook)
	app.Post("/addbook", bookHandler.CreateBook)
	app.Put("/editbook/:id", bookHandler.UpdateUser)
	app.Delete("/deletebook/:id", bookHandler.DeleteUser)
}
