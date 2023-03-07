package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"gopkg.in/gomail.v2"
)

type ForgotPasswordRequest struct {
    Email string `json:"email"`
}

func generateToken() (string, error) {
    tokenBytes := make([]byte, 32)
    _, err := rand.Read(tokenBytes)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(tokenBytes), nil
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
    email := r.FormValue("email")

    // Generate a password reset token
		token, err := generateToken()
		if err != nil {
	    http.Error(w, "Error generating token", http.StatusInternalServerError)
	    return
	}

    // Perform the necessary logic to send a password reset email
    // Create a new message
    message := gomail.NewMessage()
    message.SetHeader("From", "noreply@easyway.com")
    message.SetHeader("To", email)
    message.SetHeader("Subject", "Password reset request")
    message.SetBody("text/html", fmt.Sprintf("Here is your password reset link: <a href='http://example.com/reset?token=%s'>http://example.com/reset?token=%s</a>", token, token))

    // Create a new dialer to connect to Mailtrap
    dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, "b587fb03c51f13", "7f047af2a761af")

    // Send the message
    if err := dialer.DialAndSend(message); err != nil {
        http.Error(w, "Error sending email", http.StatusInternalServerError)
        return
    }

    // Send a success response
    w.WriteHeader(http.StatusOK)
}
