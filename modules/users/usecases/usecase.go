package controller

import (
	models "github.com/AxAxAxx/gofiber/modules/entities"
	repository "github.com/AxAxAxx/gofiber/modules/users/repositories"
	"github.com/gofiber/fiber/v2"
)

type BookUsecase struct {
	BookRepository repository.BookRepositoty
}

func NewUserUsecase(bookRepository repository.BookRepositoty) *BookUsecase {
	return &BookUsecase{
		BookRepository: bookRepository,
	}
}

func (u *BookUsecase) GetAllBooks() ([]models.Books, error) {
	var c fiber.Ctx
	books, err := u.BookRepository.GetAllBook(&c)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (u *BookUsecase) CreateBook(newBook models.Books) error {
	err := u.BookRepository.CreateBook(newBook)
	if err != nil {
		return err
	}
	return nil
}

func (u *BookUsecase) UpdateUser(userID int, updatedBook models.Books) error {
	err := u.BookRepository.UpdateUser(userID, updatedBook)
	if err != nil {
		return err
	}
	return nil
}

func (u *BookUsecase) DeleteUser(bookID int) error {
	err := u.BookRepository.DeleteUser(bookID)
	if err != nil {
		return err
	}
	return nil
}
