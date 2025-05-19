package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/services"

    // Las siguientes importaciones se utilizarán más adelante en el microservicio
    // y se comentan temporalmente para evitar errores de compilación.
    // "fmt"
    // "log"
)

type BookController struct {
    Service services.BookServiceInterface
}

func NewBookController(service services.BookServiceInterface) *BookController {
    return &BookController{Service: service}
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var book models.Book
    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        http.Error(w, "Datos inválidos", http.StatusBadRequest)
        return
    }

    book.ID, err = primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    updatedBook, err := c.Service.UpdateBook(&book)
    if err != nil {
        http.Error(w, "Error al actualizar libro", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Libro actualizado correctamente",
        "book":    updatedBook,
    })
}
