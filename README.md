## ABOUT THE PROJECT
This is a backend (i.e., server-side portion) of a CRM application without a frontend UI.While the server is on and active, users will be able to make their own HTTP requests to your server to perform CRUD operations. It consists of mock "database" to store customer data. To interact with the that data we have the following RESTFUL APIs

* GET http://localhost:3000/customers - Return All Customers present in database
* GET http://localhost:3000/customers/{id} - return specific customer if exist
* POST http://localhost:3000/customers - Adds a new customer to the database
* PATCH	http://localhost:3000/customers/{id} - Update the info of customers if exist
* DELETE http://localhost:3000/customers/{id}- delete a customer info if exist


## SETTING UP THE ENVIRONMENT
* Download Go - (https://go.dev/doc/install)
* Download Postman for your Operating System or using Postman on web - (https://www.postman.com/)

## INSTALLING PACKAGES
Once go is installed, you need to install some packages using the following commands
* github.com/google/uuid - `go get github.com/google/uuid`
* github.com/gorilla/mux - `go get github.com/gorilla/mux`


## STEPS TO RUN THE PROJECT
* Open the terminal/command line in the root of the folder
* Run `go run main.go`
* You can access the application using `http://localhost:8000`
