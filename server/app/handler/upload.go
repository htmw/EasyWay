package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/ksharma67/EasyWay/server/app/model"
)

// CreateUploadedFile creates a new uploaded file in the database
func CreateUploadedFile(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Parse the multipart form
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("file")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	// Generate a unique filename for the uploaded file
	fileName := uuid.New().String() + filepath.Ext(handler.Filename)

	// Create a new file in the uploads directory
	f, err := os.Create(filepath.Join("./uploads", fileName))
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()

	// Copy the file to the new file
	io.Copy(f, file)

	// Create a new UploadedFile
	uploadedFile := model.UploadedFile{
		FileName:    handler.Filename,
		Size:        handler.Size,
		FilePath:    "/uploads/" + fileName,
		ContentType: handler.Header.Get("Content-Type"),
	}

	// Save the uploaded file to the database
	if err := db.Create(&uploadedFile).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Return the created uploaded file to the client
	respondJSON(w, http.StatusCreated, uploadedFile)
}
