package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ksharma67/EasyWay/server/app"
	"github.com/ksharma67/EasyWay/server/app/model"
	"github.com/ksharma67/EasyWay/server/config"
	"github.com/stretchr/testify/assert"
)

// var a main.App

var a = &app.App{}
var dbName string = "test.db"

var users []model.User
var bookings []model.Booking
var user model.User
var booking model.Booking

func setUpTestDb() {
		// Drop the tables if they exist
		a.DB.AutoMigrate().DropTable(&model.User{})
		a.DB.AutoMigrate().DropTable(&model.Service{})
		a.DB.AutoMigrate().DropTable(&model.Booking{})
		a.DB.AutoMigrate().DropTable(&model.CityServiceMapping{})
		a.DB.AutoMigrate().DropTable(&model.Blog{})
		a.DB.AutoMigrate().DropTable(&model.Comment{})
		a.DB.AutoMigrate(&model.UploadedFile{})
		a.DB.AutoMigrate(&model.EmailValidationResponse{})
		a.DB.AutoMigrate(&model.ForgotUsernameReqBody{})
		a.DB.AutoMigrate(&model.ForgotPasswordReqBody{})
		a.DB.AutoMigrate(&model.CommentInput{})

		// Migrate the schema and create tables
		a.DB.AutoMigrate(&model.User{}, &model.Service{}, &model.Booking{}, &model.CityServiceMapping{}, &model.Blog{}, &model.Comment{}, &model.UploadedFile{}, &model.EmailValidationResponse{}, &model.ForgotUsernameReqBody{}, &model.ForgotPasswordReqBody{}, &model.CommentInput{})

		// Create some dummy users
		a.DB.Create(&model.User{
			Id:       1,
			Name:     "Dummy Duck",
			Username: "dummy",
			Password: "dumdum",
			Email:    "dummy@pace.edu",
			Gender:   "M",
		})
		a.DB.Create(&model.User{
			Id:       2,
			Name:     "Buzz Lightyear",
			Username: "buzz",
			Password: "busybee",
			Email:    "buzz@pace.edu",
			Gender:   "M",
		})
		a.DB.Create(&model.User{
			Id:       3,
			Name:     "Snow White",
			Username: "snow",
			Password: "abc1234",
			Email:    "snow@pace.edu",
			Gender:   "F",
		})

		// Create some dummy services
		a.DB.Create(&model.Service{
			Id:          1,
			Name:        "AC Maintanence",
			Description: "Any type of AC maintanence such as filter cleaning, part replacement, etc.",
			Category:    "Electronics",
			ImageName:   "air_conditioning.jpg",
			Price:       80,
		})
		a.DB.Create(&model.Service{
			Id:          2,
			Name:        "Plumbing",
			Description: "Sanitary and household plumbing. No sewage service.",
			Category:    "Household",
			ImageName:   "plumbing.jpg",
			Price:       100,
		})
		a.DB.Create(&model.Service{
			Id:          3,
			Name:        "Saloon",
			Description: "Haircut, massage, nailwork, makeup, etc.",
			Category:    "Personal Care",
			ImageName:   "saloon.jpg",
			Price:       25,
		})
		a.DB.Create(&model.Service{
			Id:          4,
			Name:        "Furniture Repair",
			Description: "Furniture frame repair, drilling, fitting new furniture, etc.",
			Category:    "Household",
			ImageName:   "furniture_repair.jpg",
			Price:       70,
		})
		a.DB.Create(&model.Service{
			Id:          5,
			Name:        "Exterminator",
			Description: "Pest control, wildlife evac, alligator emergency, etc.",
			Category:    "Animal/Pet",
			ImageName:   "pest_control.jpg",
			Price:       150,
		})

		// Create some dummy bookings
		a.DB.Create((&model.Booking{
			UserId:      1,
			ServiceId:   1,
			Date:        "2022-02-15",
			StartTime:   "12:30",
			EndTime:     "13:30",
			IsCancelled: false,
		}))
		a.DB.Create((&model.Booking{
			UserId:      1,
			ServiceId:   2,
			Date:        "2022-02-15",
			StartTime:   "16:30",
			EndTime:     "17:30",
			IsCancelled: false,
		}))
		a.DB.Create((&model.Booking{
			UserId:      2,
			ServiceId:   3,
			Date:        "2022-02-15",
			StartTime:   "16:30",
			EndTime:     "17:30",
			IsCancelled: false,
		}))

		// Create some dummy CityServiceMapping
		a.DB.Create((&model.CityServiceMapping{
			CityName:  "Newyork",
			ServiceId: 3,
		}))
		a.DB.Create((&model.CityServiceMapping{
			CityName:  "Newyork",
			ServiceId: 2,
		}))
		a.DB.Create((&model.CityServiceMapping{
			CityName:  "LA",
			ServiceId: 2,
		}))
		a.DB.Create((&model.CityServiceMapping{
			CityName:  "LA",
			ServiceId: 3,
		}))
		a.DB.Create((&model.CityServiceMapping{
			CityName:  "Boston",
			ServiceId: 1,
		}))
		a.DB.Create((&model.CityServiceMapping{
			CityName:  "Boston",
			ServiceId: 2,
		}))

		// Create some dummy Blog
		a.DB.Create(&model.Blog{
	    Id:          1,
	    Title:        "Why On-Demand Services are the Future of Convenience",
	    Content:      "On-demand services have become a popular trend in recent years due to their convenience and accessibility. With the rise of smartphones and apps, people are able to order almost anything they need at the touch of a button. From food delivery to ride-sharing to home cleaning services, on-demand services offer a wide range of options for consumers. They eliminate the need for physical stores and the inconvenience of waiting in long lines or driving to multiple locations. With the COVID-19 pandemic accelerating the shift towards online and contactless services, on-demand services have become even more crucial for people looking to get what they need quickly and safely.",
	    CreatedAt:   "2023-04-25 07:30",
	    UpdatedAt:   "2023-04-25 07:32",
			ImageName:   "house_cleaning.jpg",
		})
		a.DB.Create(&model.Blog{
		Id:          2,
		Title:        "The Pros and Cons of On-Demand Home Cleaning Services",
		Content:      "On-demand home cleaning services have become a popular choice for busy people looking to outsource their household chores. These services offer the convenience of scheduling and paying for cleaning online, without the need for face-to-face communication with the cleaners. However, there are some drawbacks to these services. One of the main concerns is the quality of the cleaning. With on-demand services, the cleaners may not be as thorough as you would like, and there may be a lack of consistency between different cleaners. Additionally, on-demand cleaning services can be more expensive than hiring a regular cleaner or cleaning yourself. They often charge a premium for the convenience and flexibility they offer.",
		CreatedAt:   "2023-04-25 09:15",
		UpdatedAt:   "2023-04-25 09:30",
		ImageName:   "refridgerator.jpg",
		})

		// Create some dummy Comments
		a.DB.Create(&model.Comment{
	    Id:        1,
	    BlogId:    1,
	    Content:   "Great article, very informative!",
	    CreatedAt: "2023-04-25 09:30",
	    UpdatedAt: "2023-04-25 09:30",
		})
		a.DB.Create(&model.Comment{
	    Id:        2,
	    BlogId:    1,
	    Content:   "I totally agree, on-demand services are the way to go!",
	    CreatedAt: "2023-04-25 10:15",
	    UpdatedAt: "2023-04-25 10:15",
		})
		a.DB.Create(&model.Comment{
	    Id:        3,
	    BlogId:    2,
	    Content:   "I had a bad experience with an on-demand cleaning service, they didn't do a good job",
	    CreatedAt: "2023-04-25 11:20",
	    UpdatedAt: "2023-04-25 11:20",
		})
	}

