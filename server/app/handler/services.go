package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/ksharma67/EasyWay/server/app/model"
)

func GetAllServices(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	services := []model.Service{}
	db.Find(&services)
	respondJSON(w, http.StatusOK, services)
}

func GetServiceInfo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	service := model.Service{}

	serviceId := r.URL.Query()["serviceId"]
	i, err := strconv.Atoi(serviceId[0])
	if err == nil {
		fmt.Println("No error")
	}

	if err := db.Where("id = ?", i).Find(&service).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}

	respondJSON(w, http.StatusOK, service)
}

type city struct {
	cityname string
}

func GetServicesInCity(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	cityServiceMap := []model.CityServiceMapping{}

	err := r.ParseForm()
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	// cityName := r.Form.Get("cityname")

	c := model.CityServiceMapping{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&c); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	// fmt.Fprintf(w, "city: %+v", c)
	fmt.Printf("cityname : %s", c.CityName)

	var serviceList []model.Service

	db.Where("city_name = ?", c.CityName).Find(&cityServiceMap)

	for _, cityService := range cityServiceMap {
		service := model.Service{}
		err := db.Where("id = ?", cityService.ServiceId).First(&service).Error
		if err == nil {
			serviceList = append(serviceList, service)
		}
	}
	respondJSON(w, http.StatusOK, serviceList)
}

func CreateService(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	services := model.Service{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&services); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&services).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, services)
}

// SearchServiceByName searches for services that match the given name
func SearchServiceByName(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
  w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	query := r.URL.Query().Get("name")

	if len(query) >= 3 {
		var services []model.Service
		db.Where("name LIKE ?", "%"+query+"%").Find(&services)

		if len(services) > 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(services)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
