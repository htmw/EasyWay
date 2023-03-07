package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"math/rand"
  "time"

	"github.com/jinzhu/gorm"
	"github.com/ksharma67/EasyWay/server/app/model"
)

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("Logged CreateUser:POST")
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	userFromReqBody := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userFromReqBody); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	username := userFromReqBody.Username
	password := userFromReqBody.Password
	fmt.Printf("username : %s", username)
	fmt.Printf("password : %s", password)

	user := getUserOr404(db, username, password, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}

// getUserOr404 gets a user instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, username string, password string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}

	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &user
}

// getUserDetails gets a user's information to display in the user-profile
func GetUserDetails(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	serviceId := r.URL.Query()["userId"]
	i, err := strconv.Atoi(serviceId[0])
	if err == nil {
		fmt.Println("No error")
	}

	if err := db.Where("id = ?", i).First(&user).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}

	respondJSON(w, http.StatusOK, user)
}

func ForgotPassword(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    var user model.User
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&user); err != nil {
        respondError(w, http.StatusBadRequest, err.Error())
        return
    }

    defer r.Body.Close()

    // Check if the email exists in the database
    if err := db.Where("email = ?", user.Email).First(&user).Error; err != nil {
        respondError(w, http.StatusNotFound, "Email not found")
        return
    }

    // Generate a new password
    newPassword := generatePassword()

    // Update the user's password in the database
    if err := db.Model(&user).Update("password", newPassword).Error; err != nil {
        respondError(w, http.StatusInternalServerError, err.Error())
        return
    }

    // Send the new password to the user's email
    sendEmail(user.Email, newPassword)

    respondJSON(w, http.StatusOK, map[string]string{"message": "New password sent to email"})
}

func generatePassword() string {
    // Generate a random password with 10 characters
    const passwordLength = 10
    const passwordChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    password := make([]byte, passwordLength)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < passwordLength; i++ {
        password[i] = passwordChars[rand.Intn(len(passwordChars))]
    }
    return string(password)
}

func sendEmail(to string, password string) {
    // Implement code to send the email with the new password
}
