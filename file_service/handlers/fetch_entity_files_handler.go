package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flutterninja9/shoppie/image_service/models"
	"github.com/gorilla/mux"
)

type FetchEntityFilesHandler struct{}

func NewFetchEntityFilesHandler() *FetchEntityFilesHandler {
	return &FetchEntityFilesHandler{}
}

func (f *FetchEntityFilesHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	entityType := v["entity"]
	entityId := v["entityId"]

	files, err := models.GetFilesByEntityTypeAndId(entityType, entityId)

	if err != nil {
		http.Error(rw, "Unable to fetch files", http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(files)
	if err != nil {
		http.Error(rw, "Unable to convert in json", http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprint(rw, string(bytes))
}
