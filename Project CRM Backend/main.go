package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string
	Name      string
	Role      string
	Email     string
	Phone     int
	Contacted bool
}

var c1 = Customer{ID: uuid.New().String(), Name: "Shubham", Role: "ML Developer", Email: "Shubham@gmail.com", Phone: 5556769789, Contacted: true}
var c2 = Customer{ID: uuid.New().String(), Name: "Alex", Role: "Software Developer", Email: "Alex@gmail.com", Phone: 7756769789, Contacted: false}
var c3 = Customer{ID: uuid.New().String(), Name: "Tony", Role: "Data Engineer", Email: "Tony@gmail.com", Phone: 8796769449, Contacted: true}

var database = []Customer{c1, c2, c3}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(database)

	w.WriteHeader(http.StatusOK)

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	found := false
	for _, customer := range database {
		if customer.ID == key {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customer)
			w.WriteHeader(http.StatusOK)
			found = true
		}
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Customer ID not found")

	}

}

func addCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := io.ReadAll(r.Body)

	var customer Customer

	json.Unmarshal(reqBody, &customer)
	customer.ID = uuid.New().String()
	database = append(database, customer)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(database)

}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := io.ReadAll(r.Body)

	var customer Customer

	json.Unmarshal(reqBody, &customer)
	found := false

	for index, value := range database {
		if value.ID == key {
			database[index] = customer
			database[index].ID = key
			json.NewEncoder(w).Encode(database)
			w.WriteHeader(http.StatusOK)
			found = true
		}
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Customer ID not found")
	}

}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	found := false

	for index, customer := range database {
		if customer.ID == key {

			database = append(database[:index], database[index+1:]...)
			json.NewEncoder(w).Encode(database)
			w.WriteHeader(http.StatusOK)
			found = true
		}
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Customer ID not found")
	}

}

func main() {

	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))

	router.Handle("/", fileServer).Methods("GET")

	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 8000...")
	http.ListenAndServe(":8000", router)

}
