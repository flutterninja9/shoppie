package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/flutterninja9/shoppie/image_service/models"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (u *UploadHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleFileUpload(rw, r)
		return
	}

	http.Error(rw, "Method not supported", http.StatusMethodNotAllowed)
}

func handleFileUpload(rw http.ResponseWriter, r *http.Request) {
	parseErr := r.ParseMultipartForm(2097152) // 2MB max limit
	if parseErr != nil {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	var model = models.FileModel{}
	file, header, err := r.FormFile("file")
	entityIdStr := r.FormValue("entity_id")
	ent_type := r.FormValue("entity_type")
	model.EntityId = entityIdStr
	model.EntityType = models.EntityType(ent_type)

	if err != nil {
		http.Error(rw, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	saveDir := "./files/" + entityIdStr
	savePath := saveDir + "/" + header.Filename
	dirCreateErr := os.MkdirAll(saveDir, os.ModePerm)
	if dirCreateErr != nil {
		http.Error(rw, "Error creating dir", http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(savePath)
	if err != nil {
		http.Error(rw, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(rw, "Error saving file", http.StatusInternalServerError)
		return
	}

	model.Path = entityIdStr + "/" + header.Filename
	model.Name = header.Filename
	model.Type = strings.Split(header.Filename, ".")[1]
	saveErr := model.Save()
	if saveErr != nil {
		http.Error(rw, "Error saving to db", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, "File uploaded successfully: %s", header.Filename)
}
