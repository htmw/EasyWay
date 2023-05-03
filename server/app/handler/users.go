package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"net/smtp"
	"crypto/rand"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sendgrid/sendgrid-go"
  "github.com/sendgrid/sendgrid-go/helpers/mail"
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

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username := vars["username"]
	user := getUserOr404NoPass(db, username, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}

// getUserOr404 gets a user instance if exists, or respond the 404 error otherwise
func getUserOr404NoPass(db *gorm.DB, username string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, model.User{Username: username}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	decoder := json.NewDecoder(r.Body)
	reqBody := model.ForgotUsernameReqBody{}
	if err := decoder.Decode(&reqBody); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	user := model.User{}
	if err := db.Where("email = ?", reqBody.Email).First(&user).Error; err != nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	// generate a temporary username
	tempUsername := generateRandomString(10)
	if err := db.Model(&user).Update("username", tempUsername).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// send email with temporary username
	err := sendForgotUsernameEmail(user.Email, tempUsername)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "A username has been sent to your email"})
}

func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	rand.Read(b)
	for i := 0; i < n; i++ {
		b[i] = letters[int(b[i])%len(letters)]
	}
	return string(b)
}

func sendForgotUsernameEmail(to string, tempUsername string) error {
    from := mail.NewEmail("EasyWay", "noreply@easyway.com")
    subject := "Forgot Username"
    toEmail := mail.NewEmail("", to)
    bodyMessage := "Dear user,\n\n" +
        "Your temporary username is " + tempUsername + ". Please use this username to login and reset your username.\n\n" +
        "Best regards,\n" +
        "The EasyWay team"
    plainTextContent := bodyMessage
    htmlContent := bodyMessage

    message := mail.NewSingleEmail(from, subject, toEmail, plainTextContent, htmlContent)
    client := sendgrid.NewSendClient("")
    _, err := client.Send(message)
    if err != nil {
        return err
    }
    return nil
}


// sendEmail sends an email with the given subject and message to the given email address using SMTP
func sendEmail(to, subject, message string) error {
	from := "easyway@example.com" // Update with your email address
	password := "yourpassword"    // Update with your email password

	// Setup SMTP connection
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	smtpAddr := "smtp.gmail.com:587"
	connection, err := smtp.Dial(smtpAddr)
	if err != nil {
		return err
	}
	defer connection.Close()

	if err = connection.StartTLS(nil); err != nil {
		return err
	}

	if err = connection.Auth(auth); err != nil {
		return err
	}

	if err = connection.Mail(from); err != nil {
		return err
	}

	if err = connection.Rcpt(to); err != nil {
		return err
	}

	messageBody := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		message + "\r\n")
	if err = smtp.SendMail(smtpAddr, auth, from, []string{to}, messageBody); err != nil {
		return err
	}

	return nil
}

func ForgotPassword(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	decoder := json.NewDecoder(r.Body)
	reqBody := model.ForgotPasswordReqBody{}
	if err := decoder.Decode(&reqBody); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	user := model.User{}
	if err := db.Where("email = ?", reqBody.Email).First(&user).Error; err != nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	// generate a temporary password
	tempPassword := generateRandomString(10)
	if err := db.Model(&user).Update("password", tempPassword).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// send email with temporary password
	err := sendForgotPasswordEmail(user.Email, tempPassword)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "A temporary password has been sent to your email"})
}

func sendForgotPasswordEmail(to string, tempPassword string) error {
	from := mail.NewEmail("EasyWay", "noreply@easyway.com")
	subject := "Forgot Password"
	toEmail := mail.NewEmail("", to)
	bodyMessage := "Dear user,\n\n" +
		"Your temporary password is " + tempPassword + ". Please use this password to login and reset your password.\n\n" +
		"Best regards,\n" +
		"The EasyWay team"
	plainTextContent := bodyMessage
	htmlContent := bodyMessage

	message := mail.NewSingleEmail(from, subject, toEmail, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("")
	_, err := client.Send(message)
	return err
}
