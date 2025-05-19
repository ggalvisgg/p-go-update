package services

import (
    //"fmt"
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/repositories"
)

type BookServiceInterface interface {
    UpdateBook(book *models.Book) (*models.Book, error)
}

type BookService struct {
    repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
    return &BookService{repo}
}

func (s *BookService) UpdateBook(book *models.Book) (*models.Book, error) {
    return s.repo.UpdateBook(book)
}