package model
// Import the necessary package for database dialect
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User model for database table users
type User struct {
	Id       uint   `gorm:"size:10;primary_key;" json:"id"`
	Name     string `gorm:"size:70" json:"name"`
	Username string `gorm:"size:20; unique" json:"username"`
	Password string `gorm:"size:20" json:"password"`
	Email    string `gorm:"size:50" json:"email"`
	Gender   string `gorm:"size:1; check:gender==M || gender==F" json:"gender"`
}

type EmailValidationResponse struct {
    Valid bool `gorm:"not_null" json:"valid"`
}

type ForgotUsernameReqBody struct {
		Email string `gorm:"size:50" json:"email"`
	}

type ForgotPasswordReqBody struct {
			Email string `gorm:"size:50" json:"email"`
		}

//Service model for database table Service
type Service struct {
	Id          uint   `gorm:"size:10;primary_key;" json:"id"`
	Name        string `gorm:"size:50" json:"name"`
	Description string `gorm:"size:200" json:"description"`
	Category    string `gorm:"size:30; default: 'Other'" json:"category"`
	ImageName   string `gorm:"size:80; default: 'default.jpg'" json:"image_name"`
	Price       int    `gorm:"default:100" json:"price"`
}

//Booking model for database table Booking
type Booking struct {
	Id          int     `gorm:"primary_key" json:"id"`
	UserId      uint    `gorm:"not_null" json:"user_id"`
	ServiceId   uint    `gorm:"not_null" json:"service_id"`
	Date        string  `gorm:"size:11" json:"date"`
	StartTime   string  `gorm:"size:5" json:"start_time"`
	EndTime     string  `gorm:"size:5" json:"end_time"`
	IsCancelled bool    `gorm:"type:bool;default:false" json:"is_cancelled"`
	Note        *string `gorm:"size:255" json:"note,omitempty"`
}


//City model for database table City Service Mapping
type CityServiceMapping struct {
	Id        uint   `gorm:"size:10;unique;auto_increment:true" json:"id"`
	CityName  string `gorm:"size:200" json:"cityname"`
	ServiceId uint   `gorm:"size:10" json:"service_id"`
}

//Blog model for database table Blog
type Blog struct {
	Id         int 		`gorm:"primary_key " json:"id"`
	Title      string `gorm:"size:200" json:"title"`
	Content    string `gorm:"size:5000" json:"content"`
	CreatedAt  string `gorm:"size:20" json:"created_at"`
	UpdatedAt  string `gorm:"size:20" json:"updated_at"`
	ImageName  string `gorm:"size:80; default: 'default.jpg'" json:"image_name"`
}

//Comment model for database table Blog Commnets
type Comment struct {
    Id        int    `gorm:"primary_key" json:"id"`
    BlogId    int    `gorm:"not_null" json:"blog_id"`
    Content   string `gorm:"size:500" json:"content"`
    CreatedAt string `gorm:"size:20" json:"created_at"`
    UpdatedAt string `gorm:"size:20" json:"updated_at"`
}

type CommentInput struct {
    BlogId    int  	 `gorm:"not_null" json:"blog_id"`
  	Content   string `gorm:"size:500" json:"content"`
}

// Search model for database ServiceResult
type ServiceResult struct {
	Id   uint   `gorm:"primary_key" json:"id"`
	Name string `gorm:"size:50" json:"name"`
}

type UploadedFile struct {
    ID          uint     `gorm:"primary_key" json:"id"`
    ContentType string   `gorm:"not null" json:"content_type"`
    Size        int64    `gorm:"not null" json:"size"`
    FileName    string   `gorm:"not null" json:"file_name"`
    FilePath    string   `gorm:"not null" json:"file_path"`
}

type RespondError struct {
    Error string `gorm:"not null" json:"error"`
}


type DetectionResult struct {
	Response []struct {
		Detections []struct {
			Class      string  `json:"class"`
			Confidence float64 `json:"confidence"`
		} `json:"detections"`
		Image string `json:"image"`
	} `json:"response"`
}

type Response struct {
    Detections []Detection `json:"detections"`
    Image      string      `json:"image"`
}

type Detection struct {
    Class      string  `json:"class"`
    Confidence float64 `json:"confidence"`
}


type Object struct {
    Class string  `json:"class"`
    Score float64 `json:"score"`
    Box   Box     `json:"box"`
}

type Box struct {
    X1 int `json:"x1"`
    Y1 int `json:"y1"`
    X2 int `json:"x2"`
    Y2 int `json:"y2"`
}
