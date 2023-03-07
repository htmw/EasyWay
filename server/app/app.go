package app

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/ksharma67/EasyWay/server/app/handler"
	"github.com/ksharma67/EasyWay/server/app/handlers"
	"github.com/ksharma67/EasyWay/server/app/model"
	"github.com/ksharma67/EasyWay/server/config"
	"gopkg.in/gomail.v2"
	//"gorm.io/driver/sqlite"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, _ := gorm.Open(config.DB.Dialect, dbURI)

	a.DB = db
	a.Router = mux.NewRouter().PathPrefix("/api").Subrouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/getServices", a.GetAllServices)
	a.Post("/createService", a.CreateService)
	a.Post("/register", a.CreateUser)
	a.Post("/login", a.Login)
	a.Get("/getBookings", a.GetBookings)
	a.Post("/bookService", a.CreateBooking)
	a.Post("/getServicesOfCity", a.GetServicesInCity)
	a.Get("/cancelBooking", a.CancelBooking)
	a.Get("/getCancelledBookings", a.GetCancelledBookings)
	a.Get("/getServiceInfo", a.GetServiceInfo)
	a.Get("/getUserDetails", a.GetUserDetails)
	a.Get("/searchServices", a.SearchServices)
	a.Post("/forgotPassword", a.ForgotPassword)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Services Data
func (a *App) GetAllServices(w http.ResponseWriter, r *http.Request) {
	handler.GetAllServices(a.DB, w, r)
}

func (a *App) GetCancelledBookings(w http.ResponseWriter, r *http.Request) {
	handler.GetCancelledBookings(a.DB, w, r)
}

func (a *App) CancelBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called Routes: /cancelBooking Method:POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.CancelBooking(a.DB, w, r)
}

func (a *App) CreateService(w http.ResponseWriter, r *http.Request) {
	handler.CreateService(a.DB, w, r)
}

func (a *App) GetServicesInCity(w http.ResponseWriter, r *http.Request) {
	handler.GetServicesInCity(a.DB, w, r)
}

// Handlers to manager Users Data
func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called Routes: /User Method:POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//w.Header().Set("Accept", "application/x-www-form-urlencoded")
	handler.CreateUser(a.DB, w, r)
}

func (a *App) CreateBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called Routes: /User Method:POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.CreateBooking(a.DB, w, r)
}

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.Login(a.DB, w, r)
}

func (a *App) GetBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.GetBookings(a.DB, w, r)
}

func (a *App) GetServiceInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.GetServiceInfo(a.DB, w, r)
}

func (a *App) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.GetUserDetails(a.DB, w, r)
}

// Search services by name
func (a *App) SearchServices(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query().Get("q")
  if query == "" {
    http.Error(w, "Missing query parameter", http.StatusBadRequest)
    return
  }

  // Get a list of services that match the query
  var services []model.Service
  a.DB.Where("name LIKE ?", "%"+query+"%").Find(&services)

  // Return the list of services as a JSON response
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(services)
}


// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// ForgotPasswordHandler
func (a *App) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/forgotpassword", handlers.ForgotPasswordHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}

// Handler function for the "forgot username" form submission
func (a *App) ForgotUsername(w http.ResponseWriter, r *http.Request) {
	// Parse the email or phone number from the form data
	r.ParseForm()
	email := r.Form.Get("email")

	// Validate the input
	if email == "" {
		// Input is invalid
		http.Error(w, "Please enter your email", http.StatusBadRequest)
		return
	}

	// Query the database for the user(s) with the given email
	var users []model.User
	if email != "" {
		a.DB.Where("email = ?", email).Find(&users)
	}

	// Check if any users were found
	if len(users) == 0 {
		// No user found with the given email
		http.Error(w, "We couldn't find your account. Please check your email.", http.StatusNotFound)
		return
	}

	// Send the username(s) to the user via email
	for _, user := range users {
		message := gomail.NewMessage()
		message.SetHeader("From", "noreply@easyway.com")
		message.SetHeader("To", email)
		message.SetHeader("Subject", "Username request")
		fmt.Printf("Username for %s: %s\n", user.Name, user.Username)

		// Create a new dialer to connect to Mailtrap
		dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, "b587fb03c51f13", "7f047af2a761af")

		// Send the message
		if err := dialer.DialAndSend(message); err != nil {
			http.Error(w, "Error sending email", http.StatusInternalServerError)
			return
		}
	}

	// Inform the user that their username(s) have been sent
	fmt.Fprint(w, "Your username(s) have been sent to your email.")
}
