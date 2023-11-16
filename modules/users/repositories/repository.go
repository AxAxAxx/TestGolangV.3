package repository

import (
	"log"

	models "github.com/AxAxAxx/gofiber/modules/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type BookRepositoty struct {
	DB *sqlx.DB
}

func NewBoolRepository(db *sqlx.DB) *BookRepositoty {
	return &BookRepositoty{
		DB: db,
	}
}

type Repository interface {
	CreateBook(book *models.Books) (*models.Books, error)
}

func (r *BookRepositoty) GetAllBook(c *fiber.Ctx) ([]models.Books, error) {
	var books []models.Books
	err := r.DB.Select(&books, "SELECT * FROM books Order By id")
	if err != nil {
		return nil, c.JSON(models.BookErrorResponse(err))
	}
	return books, nil
}

func (r *BookRepositoty) CreateBook(book models.Books) error {
	_, err := r.DB.Exec("INSERT INTO books (title, author, price) VALUES ($1, $2, $3)", book.Title, book.Author, book.Price)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (r *BookRepositoty) UpdateUser(bookID int, updatedBook models.Books) error {
	_, err := r.DB.Exec("UPDATE public.books SET title = $1, author = $2, price = $3 WHERE id = $4", updatedBook.Title, updatedBook.Author, updatedBook.Price, bookID)
	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepositoty) DeleteUser(bookID int) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = $1", bookID)
	if err != nil {
		return err
	}
	return nil
}