func TestMain(m *testing.M) {
	config := config.GetConfig()
	a.Initialize(config)
	setUpTestDb()
	code := m.Run()
	os.Exit(code)
}

func TestGetAUser(t *testing.T) {

	var jsonStr = []byte(`{"username":"dummy","password":"dumdum"}`)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Access-Control-Allow-Methods", "POST")
	req.Header.Set("Access-Control-Allow-Headers", "Content-Type")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(a.Login)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	assert.Equal(t, 200, w.Code)
}

func TestCreateAUser(t *testing.T) {

	var jsonStr = []byte(`{"Id":4,"Name":"xyz","Username":"xyz","Password":"xyz@pqr.com","Email":"xyz@gmail.com","Gender":"F"}`)

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))

	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Access-Control-Allow-Methods", "POST")
	req.Header.Set("Access-Control-Allow-Headers", "Content-Type")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(a.CreateUser)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expected := `{"id":4,"name":"xyz","username":"xyz","password":"xyz@pqr.com","email":"xyz@gmail.com","gender":"F"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}

}

func TestGetBookings(t *testing.T) {

	req, _ := http.NewRequest("GET", "/getBookings", nil)
	q := req.URL.Query()
	q.Add("userId", "1")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(a.GetBookings)
	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

//func TestCancelBooking(t *testing.T) {

	//var jsonStr = []byte(`{"user_id":1,"service_id":1}`)
	//req, _ := http.NewRequest("POST", "/cancelBooking", bytes.NewBuffer(jsonStr))
	//w := httptest.NewRecorder()
	//handler := http.HandlerFunc(a.CancelBooking)
	//handler.ServeHTTP(w, req)

//	assert.Equal(t, 200, w.Code)

//}

func TestGetCancelledBookings(t *testing.T) {

	req, _ := http.NewRequest("GET", "/getCancelledBookings", nil)
	q := req.URL.Query()
	q.Add("userId", "1")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(a.GetCancelledBookings)
	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetServices(t *testing.T) {

	req, _ := http.NewRequest("GET", "/getServices", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(a.GetAllServices)
	handler.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestGetServicesInCity(t *testing.T) {
	var jsonStr = []byte(`{"cityname":"Boston"}`)
	req, _ := http.NewRequest("POST", "/getServicesOfCity", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(a.GetServicesInCity)
	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}
