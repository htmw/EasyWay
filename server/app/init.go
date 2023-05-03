package app

import (
	"github.com/ksharma67/EasyWay/server/app/model"
)

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func (a *App) DBMigrate() {
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
