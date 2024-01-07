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

	r := mux.NewRouter()
	uh := handlers.NewUploadHandler()
	fh := handlers.NewFetchEntityFilesHandler()

	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(dir))))
	r.HandleFunc("/upload", uh.Handle).Methods(http.MethodPost)
	r.Handle("/{entity}/{entityId}", fh).Methods(http.MethodGet)

	server := http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
