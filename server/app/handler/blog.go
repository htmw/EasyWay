package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/ksharma67/EasyWay/server/app/model"
)

// GetAllBlogs retrieves all blog posts.   a.Get("/getAllBlogs", a.GetAllBlogs)
func GetAllBlogs(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Query the database for all blog posts.
	blogs := []model.Blog{}
	if err := db.Find(&blogs).Error; err != nil {
		http.Error(w, "Error retrieving blogs", http.StatusInternalServerError)
		return
	}

	// Convert the blogs to JSON and write them to the response.
	jsonBytes, err := json.Marshal(blogs)
	if err != nil {
		http.Error(w, "Error converting blogs to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

// GetAllComments retrieves all comments for a specific blog post.  	a.Get("/getAllComments/{id:[0-9]+}", a.GetAllComments)
func GetAllComments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Get the blog ID from the request URL.
	BlogIdStr := mux.Vars(r)["id"]
	BlogId, err := strconv.Atoi(BlogIdStr)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	// Query the database for all comments for the specified blog post.
	comments := []model.Comment{}
	if err := db.Where("blog_id = ?", BlogId).Find(&comments).Error; err != nil {
		http.Error(w, "Error retrieving comments", http.StatusInternalServerError)
		return
	}

	// Convert the comments to JSON and write them to the response.
	jsonBytes, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, "Error converting comments to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}


// Post a new comment for a specific blog post. a.Post("/blogs/{id}/comments", a.AddComment)
func AddComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

    // Parse the comment input from the request body.
    commentInput := model.CommentInput{}
    if err := json.NewDecoder(r.Body).Decode(&commentInput); err != nil {
        http.Error(w, "Error parsing comment input", http.StatusBadRequest)
        return
    }

    // Create a new Comment object and populate it with the input data.
    comment := model.Comment{
        BlogId:  commentInput.BlogId,
        Content: commentInput.Content,
    }

    // Set the current time as the comment's created and updated timestamps.
    currentTime := time.Now().Format("2006-01-02 15:04:05")
    comment.CreatedAt = currentTime
    comment.UpdatedAt = currentTime

    // Save the new comment to the database.
    if err := db.Create(&comment).Error; err != nil {
        http.Error(w, "Error creating comment", http.StatusInternalServerError)
        return
    }

    // Convert the new comment to JSON and write it to the response.
    jsonBytes, err := json.Marshal(comment)
    if err != nil {
        http.Error(w, "Error converting comment to JSON", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBytes)
}
