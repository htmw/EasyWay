package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"net/smtp"
	"crypto/rand"

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

func ForgotUsername(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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

    // Send the username to the user's email
    sendEmail(user.Email, user.Username)

    respondJSON(w, http.StatusOK, map[string]string{"message": "Username sent to email"})
}

func sendEmail(email, username string) error {
    from := "your-email@example.com"
    password := "your-email-password"
    to := []string{email}

    msg := "From: " + from + "\n" +
        "To: " + email + "\n" +
        "Subject: Your username on our platform\n\n" +
        "Hello,\n\n" +
        "Your username on our platform is: " + username + "\n\n" +
        "Thanks for using our platform!"

    auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

    err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, []byte(msg))
    if err != nil {
        return err
    }

    return nil
}

func ForgotPassword(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

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

    // Generate a new random password
    newPassword := generateRandomString(10)

    // Update the user's password in the database
    if err := db.Model(&user).Update("password", newPassword).Error; err != nil {
        respondError(w, http.StatusInternalServerError, err.Error())
        return
    }

    // Send the new password to the user's email
    sendEmail(user.Email, newPassword)

    respondJSON(w, http.StatusOK, map[string]string{"message": "New password sent to email"})
}

func generateRandomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    randomBytes := make([]byte, length)
    rand.Read(randomBytes)
    for i := 0; i < length; i++ {
        randomBytes[i] = charset[int(randomBytes[i])%len(charset)]
    }
    return string(randomBytes)
}


func sendPasswordResetEmail(email, resetLink string) error {
    from := "noreply@easyway.com"
    password := "your-email-password"
    to := []string{email}

    msg := "From: " + from + "\n" +
        "To: " + email + "\n" +
        "Subject: Password reset instructions\n\n" +
        "Hello,\n\n" +
        "Please click on the following link to reset your password: " + resetLink + "\n\n" +
        "If you did not request a password reset, please ignore this email.\n\n" +
        "Thanks for using our platform!"

    auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

    err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, []byte(msg))
    if err != nil {
        return err
    }

    return nil
}
