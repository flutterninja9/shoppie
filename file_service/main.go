package main

import (
	"log"
	"net/http"
	"os"

	db "github.com/flutterninja9/shoppie/image_service/database"
	"github.com/flutterninja9/shoppie/image_service/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	var dir = "./files"
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db.Setup()

	mux := mux.NewRouter()
	uh := handlers.NewUploadHandler()
	mux.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(dir))))

	mux.HandleFunc("/upload", uh.Handle).Methods("POST")

	server := http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
