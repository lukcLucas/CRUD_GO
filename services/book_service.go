package services

import (
	"go-mvc-api/models"
	"go-mvc-api/repositories"
)

func GetAllBooks() ([]models.Book, error) {
	return repositories.FindAllBooks()
}

func GetBook(id uint) (*models.Book, error) {
	return repositories.FindBookByID(id)
}

func AddBook(book *models.Book) error {
	return repositories.CreateBook(book)
}

func EditBook(book *models.Book) error {
	return repositories.UpdateBook(book)
}

func RemoveBook(book *models.Book) error {
	return repositories.DeleteBook(book)
}
