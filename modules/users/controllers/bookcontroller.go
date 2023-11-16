package handler

import (
	"strconv"

	models "github.com/AxAxAxx/gofiber/modules/entities"
	usecase "github.com/AxAxAxx/gofiber/modules/users/usecases"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	BookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		BookUsecase: bookUsecase,
	}
}

func (h *BookHandler) GetBook(c *fiber.Ctx) error {
	books, err := h.BookUsecase.GetAllBooks()
	if err != nil {
		return c.JSON(models.BookErrorResponse(err))
	}
	return c.JSON(models.BooksSuccessResponse(&books))
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var newBook models.Books

	if err := c.BodyParser(&newBook); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		return c.JSON(models.BookErrorResponse(err))
	}

	err := h.BookUsecase.CreateBook(newBook)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		return c.JSON(models.BookErrorResponse(err))
	}
	return c.JSON(models.BookSuccessResponse(newBook))
}

func (h *BookHandler) UpdateUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Invalid User ID")
		return c.JSON(models.BookErrorResponse(err))
	}

	var updatedUser models.Books
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.JSON(models.BookErrorResponse(err))
	}

	err = h.BookUsecase.UpdateUser(userID, updatedUser)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		return c.JSON(models.BookErrorResponse(err))
	}

	return c.JSON(models.BookSuccessResponse(updatedUser))
}

func (h *BookHandler) DeleteUser(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Invalid Book ID")
		return c.JSON(models.BookErrorResponse(err))
	}

	err = h.BookUsecase.DeleteUser(bookID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		return c.JSON(models.BookErrorResponse(err))
	}

	return c.JSON(models.BookDelete())
}
