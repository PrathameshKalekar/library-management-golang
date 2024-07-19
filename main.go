package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/PrathameshKalekar/library-management/books"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while readeing env file")
	}
	MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal("Error conneting MongoDB")
	} else {
		log.Println("Database connected ")
	}

	if err = MongoClient.Database(os.Getenv("DATABASE_NAME")).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatal("ERROR connecting databse ")
	}
	log.Println("Ping Success")
}

func main() {
	defer MongoClient.Disconnect(context.Background())
	router := mux.NewRouter()
	libraryAPT := "/api/library"
	collection := MongoClient.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("COLLECTION_BOOK"))

	BooksService := books.BooksService{MongoCollection: collection}
	router.HandleFunc(libraryAPT+"/books", BooksService.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc(libraryAPT+"/books/{name}", BooksService.GetBookByName).Methods(http.MethodPost)
	router.HandleFunc(libraryAPT+"/books/{name}", BooksService.UpdateBookByName).Methods(http.MethodPut)

	log.Printf("Server running on port %s ...", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
