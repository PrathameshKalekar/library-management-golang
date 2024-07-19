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

	if err = MongoClient.Database(os.Getenv("DATABASE_NAME")).RunCommand(context.TODO(), bson.D{{"ping", 0}}).Err(); err != nil {
		log.Fatal("ERROR connecting databse ")
	}
	log.Println("Ping Success")
}

func main() {
	defer MongoClient.Disconnect(context.Background())
	router := mux.NewRouter()
	collection := MongoClient.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("COLLECTION_BOOK"))

	BooksService := books.BooksService{MongoCollection: collection}
	router.HandleFunc("/api/library/books", BooksService.GetAllBooks).Methods(http.MethodGet)

	log.Printf("Server running on port %s ...", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
