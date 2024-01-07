package main

import (
	"fmt"
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
	r.Handle("/upload", uh).Methods(http.MethodPost)
	r.Handle("/{entity}/{entityId}", fh).Methods(http.MethodGet)

	port := os.Getenv("SERVER_PORT")
	fmt.Println(port)
	server := http.Server{
		Addr:    port,
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
