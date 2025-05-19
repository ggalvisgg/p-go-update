package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"example.com/go-mongo-app/models"
	"os"
	"log"
	// Importaciones futuras:
	// "go.mongodb.org/mongo-driver/bson/primitive" // ← Se usará si se trabaja con ObjectID manualmente
	// "fmt" // ← Útil para depuración con fmt.Println()
)

type BookRepository struct {
	collection *mongo.Collection
}

func NewBookRepository() *BookRepository {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in environment")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("library").Collection("books")
	return &BookRepository{collection}
}

func (r *BookRepository) UpdateBook(book *models.Book) (*models.Book, error) {
	filter := bson.M{"_id": book.ID}
	update := bson.M{"$set": book}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return book, nil
}
