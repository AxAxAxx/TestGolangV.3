package models

import "github.com/gofiber/fiber/v2"

type Books struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

func BookSuccessResponse(data Books) *fiber.Map {
	book := Books{
		ID:     data.ID,
		Title:  data.Title,
		Author: data.Author,
		Price:  data.Price,
	}
	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}

func BooksSuccessResponse(data *[]Books) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func BookDelete() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "delete successfully",
		"err":    nil}

}
