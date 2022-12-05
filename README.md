# EasyWay

[![Netlify Status](https://api.netlify.com/api/v1/badges/57f211b7-3340-4374-a825-55fa2f4b2a82/deploy-status)](https://app.netlify.com/sites/easywayapp/deploys)  [![Go](https://github.com/ksharma67/EasyWay/actions/workflows/go.yml/badge.svg)](https://github.com/ksharma67/EasyWay/actions/workflows/go.yml)  [![Node.js CI](https://github.com/ksharma67/EasyWay/actions/workflows/node.js.yml/badge.svg)](https://github.com/ksharma67/EasyWay/actions/workflows/node.js.yml)

The “EasyWay" web application aims to aggregate utility services such as beauty, electrical maintenance, home cleaning, pest control, etc. The application would enable the end-user to select their preferred service, book an appointment at a convenient time, pay the resultant charge and give feedback. The primary criterion of the web application would facilitate easy calendar and time slot booking to book the services according to your time and availability. The application aims to be a one-stop shop that caters to all the utility needs of the end user. Our product promises easy booking and cancellation without extra changes and you can maintain your booked services in one place easily.

## Design

The frontend of EasyWay will be implemented using Angular JS. Users should thus be able to view and interact with EasyWay within all supported browsers. The backend of EasyWay is implemented using Node JS, with Go Lang as its web framework. And the database system which we are going to be used will be MySQL.

## File Structure

Our application is structured as follows:

| File Name   | Description                                                            |
|--------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ProjectDocs      | This folder contains all the Project Deliverable files featured on the project Wiki page.    
| TeamPhotos   | This folder contains the photos of each team member that are used on the project Wiki page.   
| client   | This folder contains the codes for Front End.   
| db   | This folder contains the database schema and dummy data.   
| server   | This folder contains the codes for Back End server.  

## Technology Stack:
* Framework : Angular
* Backend : GoLang
* Database : MySQL (GORM Library)
* Version Control: Git
* Code Editor : Visual Studio Code

## Project Board:

Link : https://github.com/users/ksharma67/projects/2

## API Documentation:

Link : https://github.com/ksharma67/EasyWay/blob/main/API%20Documentation.md

## Running Backend Server:

* Clone the repository
```
git clone https://github.com/ksharma67/EasyWay.git
```
* Make sure you have mysql installed and correctly set up.
* Create a new database in MySQL using:
```
mysql -u root -p
```
Enter mysql password, then run:
```
create database easyWay;
```
* Goto config.go and update your mysql password
```
cd server/config/
code config.go
```
* Now navigate to server folder and run go server:
```
cd ./server/
go run main.go
```
Ignore any errors as it will check for required datatables (show the error), then automatically creates the datatables.

## Running Frontend Server:

Link : https://easywayapp.netlify.app

* Clone the repository
```
git clone https://github.com/ksharma67/EasyWay.git
```
* Install NodeJS LTS version from https://nodejs.org/en/ for your Operating System.
* Navigate to client folder and install required libraries:
```
cd ./client/
npm install
```
* In case of any error run audit and install once more:
```
npm audit fix --force && npm install
```
* Run the Angular Server:
```
npm start
```
