# EasyWay

[View Project Description as PDF](https://github.com/htmw/EasyWay/blob/main/ProjectDocs/Artifacts/Project%20Description/Project%20Description.pdf) | <a id="raw-url" href="https://github.com/htmw/EasyWay/blob/main/ProjectDocs/Artifacts/Project%20Description/Project%20Description.docx">Download Project Description as Word Document</a>

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
* Backend : GoLang, Flask
* Database : MySQL (GORM Library)
* Version Control: Git
* Code Editor : Visual Studio Code

## Project Board:

Link : https://github.com/users/ksharma67/projects/2

## API Documentation:

Link : https://documenter.getpostman.com/view/23815648/2s93eSZant

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

## Running Backend Server - Object Detection Server:

* Clone the repository
```
git clone https://github.com/htmw/EasyWay.git
```
* Install Python from https://www.python.org/downloads/
* Install Pip from https://pip.pypa.io/en/stable/installation/
* Navigate to server folder and run go server:
```
cd ./server/
```
* Install the required libraries
```
# TensorFlow CPU
pip install -r requirements.txt

# TensorFlow GPU
pip install -r requirements-gpu.txt
```
* For Linux: Let's download official yolov3 weights pretrained on COCO dataset.
```
# Downloading yolov3 weights
wget https://pjreddie.com/media/files/yolov3.weights -O weights/yolov3.weights
```
* Load the weights using `load_weights.py` script. This will convert the yolov3 weights into TensorFlow .ckpt model files!
```
# Loading yolov3 weights
python load_weights.py
```
* Starting the Flask Server
```
python app.py
```

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
