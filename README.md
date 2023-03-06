# EasyWay

[View Project Description as PDF](https://github.com/ksharma67/EasyWay/blob/main/ProjectDocs/Artifacts/Product%20Description.pdf) | <a id="raw-url" href="https://github.com/ksharma67/EasyWay/blob/main/ProjectDocs/Artifacts/Product%20Description.docx">Download Project Description as Word Document</a>

### Project Overview

The "EasyWay" web application is a comprehensive platform designed to aggregate various utility services, including beauty, electrical maintenance, home cleaning, pest control, and more. The primary objective of the application is to provide a convenient and hassle-free experience to the end-user, enabling them to book services, pay for them, and give feedback, all in one place.


### Key Features

The key features of the EasyWay web application include:

* **Easy service selection:** The end-user can select their preferred service from a list of available options.
Convenient appointment booking: The application facilitates easy calendar and time slot booking, allowing the user to schedule services at a convenient time.
* **Seamless payment process:** The end-user can pay for their services securely and conveniently.
* **Feedback mechanism:** The application enables the end-user to give feedback on the services they have availed, thus ensuring quality control.
* **One-stop-shop:** The application serves as a one-stop-shop, catering to all the utility needs of the end-user.


## Project Design

### Architecture Overview

The EasyWay web application is built using a client-server architecture, with the front-end implemented in Angular JS, the backend in Node JS, and the server in GOLang. The database system used is MySQL.

### Front-end Design

The front-end of the EasyWay web application is designed using Angular JS, a popular framework for building single-page applications. The front-end design includes the following components:

![Node.js Automated Testing](https://github.com/ksharma67/EasyWay/actions/workflows/node.js.yml/badge.svg) - [Front-end Smoke Testing](https://github.com/ksharma67/EasyWay/actions/workflows/node.js.yml)

* **User Interface:** The user interface is designed to be intuitive and user-friendly, with clear and concise layouts and color schemes.
* **Navigation:** The navigation system is designed to provide easy access to all the features of the application, with clearly labeled menus and icons.
* **Forms and Input Fields:** The forms and input fields are designed to be easy to use, with clear instructions and error messages.
Interactive Elements: The interactive elements, such as buttons and links, are designed to provide a responsive and smooth user experience.

### Back-end Design

The back-end of the EasyWay web application is designed using Node JS, a popular framework for building scalable and performant applications. The back-end design includes the following components:

![Go Automated Testing](https://github.com/ksharma67/EasyWay/actions/workflows/go.yml/badge.svg) - [Back-end Smoke Testing](https://github.com/ksharma67/EasyWay/actions/workflows/go.yml)

* **RESTful API:** The back-end provides a RESTful API for the front-end to communicate with the server.
* **Database Access:** The back-end interacts with the MySQL database system to store and retrieve data.
* **Server:** The server component of the back-end is implemented in GOLang, a high-performance programming language designed for building scalable and efficient applications.

### Deployment

The EasyWay web application is deployed on a cloud platform named Amazon Web Services (AWS). The front-end and back-end components can be deployed separately to ensure scalability and reliability.

![Deployment Status](https://api.netlify.com/api/v1/badges/57f211b7-3340-4374-a825-55fa2f4b2a82/deploy-status) - [Deployed Front-End](https://easywayapp.netlify.app/)


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

## API Documentation:

Link : https://github.com/htmw/EasyWay/blob/main/API%20Documentation.md

## Running Backend Server:

* Clone the repository
```
git clone https://github.com/htmw/EasyWay.git
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
git clone https://github.com/htmw/EasyWay.git
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
