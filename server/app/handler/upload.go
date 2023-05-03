package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
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
		log.Printf("Error parsing multipart form: %v\n", err)
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error getting file from form: %v\n", err)
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	// Generate a unique filename for the uploaded file
	fileName := uuid.New().String() + filepath.Ext(handler.Filename)

	// Create a new file in the uploads directory
	f, err := os.Create(filepath.Join("./uploads", fileName))
	if err != nil {
		log.Printf("Error creating new file: %v\n", err)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()

	// Copy the file to the new file
	_, err = io.Copy(f, file)
	if err != nil {
		log.Printf("Error copying file: %v\n", err)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Create a new UploadedFile
	uploadedFile := model.UploadedFile{
		FileName:    handler.Filename,
		Size:        handler.Size,
		FilePath:    "./uploads/" + fileName,
		ContentType: handler.Header.Get("Content-Type"),
	}

	// Save the uploaded file to the database
	if err := db.Create(&uploadedFile).Error; err != nil {
		log.Printf("Error saving uploaded file to the database: %v\n", err)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Send the uploaded file to the Python server for object detection
	responseBody, err := SendImageToPythonServer(uploadedFile.FilePath)
	if err != nil {
		log.Printf("Error sending image to Python server: %v\n", err)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Parse the JSON response from the Python server
	var detectionResults model.DetectionResult
	if err := json.Unmarshal(responseBody, &detectionResults); err != nil {
	    log.Printf("Error parsing JSON response: %v\n", err)
	    respondError(w, http.StatusInternalServerError, err.Error())
	    return
	}


	// Return the detection results to the client
	respondJSON(w, http.StatusOK, detectionResults)
}

// SendImageToPythonServer sends an image file to the Python server and returns the response
func SendImageToPythonServer(filename string) ([]byte, error) {
    log.Printf("Opening image file %s", filename)

    // Open the image file
    file, err := os.Open(filename)
    if err != nil {
        log.Printf("Error opening image file: %v", err)
        return nil, err
    }
    defer file.Close()

    log.Println("Creating multipart form body")

    // Create a new multipart form body
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    log.Println("Adding image file to form body")

    // Add the image file to the form body
    part, err := writer.CreateFormFile("images", filepath.Base(filename))
    if err != nil {
        log.Printf("Error adding image file to form body: %v", err)
        return nil, err
    }
    _, err = io.Copy(part, file)
    if err != nil {
        log.Printf("Error copying image file to form body: %v", err)
        return nil, err
    }

    log.Println("Closing multipart form body")

    // Close the multipart form body
    err = writer.Close()
    if err != nil {
        log.Printf("Error closing multipart form body: %v", err)
        return nil, err
    }

    log.Println("Creating HTTP request")

    // Create a new HTTP request to send to the Python server
    req, err := http.NewRequest("POST", "http://localhost:3600/detections", body)
    if err != nil {
        log.Printf("Error creating HTTP request: %v", err)
        return nil, err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())

    log.Println("Sending request to Python server")

    // Send the request to the Python server
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Error sending request to Python server: %v", err)
        return nil, err
    }
    defer resp.Body.Close()

    log.Println("Reading response body")

    // Read the response body
    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Error reading response body: %v", err)
        return nil, err
    }

    return responseBody, nil
}

func GetDetectionImage(w http.ResponseWriter, r *http.Request) {
	// Extract the folder name from the query parameters
	folderName := r.URL.Query().Get("folderName")

	// Get the path to the folder
	folderPath := filepath.Join(".", folderName)

	// Open the file
	filePath := filepath.Join(folderPath, "detection1.jpg")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "image/jpeg")

	// Write the file contents to the response body
	w.Write(file)
}
