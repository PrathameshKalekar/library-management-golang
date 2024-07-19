package books

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/PrathameshKalekar/library-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BooksService struct {
	MongoCollection *mongo.Collection
}

type ResponseJSON struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func (books *BooksService) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	bookList, err := books.MongoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	var retrievedBooks []models.Book
	if err = bookList.All(context.Background(), &retrievedBooks); err != nil {
		http.Error(w, "NO BOOKS", http.StatusForbidden)
		return
	}

	response := ResponseJSON{
		Message: "Books retrieved successfully",
		Status:  true,
		Data:    retrievedBooks,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
