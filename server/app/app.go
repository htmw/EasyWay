package app

import (
	"fmt"
	"net/http"
	"log"

	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/ksharma67/EasyWay/server/app/handler"
	"github.com/ksharma67/EasyWay/server/config"
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

	// Enable CORS
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
    })

	// Use the CORS middleware
	a.Router.Use(c.Handler)

	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/getServices", a.GetAllServices)
	a.Post("/user", a.CreateUser)
	a.Get("/user/{username}", a.GetUser)
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
	a.Get("/getAllBlogs", a.GetAllBlogs)
	a.Get("/getAllComments/{id:[0-9]+}", a.GetAllComments)
	a.Get("/services/search", a.SearchServiceByName)
	a.Post("/forgotUsername", a.ForgotUsername)
	a.Post("/forgotPassword", a.ForgotPassword)
	a.Post("/createUploadedFile", a.CreateUploadedFile)
	a.Post("/blogs/{id}/comments", a.AddComment)
	a.Put("/updateBooking", a.EditBooking)
	a.Get("/detection", a.GetDetectionImage)
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

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.DB, w, r)
}

// Handlers to manage Services Data
func (a *App) GetAllServices(w http.ResponseWriter, r *http.Request) {
	handler.GetAllServices(a.DB, w, r)
}

// Handlers to Get Cancelled Bookings
func (a *App) GetCancelledBookings(w http.ResponseWriter, r *http.Request) {
	handler.GetCancelledBookings(a.DB, w, r)
}

// Handlers to Cancel Booking
func (a *App) CancelBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called Routes: /cancelBooking Method:POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.CancelBooking(a.DB, w, r)
}

// Handlers to Create Services
func (a *App) CreateService(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called Routes: /cancelBooking Method:POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.CreateService(a.DB, w, r)
}

// Handlers to Get Services In City
func (a *App) GetServicesInCity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called Routes: /cancelBooking Method:POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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

// Handlers to create Bookings
func (a *App) CreateBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called Routes: /User Method:POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.CreateBooking(a.DB, w, r)
}

// Handlers to manage Login
func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.Login(a.DB, w, r)
}

// Handlers to get Bookings Data
func (a *App) GetBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.GetBookings(a.DB, w, r)
}

// Handlers to Get Services Data
func (a *App) GetServiceInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.GetServiceInfo(a.DB, w, r)
}

// Handlers to Get User Data
func (a *App) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.GetUserDetails(a.DB, w, r)
}

// Handlers to Get Blogs Data
func (a *App) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  handler.GetAllBlogs(a.DB, w, r)
}

// Handlers to Get Comments Data
func (a *App) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  handler.GetAllComments(a.DB, w, r)
}

// Handlers to Search Services
func (a *App) SearchServiceByName(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
    handler.SearchServiceByName(a.DB, w, r)
}

//ForgotUsername
func (a *App) ForgotUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.ForgotUsername(a.DB, w, r)
}

//ForgotPassword
func (a *App) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	handler.ForgotPassword(a.DB, w, r)
}

//Upload
func (a *App) CreateUploadedFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  handler.CreateUploadedFile(a.DB, w, r)
}


//AddComment
func (a *App) AddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  handler.AddComment(a.DB, w, r)
}

// Handlers to update Bookings
func (a *App) EditBooking(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Called Routes: /bookService Method:PUT")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    handler.EditBooking(a.DB, w, r)
}

// Handlers to Get Detected Image
func (a *App) GetDetectionImage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Called Routes: /bookService Method:PUT")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    handler.GetDetectionImage(w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
