package repositories

import (
	"go-mvc-api/config"
	"go-mvc-api/models"
)

func FindAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := config.DB.Find(&books).Error
	return books, err
}

func FindBookByID(id uint) (*models.Book, error) {
	var book models.Book
	err := config.DB.First(&book, id).Error
	return &book, err
}

func CreateBook(book *models.Book) error {
	return config.DB.Create(book).Error
}

func UpdateBook(book *models.Book) error {
	return config.DB.Save(book).Error
}

func DeleteBook(book *models.Book) error {
	return config.DB.Delete(book).Error
}
