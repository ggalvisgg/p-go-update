package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"example.com/go-mongo-app/controllers"
	"example.com/go-mongo-app/models"
)

type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) UpdateBook(book *models.Book) (*models.Book, error) {
	args := m.Called(book)
	return args.Get(0).(*models.Book), args.Error(1)
}

func TestUpdateBook_Success(t *testing.T) {
	mockService := new(MockBookService)
	controller := controllers.NewBookController(mockService)

	id := primitive.NewObjectID().Hex()
	updatedBook := &models.Book{
		ID:     primitive.NewObjectID(),
		Title:  "Updated Title",
		Author: "Updated Author",
		ISBN:   "0987654321",
	}

	body, _ := json.Marshal(updatedBook)
	req := httptest.NewRequest("PUT", "/books/"+id, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": id})
	resp := httptest.NewRecorder()

	mockService.On("UpdateBook", mock.MatchedBy(func(b *models.Book) bool {
		return b.Title == updatedBook.Title && b.Author == updatedBook.Author && b.ISBN == updatedBook.ISBN
	})).Return(updatedBook, nil)

	controller.UpdateBook(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateBook_InvalidID(t *testing.T) {
	mockService := new(MockBookService)
	controller := controllers.NewBookController(mockService)

	req := httptest.NewRequest("PUT", "/books/invalid-id", bytes.NewReader([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": "invalid-id"})
	resp := httptest.NewRecorder()

	controller.UpdateBook(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestUpdateBook_InvalidData(t *testing.T) {
	mockService := new(MockBookService)
	controller := controllers.NewBookController(mockService)

	id := primitive.NewObjectID().Hex()
	req := httptest.NewRequest("PUT", "/books/"+id, bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": id})
	resp := httptest.NewRecorder()

	controller.UpdateBook(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
